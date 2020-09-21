// Package git that implements VersionControl interface
package git

import (
	"os/exec"
	"strings"

	"github.com/lucasvmiguel/task/internal/versioncontrol"
	"github.com/pkg/errors"
)

// Client to use git commands
type Client struct{}

// CreateBranchAndPush uses git create a local branch and push it to remote
// is going to fail in case it cannot be created or pushed
func (c *Client) CreateBranchAndPush(name string) error {
	if name == "" {
		return errors.New("name cannot be empty")
	}

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

// Origin gets origin repository information like organization and repository
func (c *Client) Origin() (*versioncontrol.Origin, error) {
	out, err := exec.Command("git", "remote", "get-url", "origin").Output()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get origin from command line")
	}

	return splitOriginURL(string(out))
}

func splitOriginURL(originURL string) (*versioncontrol.Origin, error) {
	// git@github.com:lucasvmiguel/task.git
	urlSplitted := strings.Split(originURL, ":")
	if len(urlSplitted) != 2 {
		return nil, errors.New("invalid remote url")
	}

	// lucasvmiguel/task.git
	orgAndRepoSplitted := strings.Split(string(urlSplitted[1]), "/")
	if len(orgAndRepoSplitted) != 2 {
		return nil, errors.New("invalid remote org and repository")
	}

	r := strings.NewReplacer(
		".git", "",
		"\n", "",
	)

	return &versioncontrol.Origin{
		// lucasvmiguel
		Org: orgAndRepoSplitted[0],
		// task
		Repository: r.Replace(orgAndRepoSplitted[1]),
	}, nil
}
