package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/lucasvmiguel/task/internal/issuetracker"
	"github.com/pkg/errors"
)

type Client struct {
	Host     string
	Username string
	Key      string
	client   *jira.Client
}

func (c *Client) Authenticate() error {
	tp := jira.BasicAuthTransport{
		Username: c.Username,
		Password: c.Key,
	}

	client, err := jira.NewClient(tp.Client(), c.Host)
	if err != nil {
		return errors.Wrap(err, "failed to authenticate with jira")
	}

	c.client = client
	return nil
}

func (c *Client) Issue(ID string) (*issuetracker.Issue, error) {
	issue, _, err := c.client.Issue.Get(ID, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch issue on jira")
	}

	return &issuetracker.Issue{
		ID:          issue.ID,
		Title:       issue.Fields.Summary,
		Description: issue.Fields.Description,
	}, nil
}
