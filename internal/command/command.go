package command

import (
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

// Config is a struct passed as param to create a new Command struct
type Config struct {
	IssueTracker   issuetracker.IssueTracker
	GitRepo        gitrepo.GitRepo
	VersionControl versioncontrol.VersionControl
}

// New is a function to create a new Command struct
func New(config Config) (Command, error) {
	return Command{
		IssueTracker:   config.IssueTracker,
		GitRepo:        config.GitRepo,
		VersionControl: config.VersionControl,
	}, nil
}
