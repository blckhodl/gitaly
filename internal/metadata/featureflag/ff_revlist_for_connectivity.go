package featureflag

var RevlistForConnectivity = NewFeatureFlag(
	"revlist_for_connectivity",
	"v15.4.0",
	"https://gitlab.com/gitlab-org/gitaly/-/issues/4489",
	false,
)
