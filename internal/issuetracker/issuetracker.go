// Package issuetracker provides all the required structs and
// interfaces to use a issue tracker (eg: trello, jira, etc)
package issuetracker

// IssueTracker is a interface to represent any issue tracker (eg: trello, jira, etc)
type IssueTracker interface {
	// Authenticate to a issue tracker server
	Authenticate(host, username, key string) error
	// Issue fetches an issue on a issue tracker server
	Issue(ID string) (*Issue, error)
}

// Issue is a struct that contains all relevant issue fields on a issue tracker
type Issue struct {
	ID          string
	Title       string
	Description string
	Link        string
}
