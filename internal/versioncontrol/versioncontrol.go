package versioncontrol

type VersionControl interface {
	CreateBranch(name string) error
}
