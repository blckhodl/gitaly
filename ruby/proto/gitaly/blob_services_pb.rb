# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: blob.proto for package 'gitaly'

require 'grpc'
require 'blob_pb'

module Gitaly
  module BlobService
    # BlobService is a service which provides RPCs to retrieve Git blobs from a
    # specific repository.
    class Service

      include ::GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.BlobService'

      # GetBlob returns the contents of a blob object referenced by its object
      # ID. We use a stream to return a chunked arbitrarily large binary
      # response
      rpc :GetBlob, ::Gitaly::GetBlobRequest, stream(::Gitaly::GetBlobResponse)
      # This comment is left unintentionally blank.
      rpc :GetBlobs, ::Gitaly::GetBlobsRequest, stream(::Gitaly::GetBlobsResponse)
      # ListBlobs will list all blobs reachable from a given set of revisions by
      # doing a graph walk. It is not valid to pass revisions which do not resolve
      # to an existing object.
      rpc :ListBlobs, ::Gitaly::ListBlobsRequest, stream(::Gitaly::ListBlobsResponse)
      # ListAllBlobs retrieves all blobs pointers in the repository, including
      # those not reachable by any reference.
      rpc :ListAllBlobs, ::Gitaly::ListAllBlobsRequest, stream(::Gitaly::ListAllBlobsResponse)
      # GetLFSPointers retrieves LFS pointers from a given set of object IDs.
      # This RPC filters all requested objects and only returns those which refer
      # to a valid LFS pointer.
      rpc :GetLFSPointers, ::Gitaly::GetLFSPointersRequest, stream(::Gitaly::GetLFSPointersResponse)
      # ListLFSPointers retrieves LFS pointers reachable from a given set of
      # revisions by doing a graph walk. This includes both normal revisions like
      # an object ID or branch, but also the pseudo-revisions "--all" and "--not"
      # as documented in git-rev-parse(1). Revisions which don't directly or
      # transitively reference any LFS pointers are ignored. It is not valid to
      # pass revisions which do not resolve to an existing object.
      rpc :ListLFSPointers, ::Gitaly::ListLFSPointersRequest, stream(::Gitaly::ListLFSPointersResponse)
      # ListAllLFSPointers retrieves all LFS pointers in the repository, including
      # those not reachable by any reference.
      rpc :ListAllLFSPointers, ::Gitaly::ListAllLFSPointersRequest, stream(::Gitaly::ListAllLFSPointersResponse)
    end

    Stub = Service.rpc_stub_class
  end
end
