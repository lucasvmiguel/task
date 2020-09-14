package git

import (
	"os/exec"

	"github.com/pkg/errors"
)

type Client struct {
}

func (c *Client) CreateBranch(name string) error {
	err := exec.Command("git checkout -b " + name).Run()
	if err != nil {
		return errors.Wrap(err, "failed to create a git branch")
	}
	err = exec.Command("git commit --allow-empty -m \"First commit for \"").Run()
	if err != nil {
		return errors.Wrap(err, "failed to create first commit")
	}
	err = exec.Command("git push origin " + name).Run()
	if err != nil {
		return errors.Wrap(err, "failed to push branch")
	}

	return nil
}
