# Generated by the protocol buffer compiler.  DO NOT EDIT!
# Source: cleanup.proto for package 'gitaly'

require 'grpc'
require 'cleanup_pb'

module Gitaly
  module CleanupService
    # CleanupService provides RPCs to clean up a repository's contents.
    class Service

      include ::GRPC::GenericService

      self.marshal_class_method = :encode
      self.unmarshal_class_method = :decode
      self.service_name = 'gitaly.CleanupService'

      # This comment is left unintentionally blank.
      rpc :ApplyBfgObjectMapStream, stream(::Gitaly::ApplyBfgObjectMapStreamRequest), stream(::Gitaly::ApplyBfgObjectMapStreamResponse)
    end

    Stub = Service.rpc_stub_class
  end
end
