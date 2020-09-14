package command

import (
	"github.com/lucasvmiguel/task/internal/gitrepo"
	"github.com/lucasvmiguel/task/internal/issuetracker"
	"github.com/lucasvmiguel/task/internal/versioncontrol"
)

type Command struct {
	IssueTracker   issuetracker.IssueTracker
	GitRepo        gitrepo.GitRepo
	VersionControl versioncontrol.VersionControl
}

type Config struct {
	IssueTracker   issuetracker.IssueTracker
	GitRepo        gitrepo.GitRepo
	VersionControl versioncontrol.VersionControl
}

func New(config Config) (Command, error) {
	return Command{
		IssueTracker:   config.IssueTracker,
		GitRepo:        config.GitRepo,
		VersionControl: config.VersionControl,
	}, nil
}
