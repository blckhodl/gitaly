package praefect

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	logFieldRepositoryID          = "repository_id"
	logFieldReplicaPath           = "replica_path"
	logFieldAdditionalReplicaPath = "additional_replica_path"
	logFieldPrimary               = "primary"
	logFieldSecondaries           = "secondaries"
	logFieldStorage               = "storage"
)

func addRouteLogField(ctx context.Context, fields logrus.Fields) {
	ctxlogrus.AddFields(ctx, logrus.Fields{"route": fields})
}

func routerNodeStorages(secondaries []RouterNode) []string {
	storages := make([]string, len(secondaries))
	for i := range secondaries {
		storages[i] = secondaries[i].Storage
	}
	return storages
}

// RepositoryAccessorRoute describes how to route a repository scoped accessor call.
type RepositoryAccessorRoute struct {
	// ReplicaPath is the disk path where the replicas are stored.
	ReplicaPath string
	// Node contains the details of the node that should handle the request.
	Node RouterNode
}

func (r RepositoryAccessorRoute) addLogFields(ctx context.Context) {
	addRouteLogField(ctx, logrus.Fields{
		logFieldReplicaPath: r.ReplicaPath,
		logFieldStorage:     r.Node.Storage,
	})
}

// RouterNode is a subset of a node's configuration needed to perform
// request routing.
type RouterNode struct {
	// Storage is storage of the node.
	Storage string
	// Connection is the connection to the node.
	Connection *grpc.ClientConn
}

func (r RouterNode) addLogFields(ctx context.Context) {
	addRouteLogField(ctx, logrus.Fields{
		logFieldStorage: r.Storage,
	})
}

// StorageMutatorRoute describes how to route a storage scoped mutator call.
type StorageMutatorRoute struct {
	// Primary is the primary node of the routing decision.
	Primary RouterNode
	// Secondaries are the secondary nodes of the routing decision.
	Secondaries []RouterNode
}

func (r StorageMutatorRoute) addLogFields(ctx context.Context) {
	addRouteLogField(ctx, logrus.Fields{
		logFieldPrimary:     r.Primary,
		logFieldSecondaries: routerNodeStorages(r.Secondaries),
	})
}

// RepositoryMutatorRoute describes how to route a repository scoped mutator call.
type RepositoryMutatorRoute struct {
	// RepositoryID is the repository's ID as Praefect identifies it.
	RepositoryID int64
	// ReplicaPath is the disk path where the replicas are stored.
	ReplicaPath string
	// AdditionalReplicaPath is the disk path where the possible additional repository in the request
	// is stored. This is only used for object pools.
	AdditionalReplicaPath string
	// Primary is the primary node of the transaction.
	Primary RouterNode
	// Secondaries are the secondary participating in a transaction.
	Secondaries []RouterNode
	// ReplicationTargets are additional nodes that do not participate in a transaction
	// but need the changes replicated.
	ReplicationTargets []string
}

func (r RepositoryMutatorRoute) addLogFields(ctx context.Context) {
	addRouteLogField(ctx, logrus.Fields{
		logFieldRepositoryID:          r.RepositoryID,
		logFieldReplicaPath:           r.ReplicaPath,
		logFieldAdditionalReplicaPath: r.AdditionalReplicaPath,
		logFieldPrimary:               r.Primary,
		logFieldSecondaries:           routerNodeStorages(r.Secondaries),
		"replication_targets":         r.ReplicationTargets,
	})
}

// RepositoryMaintenanceRoute describes how to route a repository scoped maintenance call.
type RepositoryMaintenanceRoute struct {
	// RepositoryID is the repository's ID as Praefect identifies it.
	RepositoryID int64
	// ReplicaPath is the disk path where the replicas are stored.
	ReplicaPath string
	// Nodes contains all the nodes the call should be routed to.
	Nodes []RouterNode
}

func (r RepositoryMaintenanceRoute) addLogFields(ctx context.Context) {
	addRouteLogField(ctx, logrus.Fields{
		logFieldRepositoryID: r.RepositoryID,
		logFieldReplicaPath:  r.ReplicaPath,
		"storages":           routerNodeStorages(r.Nodes),
	})
}

// Router decides which nodes to direct accessor and mutator RPCs to.
type Router interface {
	// RouteStorageAccessor returns the node which should serve the storage accessor request.
	RouteStorageAccessor(ctx context.Context, virtualStorage string) (RouterNode, error)
	// RouteStorageAccessor returns the primary and secondaries that should handle the storage
	// mutator request.
	RouteStorageMutator(ctx context.Context, virtualStorage string) (StorageMutatorRoute, error)
	// RouteRepositoryAccessor returns the node that should serve the repository accessor
	// request. If forcePrimary is set to `true`, it returns the primary node.
	RouteRepositoryAccessor(ctx context.Context, virtualStorage, relativePath string, forcePrimary bool) (RepositoryAccessorRoute, error)
	// RouteRepositoryMutatorTransaction returns the primary and secondaries that should handle the repository mutator request.
	// Additionally, it returns nodes which should have the change replicated to. RouteRepositoryMutator should only be used
	// with existing repositories.
	RouteRepositoryMutator(ctx context.Context, virtualStorage, relativePath, additionalRepoRelativePath string) (RepositoryMutatorRoute, error)
	// RouteRepositoryCreation decides returns the primary and secondaries that should handle the repository creation
	// request. It is up to the caller to store the assignments and primary information after finishing the RPC.
	RouteRepositoryCreation(ctx context.Context, virtualStorage, relativePath, additionalRepoRelativePath string) (RepositoryMutatorRoute, error)
	// RouteRepositoryMaintenance routes the given maintenance-style RPC to all nodes which
	// should perform maintenance. This would typically include all online nodes, regardless of
	// whether the repository hosted by them is up-to-date or not. Maintenance tasks should
	// never be replicated.
	RouteRepositoryMaintenance(ctx context.Context, virtualStorage, relativePath string) (RepositoryMaintenanceRoute, error)
}
