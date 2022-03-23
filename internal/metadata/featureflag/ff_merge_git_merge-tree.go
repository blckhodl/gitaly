package featureflag

// MergeGitMergeTree enables implementation of UserMergeBranch using
// `git merge-tree` instead of s.git2goExecutor.Merge()
var MergeGitMergeTree = NewFeatureFlag("merge_git_merge-tree", false)
