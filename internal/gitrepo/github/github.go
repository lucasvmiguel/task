// Package github that implements GitRepo interface
package github

import (
	"context"

	"github.com/lucasvmiguel/task/internal/gitrepo"

	"github.com/google/go-github/v32/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

// Client to communicate with Github
type Client struct {
	client *github.Client
}

// Authenticate to github
func (c *Client) Authenticate(host, token string) error {
	if token == "" {
		return errors.New("token cannot be empty")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	var client *github.Client
	var err error
	if host == "" {
		client = github.NewClient(tc)
	} else {
		client, err = github.NewEnterpriseClient(host, host, nil)
		if err != nil {
			return errors.Wrap(err, "failed to authenticate with enterprise github")
		}
	}

	c.client = client
	return nil
}

// CreatePR creates a pull request on github
func (c *Client) CreatePR(newPR gitrepo.NewPR) (string, error) {
	if newPR.Title == "" || newPR.Branch == "" || newPR.Org == "" || newPR.Repository == "" {
		return "", errors.New("title, branch, org and repository cannot be empty")
	}

	pr := &github.NewPullRequest{
		Title: github.String(newPR.Title),
		Base:  github.String("master"),
		Head:  github.String(newPR.Branch),
		Body:  github.String(newPR.Description),
	}

	p, _, err := c.client.PullRequests.Create(context.Background(), newPR.Org, newPR.Repository, pr)
	if err != nil {
		return "", errors.Wrap(err, "failed to create PR on github")
	}

	return p.GetHTMLURL(), nil
}
