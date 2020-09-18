// Package command is responsible for executing all the CLI commands (eg: start)
package command

import (
	"errors"
	"strings"

	"github.com/lucasvmiguel/task/internal/gitrepo"
	"github.com/lucasvmiguel/task/internal/issuetracker"
	"github.com/lucasvmiguel/task/internal/versioncontrol"
)

// Command struct is responsable for execute CLI commands
type Command struct {
	IssueTracker   issuetracker.IssueTracker
	GitRepo        gitrepo.GitRepo
	VersionControl versioncontrol.VersionControl
}

// NewParams is a struct passed as param to create a new Command struct
type NewParams struct {
	IssueTracker   issuetracker.IssueTracker
	GitRepo        gitrepo.GitRepo
	VersionControl versioncontrol.VersionControl
}

// New is a function to create a new Command struct
func New(params NewParams) (*Command, error) {
	if params.IssueTracker == nil {
		return nil, errors.New("issue tracker cannot be nil")
	}

	if params.GitRepo == nil {
		return nil, errors.New("git repo cannot be nil")
	}

	if params.VersionControl == nil {
		return nil, errors.New("version control cannot be nil")
	}

	return &Command{
		IssueTracker:   params.IssueTracker,
		GitRepo:        params.GitRepo,
		VersionControl: params.VersionControl,
	}, nil
}

func replaceInTemplate(text string, issue *issuetracker.Issue) string {
	r := strings.NewReplacer(
		"{{ISSUE_TRACKER.ID}}", issue.ID,
		"{{ISSUE_TRACKER.TITLE}}", issue.Title,
		"{{ISSUE_TRACKER.DESCRIPTION}}", issue.Description,
	)

	return r.Replace(text)
}
