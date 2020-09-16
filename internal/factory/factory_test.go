package factory

import (
	"testing"

	"github.com/lucasvmiguel/task/internal/gitrepo"
	"github.com/lucasvmiguel/task/internal/gitrepo/github"
	"github.com/lucasvmiguel/task/internal/issuetracker"
	"github.com/lucasvmiguel/task/internal/issuetracker/jira"
	"github.com/pkg/errors"

	"github.com/go-test/deep"
)

func TestGitRepo(t *testing.T) {
	var tests = []struct {
		params        GitRepoParams
		expected      gitrepo.GitRepo
		expectedError error
	}{
		{
			GitRepoParams{Provider: Github, Host: "test", Token: "test"},
			&github.Client{},
			nil,
		},
		{
			GitRepoParams{Provider: "invalid", Host: "test", Token: "test"},
			nil,
			errors.New("invalid git repository"),
		},
	}

	for _, tt := range tests {
		result, err := GitRepo(tt.params)

		if tt.expectedError == nil {
			diff := deep.Equal(result, tt.expected)
			if diff != nil {
				t.Error(diff)
			}

			if err != nil {
				t.Error(err)
			}

			return
		}

		if diff := deep.Equal(err, tt.expectedError); diff != nil {
			t.Error(diff)
		}
	}
}

func TestIssueTracker(t *testing.T) {
	var tests = []struct {
		params        IssueTrackerParams
		expected      issuetracker.IssueTracker
		expectedError error
	}{
		{
			IssueTrackerParams{Provider: Jira, Host: "test", Username: "test", Key: "test"},
			&jira.Client{},
			nil,
		},
		{
			IssueTrackerParams{Provider: "invalid", Host: "test", Username: "test", Key: "test"},
			nil,
			errors.New("invalid git repository"),
		},
	}

	for _, tt := range tests {
		result, err := IssueTracker(tt.params)

		if tt.expectedError == nil {
			diff := deep.Equal(result, tt.expected)
			if diff != nil {
				t.Error(diff)
			}

			if err != nil {
				t.Error(err)
			}

			return
		}

		if diff := deep.Equal(err, tt.expectedError); diff != nil {
			t.Error(diff)
		}
	}
}
