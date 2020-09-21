// Package jira that implements IssueTracker interface
package jira

import (
	"github.com/lucasvmiguel/task/internal/issuetracker"

	"github.com/andygrunwald/go-jira"
	"github.com/pkg/errors"
)

// Client to communicate with Jira
type Client struct {
	client *jira.Client
}

// Authenticate to a Jira server
func (c *Client) Authenticate(host, username, key string) error {
	if username == "" || key == "" {
		return errors.New("username and key cannot be empty")
	}

	tp := jira.BasicAuthTransport{
		Username: username,
		Password: key,
	}

	client, err := jira.NewClient(tp.Client(), host)
	if err != nil {
		return errors.Wrap(err, "failed to authenticate with jira")
	}

	c.client = client
	return nil
}

// Issue fetches an issue on Jira by the issue ID
func (c *Client) Issue(ID string) (*issuetracker.Issue, error) {
	if ID == "" {
		return nil, errors.New("ID cannot be empty")
	}

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
