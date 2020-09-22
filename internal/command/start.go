package command

import (
	"github.com/lucasvmiguel/task/internal/gitrepo"

	"github.com/davecgh/go-spew/spew"
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
	c.logger.Debugf("start params: %v", spew.Sdump(params))

	if params.ID == "" {
		return errors.New("task ID cannot be empty")
	}

	issue, err := c.issueTracker.Issue(params.ID)
	if err != nil {
		return errors.Wrap(err, "failed to fetch issue")
	}
	issue.ID = params.ID

	origin, err := c.versionControl.Origin()
	if err != nil {
		return errors.Wrap(err, "failed to get origin from the repository")
	}

	err = c.versionControl.CreateBranchAndPush(issue.ID)
	if err != nil {
		return errors.Wrap(err, "failed to create branch, branch may exist already")
	}

	newPR := gitrepo.NewPR{
		Branch:      issue.ID,
		Title:       replaceInTemplate(params.TitleTemplate, issue),
		Description: replaceInTemplate(params.DescriptionTemplate, issue),
		Org:         origin.Org,
		Repository:  origin.Repository,
	}

	c.logger.Debugf("start params: %v", spew.Sdump(newPR))
	prLink, err := c.gitRepo.CreatePR(newPR)
	if err != nil {
		return errors.Wrap(err, "failed to create pull request")
	}

	c.logger.Infof("PR created, here is the link: %s", prLink)

	return nil
}
