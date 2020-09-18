package command

import (
	"fmt"

	"github.com/lucasvmiguel/task/internal/gitrepo"

	"github.com/pkg/errors"
)

// StartParams is a struct passed as param to run the start command
type StartParams struct {
	ID                  string
	TitleTemplate       string
	DescriptionTemplate string
}

// Start command flow:
// 1. fetch issue on a issue tracker
// 2. creates and pushes a branch to a git repository
// 3. creates a PR on a git repository
// 4. returns the PR link
func (c *Command) Start(params StartParams) error {
	issue, err := c.IssueTracker.Issue(params.ID)
	if err != nil {
		return errors.Wrap(err, "failed to fetch issue")
	}
	issue.ID = params.ID

	origin, err := c.VersionControl.Origin()
	if err != nil {
		return errors.Wrap(err, "failed to get origin from the repository")
	}

	err = c.VersionControl.CreateBranchAndPush(issue.ID)
	if err != nil {
		return errors.Wrap(err, "failed to create branch, branch may exist already")
	}

	prLink, err := c.GitRepo.CreatePR(gitrepo.NewPR{
		Branch:      issue.ID,
		Title:       replaceInTemplate(params.TitleTemplate, issue),
		Description: replaceInTemplate(params.DescriptionTemplate, issue),
		Org:         origin.Org,
		Repository:  origin.Repository,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create pull request")
	}

	fmt.Println("here is the PR: " + prLink)

	return nil
}
