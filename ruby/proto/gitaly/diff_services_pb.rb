# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: diff.proto for package 'gitaly'

require 'grpc'
require 'diff_pb'

module Gitaly
  module DiffService
    # DiffService is a service which provides RPCs to inspect differences
    # introduced between a set of commits.
    class Service

      include ::GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.DiffService'

      # Returns stream of CommitDiffResponse with patches chunked over messages
      rpc :CommitDiff, ::Gitaly::CommitDiffRequest, stream(::Gitaly::CommitDiffResponse)
      # Return a stream so we can divide the response in chunks of deltas
      rpc :CommitDelta, ::Gitaly::CommitDeltaRequest, stream(::Gitaly::CommitDeltaResponse)
      # This comment is left unintentionally blank.
      rpc :RawDiff, ::Gitaly::RawDiffRequest, stream(::Gitaly::RawDiffResponse)
      # This comment is left unintentionally blank.
      rpc :RawPatch, ::Gitaly::RawPatchRequest, stream(::Gitaly::RawPatchResponse)
      # This comment is left unintentionally blank.
      rpc :DiffStats, ::Gitaly::DiffStatsRequest, stream(::Gitaly::DiffStatsResponse)
      # Return a list of files changed along with the status of each file
      rpc :FindChangedPaths, ::Gitaly::FindChangedPathsRequest, stream(::Gitaly::FindChangedPathsResponse)
    end

    Stub = Service.rpc_stub_class
  end
end
