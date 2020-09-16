package command

import (
	"errors"
	"testing"

	"github.com/go-test/deep"
	"github.com/lucasvmiguel/task/internal/gitrepo/github"
	"github.com/lucasvmiguel/task/internal/issuetracker/jira"
	"github.com/lucasvmiguel/task/internal/versioncontrol/git"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		params        NewParams
		expected      *Command
		expectedError error
	}{
		{
			NewParams{
				IssueTracker:   &jira.Client{},
				GitRepo:        &github.Client{},
				VersionControl: &git.Client{},
			},
			&Command{
				IssueTracker:   &jira.Client{},
				GitRepo:        &github.Client{},
				VersionControl: &git.Client{},
			},
			nil,
		},
		{
			NewParams{
				IssueTracker:   nil,
				GitRepo:        &github.Client{},
				VersionControl: &git.Client{},
			},
			nil,
			errors.New("issue tracker cannot be nil"),
		},
		{
			NewParams{
				IssueTracker:   &jira.Client{},
				GitRepo:        nil,
				VersionControl: &git.Client{},
			},
			nil,
			errors.New("git repo cannot be nil"),
		},
		{
			NewParams{
				IssueTracker:   &jira.Client{},
				GitRepo:        &github.Client{},
				VersionControl: nil,
			},
			nil,
			errors.New("version control cannot be nil"),
		},
	}

	for _, tt := range tests {
		result, err := New(tt.params)

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
