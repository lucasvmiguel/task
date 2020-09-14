package command

import (
	"github.com/pkg/errors"
)

func (c *Command) Start(ID string) error {
	err := c.IssueTracker.Authenticate()
	if err != nil {
		return errors.Wrap(err, "failed to authenticate with issue tracker")
	}
	err = c.GitRepo.Authenticate()
	if err != nil {
		return errors.Wrap(err, "failed to authenticate with git repository")
	}
	issue, err := c.IssueTracker.Issue(ID)
	if err != nil {
		return errors.Wrap(err, "failed to fetch issue")
	}
	err = c.VersionControl.CreateBranch(issue.ID)
	if err != nil {
		return errors.Wrap(err, "failed to create branch")
	}
	err = c.GitRepo.CreatePR(issue.ID, issue.Title, issue.Description)
	if err != nil {
		return errors.Wrap(err, "failed to create pull request")
	}

	return nil
}
