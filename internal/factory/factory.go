// Package factory provides the factory method pattern for different structs in the project
// reference: https://en.wikipedia.org/wiki/Factory_method_pattern
package factory

import (
	"github.com/lucasvmiguel/task/internal/gitrepo"
	"github.com/lucasvmiguel/task/internal/gitrepo/github"
	"github.com/lucasvmiguel/task/internal/issuetracker"
	"github.com/lucasvmiguel/task/internal/issuetracker/jira"
	"github.com/pkg/errors"
)

// GitRepoProvider string type
type GitRepoProvider string

// IssueTrackerProvider string type
type IssueTrackerProvider string

const (
	// Github GitRepoProvider
	Github GitRepoProvider = "github"

	// Jira IssueTrackerProvider
	Jira IssueTrackerProvider = "jira"
)

// GitRepoParams struct to be used in the function GitRepo
type GitRepoParams struct {
	Provider GitRepoProvider
	Host     string
	Token    string
}

// IssueTrackerParams struct to be used in the function IssueTracker
type IssueTrackerParams struct {
	Provider IssueTrackerProvider
	Host     string
	Username string
	Key      string
}

// GitRepo returns a struct that implements GitRepo interface
func GitRepo(params GitRepoParams) (gitrepo.GitRepo, error) {
	var gitRepo gitrepo.GitRepo
	switch params.Provider {
	case Github:
		gitRepo = &github.Client{}
	default:
		return nil, errors.New("invalid git repository")
	}

	err := gitRepo.Authenticate(params.Host, params.Token)
	if err != nil {
		return nil, errors.Wrap(err, "failed to authenticate to git repository")
	}

	return gitRepo, nil
}

// IssueTracker returns a struct that implements IssueTracker interface
func IssueTracker(params IssueTrackerParams) (issuetracker.IssueTracker, error) {
	var issueTracker issuetracker.IssueTracker
	switch params.Provider {
	case "jira":
		issueTracker = &jira.Client{}
	default:
		return nil, errors.New("invalid issue tracker")
	}

	err := issueTracker.Authenticate(params.Host, params.Username, params.Key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to authenticate to issue tracker")
	}

	return issueTracker, nil
}
