package issuetracker

type IssueTracker interface {
	Authenticate() error
	Issue(ID string) (*Issue, error)
}

type Issue struct {
	ID          string
	Title       string
	Description string
}
