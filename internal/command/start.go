package command

import (
	"fmt"

	"github.com/lucasvmiguel/task/internal/gitrepo"

	"github.com/pkg/errors"
)

// Start command flow:
// 1. fetch issue on a issue tracker
// 2. creates and pushes a branch to a git repository
// 3. creates a PR on a git repository
// 4. returns the PR link
func (c *Command) Start(ID, org, repo string) error {
	issue, err := c.IssueTracker.Issue(ID)
	if err != nil {
		return errors.Wrap(err, "failed to fetch issue")
	}

	err = c.VersionControl.CreateBranchAndPush(ID)
	if err != nil {
		return errors.Wrap(err, "failed to create branch, branch may exist already")
	}

	prLink, err := c.GitRepo.CreatePR(gitrepo.NewPR{
		Branch:      ID,
		Title:       issue.Title,
		Description: issue.Description,
		Org:         org,
		Repository:  repo,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create pull request")
	}

	fmt.Println("here is the PR: " + prLink)

	return nil
}
