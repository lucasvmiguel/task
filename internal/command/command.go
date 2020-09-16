// Package command is responsible for executing all the CLI commands (eg: start)
package command

import (
	"errors"

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
func New(config NewParams) (*Command, error) {
	if config.IssueTracker == nil {
		return nil, errors.New("issue tracker cannot be nil")
	}

	if config.GitRepo == nil {
		return nil, errors.New("git repo cannot be nil")
	}

	if config.VersionControl == nil {
		return nil, errors.New("version control cannot be nil")
	}

	return &Command{
		IssueTracker:   config.IssueTracker,
		GitRepo:        config.GitRepo,
		VersionControl: config.VersionControl,
	}, nil
}
