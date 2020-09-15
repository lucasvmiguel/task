package versioncontrol

// VersionControl is an interface to represent any version control like git
type VersionControl interface {
	CreateBranchAndPush(name string) error
}
