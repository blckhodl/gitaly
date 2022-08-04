package limithandler

import (
	"context"
	"strings"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/prometheus/client_golang/prometheus"
)

const acquireDurationLogThreshold = 10 * time.Millisecond

// ConcurrencyMonitor allows the concurrency monitor to be observed
type ConcurrencyMonitor interface {
	Queued(ctx context.Context)
	Dequeued(ctx context.Context)
	Enter(ctx context.Context, acquireTime time.Duration)
	Exit(ctx context.Context)
	Dropped(ctx context.Context, message string)
}

type noopConcurrencyMonitor struct{}

func (c *noopConcurrencyMonitor) Queued(ctx context.Context)                           {}
func (c *noopConcurrencyMonitor) Dequeued(ctx context.Context)                         {}
func (c *noopConcurrencyMonitor) Enter(ctx context.Context, acquireTime time.Duration) {}
func (c *noopConcurrencyMonitor) Exit(ctx context.Context)                             {}
func (c *noopConcurrencyMonitor) Dropped(ctx context.Context, reason string)           {}

// NewNoopConcurrencyMonitor returns a noopConcurrencyMonitor
func NewNoopConcurrencyMonitor() ConcurrencyMonitor {
	return &noopConcurrencyMonitor{}
}

type promMonitor struct {
	queuedMetric           prometheus.Gauge
	inProgressMetric       prometheus.Gauge
	acquiringSecondsMetric prometheus.Observer
	requestsDroppedMetric  *prometheus.CounterVec

	acquiringSecondsHistogramVec *prometheus.HistogramVec
}

// newPerRPCPromMonitor creates a new ConcurrencyMonitor that tracks limiter
// activity in Prometheus.
func newPerRPCPromMonitor(
	system, fullMethod string,
	queuedMetric, inProgressMetric *prometheus.GaugeVec,
	acquiringSecondsMetric prometheus.ObserverVec,
	requestsDroppedMetric *prometheus.CounterVec,
) ConcurrencyMonitor {
	serviceName, methodName := splitMethodName(fullMethod)

	return &promMonitor{
		queuedMetric:           queuedMetric.WithLabelValues(system, serviceName, methodName),
		inProgressMetric:       inProgressMetric.WithLabelValues(system, serviceName, methodName),
		acquiringSecondsMetric: acquiringSecondsMetric.WithLabelValues(system, serviceName, methodName),
		requestsDroppedMetric: requestsDroppedMetric.MustCurryWith(prometheus.Labels{
			"system":       system,
			"grpc_service": serviceName,
			"grpc_method":  methodName,
		}),
	}
}

func (c *promMonitor) Queued(ctx context.Context) {
	c.queuedMetric.Inc()
}

func (c *promMonitor) Dequeued(ctx context.Context) {
	c.queuedMetric.Dec()
}

func (c *promMonitor) Enter(ctx context.Context, acquireTime time.Duration) {
	c.inProgressMetric.Inc()

	if acquireTime > acquireDurationLogThreshold {
		logger := ctxlogrus.Extract(ctx)
		logger.WithField("acquire_ms", acquireTime.Seconds()*1000).Info("Rate limit acquire wait")
	}

	c.acquiringSecondsMetric.Observe(acquireTime.Seconds())
}

func (c *promMonitor) Exit(ctx context.Context) {
	c.inProgressMetric.Dec()
}

func (c *promMonitor) Dropped(ctx context.Context, reason string) {
	c.requestsDroppedMetric.WithLabelValues(reason).Inc()
}

func newPromMonitor(
	keyType string,
	queuedVec, inProgressVec *prometheus.GaugeVec,
	acquiringSecondsVec *prometheus.HistogramVec,
	requestsDroppedVec *prometheus.CounterVec,
) *promMonitor {
	return &promMonitor{
		queuedMetric:                 queuedVec.WithLabelValues(keyType),
		inProgressMetric:             inProgressVec.WithLabelValues(keyType),
		acquiringSecondsMetric:       acquiringSecondsVec.WithLabelValues(keyType),
		requestsDroppedMetric:        requestsDroppedVec,
		acquiringSecondsHistogramVec: acquiringSecondsVec,
	}
}

func (p *promMonitor) Collect(metrics chan<- prometheus.Metric) {
	p.queuedMetric.Collect(metrics)
	p.inProgressMetric.Collect(metrics)
	p.acquiringSecondsHistogramVec.Collect(metrics)
	p.requestsDroppedMetric.Collect(metrics)
}

func (p *promMonitor) Describe(descs chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(p, descs)
}

func splitMethodName(fullMethodName string) (string, string) {
	fullMethodName = strings.TrimPrefix(fullMethodName, "/") // remove leading slash
	if i := strings.Index(fullMethodName, "/"); i >= 0 {
		return fullMethodName[:i], fullMethodName[i+1:]
	}
	return "unknown", "unknown"
}

// NewPackObjectsConcurrencyMonitor returns a concurrency monitor for use
// with limiting pack objects processes
func NewPackObjectsConcurrencyMonitor(keyType string, latencyBuckets []float64) *promMonitor {
	acquiringSecondsVec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "gitaly_pack_objects_acquiring_seconds",
			Help:    "Histogram of time calls are rate limited (in seconds)",
			Buckets: latencyBuckets,
		},
		[]string{"type"},
	)

	inProgressVec := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gitaly_pack_objects_in_progress",
			Help: "Gauge of number of concurrent in-progress calls",
		},
		[]string{"type"},
	)

	queuedVec := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "gitaly_pack_objects_queued",
			Help: "Gauge of number of queued calls",
		},
		[]string{"type"},
	)

	requestsDroppedVec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "gitaly_pack_objects_dropped_total",
			Help: "Number of requests dropped from the queue",
		},
		[]string{"type", "reason"},
	).MustCurryWith(prometheus.Labels{"type": keyType})

	return newPromMonitor(
		keyType,
		queuedVec,
		inProgressVec,
		acquiringSecondsVec,
		requestsDroppedVec,
	)
}
