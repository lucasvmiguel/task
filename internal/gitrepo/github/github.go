package github

import (
	"context"

	"github.com/google/go-github/v32/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

type Client struct {
	Host   string
	Token  string
	Org    string
	Repo   string
	client *github.Client
}

func (c *Client) Authenticate() error {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: c.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	var client *github.Client
	var err error
	if c.Host == "" {
		client = github.NewClient(tc)
	} else {
		client, err = github.NewEnterpriseClient(c.Host, c.Host, nil)
		if err != nil {
			return errors.Wrap(err, "failed to authenticate with enterprise github")
		}
	}

	c.client = client
	return nil
}

func (c *Client) CreatePR(branch, title, description string) error {
	base := "master"

	newPR := &github.NewPullRequest{
		Title: &title,
		Base:  &base,
		Head:  &branch,
		Body:  &description,
	}

	_, _, err := c.client.PullRequests.Create(context.Background(), c.Org, c.Repo, newPR)
	if err != nil {
		return errors.Wrap(err, "failed to authenticate with jira")
	}
	return nil
}
