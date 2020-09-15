package gitrepo

// GitRepo is a interface to represent any git repository (eg: github, gitlab, bitbucket)
type GitRepo interface {
	// Authenticate to a git repository server
	Authenticate(host, token string) error
	// CreatePR in a git repository server
	CreatePR(pr NewPR) (string, error)
}

// NewPR is a struct used to create a new pull request in a git repository
type NewPR struct {
	Branch      string
	Title       string
	Description string
	Org         string
	Repository  string
}
