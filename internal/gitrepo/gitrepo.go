package gitrepo

type GitRepo interface {
	Authenticate() error
	CreatePR(branch, title, description string) error
}
