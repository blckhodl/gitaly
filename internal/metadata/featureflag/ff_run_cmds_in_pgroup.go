package featureflag

var RunCmdsInProcessGroup = NewFeatureFlag(
	"run_cmds_in_process_group",
	"v15.5.0",
	"https://gitlab.com/gitlab-org/gitaly/-/issues/4494",
	false,
)
