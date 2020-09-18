// Package versioncontrol provides all the required structs and
// interfaces to use a version control (eg: git)
package versioncontrol

// VersionControl is an interface to represent any version control like git
type VersionControl interface {
	CreateBranchAndPush(name string) error
	Origin() (*Origin, error)
}

// Origin info about a version control
type Origin struct {
	Org        string
	Repository string
}
