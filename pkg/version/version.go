package version

var (
	// The variables are set during build time using -ldflags
	Version   = "dev"
	CommitSHA = "none"
	BuildTime = "unknown"
)

type Info struct {
	Version   string `json:"version"`
	CommitSHA string `json:"commit_sha"`
	BuildTime string `json:"build_time"`
}

func Get() Info {
	return Info{
		Version:   Version,
		CommitSHA: CommitSHA,
		BuildTime: BuildTime,
	}
}
