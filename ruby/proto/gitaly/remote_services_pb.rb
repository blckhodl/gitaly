# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: remote.proto for package 'gitaly'

require 'grpc'
require 'remote_pb'

module Gitaly
  module RemoteService
    # RemoteService is a service providing RPCs to interact with a remote
    # repository that is hosted on another Git server.
    class Service

      include ::GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.RemoteService'

      # UpdateRemoteMirror compares the references in the target repository and its remote mirror
      # repository. Any differences in the references are then addressed by pushing the differing
      # references to the mirror. Created and modified references are updated, removed references are
      # deleted from the mirror. UpdateRemoteMirror updates all tags. Branches are updated if they match
      # the patterns specified in the requests.
      rpc :UpdateRemoteMirror, stream(::Gitaly::UpdateRemoteMirrorRequest), ::Gitaly::UpdateRemoteMirrorResponse
      # This comment is left unintentionally blank.
      rpc :FindRemoteRepository, ::Gitaly::FindRemoteRepositoryRequest, ::Gitaly::FindRemoteRepositoryResponse
      # FindRemoteRootRef tries to find the root reference of a remote
      # repository. The root reference is the default branch as pointed to by
      # the remotes HEAD reference. Returns an InvalidArgument error if the
      # specified remote does not exist and a NotFound error in case no HEAD
      # branch was found.
      rpc :FindRemoteRootRef, ::Gitaly::FindRemoteRootRefRequest, ::Gitaly::FindRemoteRootRefResponse
    end

    Stub = Service.rpc_stub_class
  end
end
