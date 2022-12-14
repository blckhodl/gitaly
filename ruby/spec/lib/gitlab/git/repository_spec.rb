require 'spec_helper'

describe Gitlab::Git::Repository do # rubocop:disable Metrics/BlockLength
  include TestRepo
  include Gitlab::EncodingHelper
  using RSpec::Parameterized::TableSyntax

  let(:mutable_repository) { gitlab_git_from_gitaly(new_mutable_git_test_repo) }
  let(:repository) { gitlab_git_from_gitaly(git_test_repo_read_only) }
  let(:repository_path) { repository.path }
  let(:repository_rugged) { Rugged::Repository.new(repository_path) }
  let(:storage_path) { DEFAULT_STORAGE_DIR }
  let(:user) { Gitlab::Git::User.new('johndone', 'John Doe', 'johndoe@mail.com', 'user-1') }

  describe '.from_gitaly_with_block' do
    let(:call_metadata) do
      {
        'user-agent' => 'grpc-go/1.9.1',
        'gitaly-storage-path' => DEFAULT_STORAGE_DIR,
        'gitaly-repo-path' => TEST_REPO_PATH,
        'gitaly-gl-repository' => 'project-52',
        'gitaly-repo-alt-dirs' => ''
      }
    end
    let(:call) { double(metadata: call_metadata) }

    it 'cleans up the repository' do
      described_class.from_gitaly_with_block(test_repo_read_only, call) do |repository|
        expect(repository.rugged).to receive(:close)
      end
    end

    it 'returns the passed result of the block passed' do
      result = described_class.from_gitaly_with_block(test_repo_read_only, call) { 'Hello world' }

      expect(result).to eq('Hello world')
    end
  end

  describe "Respond to" do
    subject { repository }

    it { is_expected.to respond_to(:root_ref) }
    it { is_expected.to respond_to(:tags) }
  end

  describe '#root_ref' do
    it 'calls #discover_default_branch' do
      expect(repository).to receive(:discover_default_branch)
      repository.root_ref
    end
  end

  describe '#branch_names' do
    subject { repository.branch_names }

    it 'has SeedRepo::Repo::BRANCHES.size elements' do
      expect(subject.size).to eq(SeedRepo::Repo::BRANCHES.size)
    end

    it { is_expected.to include("master") }
    it { is_expected.not_to include("branch-from-space") }
  end

  describe '#tags' do
    describe 'first tag' do
      let(:tag) { repository.tags.first }

      it { expect(tag.name).to eq("v1.0.0") }
      it { expect(tag.target).to eq("f4e6814c3e4e7a0de82a9e7cd20c626cc963a2f8") }
      it { expect(tag.dereferenced_target.sha).to eq("6f6d7e7ed97bb5f0054f2b1df789b39ca89b6ff9") }
      it { expect(tag.message).to eq("Release") }
    end

    describe 'last tag' do
      let(:tag) { repository.tags.last }

      it { expect(tag.name).to eq("v1.2.1") }
      it { expect(tag.target).to eq("2ac1f24e253e08135507d0830508febaaccf02ee") }
      it { expect(tag.dereferenced_target.sha).to eq("fa1b1e6c004a68b7d8763b86455da9e6b23e36d6") }
      it { expect(tag.message).to eq("Version 1.2.1") }
    end

    it { expect(repository.tags.size).to eq(SeedRepo::Repo::TAGS.size) }
  end

  describe '#empty?' do
    it { expect(repository).not_to be_empty }
  end

  describe '#merge_base' do
    where(:from, :to, :result) do
      '570e7b2abdd848b95f2f578043fc23bd6f6fd24d' | '40f4a7a617393735a95a0bb67b08385bc1e7c66d' | '570e7b2abdd848b95f2f578043fc23bd6f6fd24d'
      '40f4a7a617393735a95a0bb67b08385bc1e7c66d' | '570e7b2abdd848b95f2f578043fc23bd6f6fd24d' | '570e7b2abdd848b95f2f578043fc23bd6f6fd24d'
      '40f4a7a617393735a95a0bb67b08385bc1e7c66d' | 'foobar' | nil
      'foobar' | '40f4a7a617393735a95a0bb67b08385bc1e7c66d' | nil
    end

    with_them do
      it { expect(repository.merge_base(from, to)).to eq(result) }
    end
  end

  describe '#find_branch' do
    it 'should return a Branch for master' do
      branch = repository.find_branch('master')

      expect(branch).to be_a_kind_of(Gitlab::Git::Branch)
      expect(branch.name).to eq('master')
    end

    it 'should handle non-existent branch' do
      branch = repository.find_branch('this-is-garbage')

      expect(branch).to eq(nil)
    end
  end

  describe '#branches' do
    subject { repository.branches }

    context 'with local and remote branches' do
      let(:repository) { mutable_repository }

      before do
        create_remote_branch('joe', 'remote_branch', 'master')
        create_branch(repository, 'local_branch', 'master')
      end

      it 'returns the local and remote branches' do
        expect(subject.any? { |b| b.name == 'joe/remote_branch' }).to eq(true)
        expect(subject.any? { |b| b.name == 'local_branch' }).to eq(true)
      end
    end
  end

  describe '#branch_exists?' do
    it 'returns true for an existing branch' do
      expect(repository.branch_exists?('master')).to eq(true)
    end

    it 'returns false for a non-existing branch' do
      expect(repository.branch_exists?('kittens')).to eq(false)
    end

    it 'returns false when using an invalid branch name' do
      expect(repository.branch_exists?('.bla')).to eq(false)
    end
  end

  describe '#with_repo_branch_commit' do
    context 'when repository is empty' do
      let(:repository) { gitlab_git_from_gitaly(new_empty_test_repo) }

      it 'yields nil' do
        expect do |block|
          repository.with_repo_branch_commit('master', &block)
        end.to yield_with_args(nil)
      end
    end

    context 'when repository is not empty' do
      let(:start_commit) { repository.commit }

      it 'yields the commit for the SHA' do
        expect do |block|
          repository.with_repo_branch_commit(start_commit.sha, &block)
        end.to yield_with_args(start_commit)
      end

      it 'yields the commit for the branch' do
        expect do |block|
          repository.with_repo_branch_commit('master', &block)
        end.to yield_with_args(start_commit)
      end
    end
  end

  describe '#cleanup' do
    context 'when Rugged has been called' do
      it 'calls close on Rugged::Repository' do
        rugged = repository.rugged

        expect(rugged).to receive(:close).and_call_original

        repository.cleanup
      end
    end

    context 'when Rugged has not been called' do
      it 'does not call close on Rugged::Repository' do
        expect(repository).not_to receive(:rugged)

        repository.cleanup
      end
    end
  end

  describe '#rugged' do
    after do
      Thread.current[described_class::RUGGED_KEY] = nil
    end

    it 'stores reference in Thread.current' do
      Thread.current[described_class::RUGGED_KEY] = []

      2.times do
        rugged = repository.rugged

        expect(rugged).to be_a(Rugged::Repository)
        expect(Thread.current[described_class::RUGGED_KEY]).to eq([rugged])
      end
    end

    it 'does not store reference if Thread.current is not set up' do
      rugged = repository.rugged

      expect(rugged).to be_a(Rugged::Repository)
      expect(Thread.current[described_class::RUGGED_KEY]).to be_nil
    end
  end

  describe '#head_symbolic_ref' do
    subject { repository.head_symbolic_ref }

    it 'returns the symbolic ref in HEAD' do
      expect(subject).to eq('master')
    end

    context 'when repo is empty' do
      let(:repository) { gitlab_git_from_gitaly(new_empty_test_repo) }

      it 'returns the symbolic ref in HEAD' do
        repository.rugged.head = 'refs/heads/foo'

        expect(subject).to eq('foo')
      end
    end
  end

  def create_remote_branch(remote_name, branch_name, source_branch_name)
    source_branch = repository.branches.find { |branch| branch.name == source_branch_name }
    repository_rugged.references.create("refs/remotes/#{remote_name}/#{branch_name}", source_branch.dereferenced_target.sha)
  end

  def create_branch(repository, branch_name, start_point = 'HEAD')
    repository.rugged.branches.create(branch_name, start_point)
  end
end
