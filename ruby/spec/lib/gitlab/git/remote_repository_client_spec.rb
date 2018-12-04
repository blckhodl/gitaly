require 'spec_helper'

describe Gitlab::Git::GitalyRemoteRepository do
  include TestRepo
  include IntegrationClient

  let(:repository) { gitlab_git_from_gitaly_with_gitlab_projects(new_mutable_test_repo) }
  describe 'certs' do
    let(:client) { get_client("tls://localhost:#{GitalyConfig.dynamic_port('tls')}") }

    context 'when neither SSL_CERT_FILE and SSL_CERT_DIR is set' do
      it 'Raises an error' do
        expect { client.certs }.to raise_error 'SSL_CERT_DIR and/or SSL_CERT_FILE environment variable must be set'
      end
    end

    context 'when SSL_CERT_FILE is set' do
      it 'Should return the correct certificate' do
        cert = File.join(File.dirname(__FILE__), "testdata/certs/gitalycert.pem")
        allow(ENV).to receive(:[]).with('SSL_CERT_DIR').and_return(nil)
        allow(ENV).to receive(:[]).with('SSL_CERT_FILE').and_return(cert)
        certs = client.certs
        expect(certs).to eq File.read(cert)
      end
    end

    context 'when SSL_CERT_DIR is set' do
      it 'Should return concatenation of gitalycert and gitalycert2' do
        cert_pool_dir = File.join(File.dirname(__FILE__), "testdata/certs")
        allow(ENV).to receive(:[]).with('SSL_CERT_DIR').and_return(cert_pool_dir)
        allow(ENV).to receive(:[]).with('SSL_CERT_FILE').and_return(nil)
        certs = client.certs
        expected_certs = [File.read(File.join(cert_pool_dir, "gitalycert2.pem")), File.read(File.join(cert_pool_dir, "gitalycert.pem"))].join "\n"
        expect(certs).to eq expected_certs
      end
    end

    context 'when both SSL_CERT_DIR and SSL_CERT_FILE are set' do
      it 'Should return all certs in SSL_CERT_DIR + SSL_CERT_FILE' do
        cert_pool_dir = File.join(File.dirname(__FILE__), "testdata/certs")
        cert1_file = File.join(File.dirname(__FILE__), "testdata/gitalycert.pem")
        allow(ENV).to receive(:[]).with('SSL_CERT_DIR').and_return(cert_pool_dir)
        allow(ENV).to receive(:[]).with('SSL_CERT_FILE').and_return(cert1_file)
        expected_certs_paths = [File.join(cert_pool_dir, "gitalycert2.pem"), File.join(cert_pool_dir, "gitalycert.pem"), cert1_file]

        expected_certs = expected_certs_paths.map do |cert|
          File.read cert
        end.join("\n")
        certs = client.certs
        expect(certs).to eq expected_certs
      end
    end
  end

  describe 'Connectivity' do
    context 'tcp' do
      let(:client) do
        get_client("tcp://localhost:#{GitalyConfig.dynamic_port('tcp')}")
      end

      it 'Should connect over tcp' do
        expect(client).not_to be_empty
      end
    end

    context 'unix' do
      let(:client) { get_client("unix:#{File.join(TMP_DIR_NAME, SOCKET_PATH)}") }

      it 'Should connect over unix' do
        expect(client).not_to be_empty
      end
    end

    context 'tls' do
      let(:client) { get_client("tls://localhost:#{GitalyConfig.dynamic_port('tls')}") }

      it 'Should connect over tls' do
        cert = File.join(File.dirname(__FILE__), "testdata/certs/gitalycert.pem")
        allow(ENV).to receive(:[]).with('SSL_CERT_DIR').and_return(nil)
        allow(ENV).to receive(:[]).with('SSL_CERT_FILE').and_return(cert)

        expect(client).not_to be_empty
      end
    end
  end
end
