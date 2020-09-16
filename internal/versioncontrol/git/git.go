// Package git that implements VersionControl interface
package git

import (
	"os/exec"

	"github.com/pkg/errors"
)

// Client to use git commands
type Client struct{}

// CreateBranchAndPush uses git create a local branch and push it to remote
// is going to fail in case it cannot be created or pushed
func (c *Client) CreateBranchAndPush(name string) error {
	err := exec.Command("git", "checkout", "-b", name).Run()
	if err != nil {
		return errors.Wrap(err, "failed to create a git branch")
	}

	err = exec.Command("git", "commit", "--allow-empty", "-m", "First commit for").Run()
	if err != nil {
		return errors.Wrap(err, "failed to create first commit")
	}

	err = exec.Command("git", "push", "origin", name).Run()
	if err != nil {
		return errors.Wrap(err, "failed to push branch")
	}

	return nil
}
