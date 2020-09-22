package command

import (
	"errors"
	"testing"

	"github.com/lucasvmiguel/task/internal/gitrepo/github"
	"github.com/lucasvmiguel/task/internal/issuetracker/jira"
	"github.com/lucasvmiguel/task/internal/log/terminal"
	"github.com/lucasvmiguel/task/internal/versioncontrol/git"

	"github.com/go-test/deep"
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
				Logger:         &terminal.Logger{},
			},
			&Command{
				issueTracker:   &jira.Client{},
				gitRepo:        &github.Client{},
				versionControl: &git.Client{},
				logger:         &terminal.Logger{},
			},
			nil,
		},
		{
			NewParams{
				IssueTracker:   nil,
				GitRepo:        &github.Client{},
				VersionControl: &git.Client{},
				Logger:         &terminal.Logger{},
			},
			nil,
			errors.New("issue tracker cannot be nil"),
		},
		{
			NewParams{
				IssueTracker:   &jira.Client{},
				GitRepo:        nil,
				VersionControl: &git.Client{},
				Logger:         &terminal.Logger{},
			},
			nil,
			errors.New("git repo cannot be nil"),
		},
		{
			NewParams{
				IssueTracker:   &jira.Client{},
				GitRepo:        &github.Client{},
				VersionControl: nil,
				Logger:         &terminal.Logger{},
			},
			nil,
			errors.New("version control cannot be nil"),
		},
		{
			NewParams{
				IssueTracker:   &jira.Client{},
				GitRepo:        &github.Client{},
				VersionControl: &git.Client{},
				Logger:         nil,
			},
			nil,
			errors.New("logger cannot be nil"),
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

			continue
		}

		if diff := deep.Equal(err, tt.expectedError); diff != nil {
			t.Error(diff)
		}
	}
}
