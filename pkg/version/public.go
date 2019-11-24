package version

func buildVersion() string {
	if Metadata == "" {
		return Version
	}
	return Version + "-" + Metadata
}

// GetVersion is here to get version of the cli
func GetVersion() *AppVersion {
	return &AppVersion{
		Version:   buildVersion(),
		GitCommit: GitCommit,
		BuildDate: BuildDate,
	}
}
