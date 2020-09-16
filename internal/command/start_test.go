package command

import (
	"errors"
	"testing"

	"github.com/go-test/deep"

	"github.com/lucasvmiguel/task/internal/gitrepo"
	"github.com/lucasvmiguel/task/internal/issuetracker"
)

type IssueTrackerMock struct{}

func (i *IssueTrackerMock) Authenticate(host, username, key string) error {
	return nil
}

func (i *IssueTrackerMock) Issue(ID string) (*issuetracker.Issue, error) {
	return &issuetracker.Issue{ID: "test", Title: "test", Description: "test"}, nil
}

type IssueTrackerFailedMock struct{}

func (i *IssueTrackerFailedMock) Authenticate(host, username, key string) error {
	return nil
}

func (i *IssueTrackerFailedMock) Issue(ID string) (*issuetracker.Issue, error) {
	return nil, errors.New("error fetching issue")
}

type GitRepoMock struct{}

func (i *GitRepoMock) Authenticate(host, token string) error {
	return nil
}

func (i *GitRepoMock) CreatePR(pr gitrepo.NewPR) (string, error) {
	return "pr link", nil
}

type GitRepoFailedMock struct{}

func (i *GitRepoFailedMock) Authenticate(host, token string) error {
	return nil
}

func (i *GitRepoFailedMock) CreatePR(pr gitrepo.NewPR) (string, error) {
	return "", errors.New("error creating pr")
}

type VersionControlMock struct{}

func (i *VersionControlMock) CreateBranchAndPush(name string) error {
	return nil
}

type VersionControlFailedMock struct{}

func (i *VersionControlFailedMock) CreateBranchAndPush(name string) error {
	return errors.New("error creating or pushing branch")
}

func TestStart(t *testing.T) {
	var tests = []struct {
		params struct {
			command Command
			ID      string
			org     string
			repo    string
		}
		expected string
	}{
		{
			params: struct {
				command Command
				ID      string
				org     string
				repo    string
			}{
				command: Command{
					IssueTracker:   &IssueTrackerMock{},
					GitRepo:        &GitRepoMock{},
					VersionControl: &VersionControlMock{},
				},
				ID:   "test",
				org:  "org-test",
				repo: "repo-test",
			},
			expected: "",
		},
		{
			params: struct {
				command Command
				ID      string
				org     string
				repo    string
			}{
				command: Command{
					IssueTracker:   &IssueTrackerFailedMock{},
					GitRepo:        &GitRepoMock{},
					VersionControl: &VersionControlMock{},
				},
				ID:   "test",
				org:  "org-test",
				repo: "repo-test",
			},
			expected: "failed to fetch issue: error fetching issue",
		},
		{
			params: struct {
				command Command
				ID      string
				org     string
				repo    string
			}{
				command: Command{
					IssueTracker:   &IssueTrackerMock{},
					GitRepo:        &GitRepoFailedMock{},
					VersionControl: &VersionControlMock{},
				},
				ID:   "test",
				org:  "org-test",
				repo: "repo-test",
			},
			expected: "failed to create pull request: error creating pr",
		},
		{
			params: struct {
				command Command
				ID      string
				org     string
				repo    string
			}{
				command: Command{
					IssueTracker:   &IssueTrackerMock{},
					GitRepo:        &GitRepoMock{},
					VersionControl: &VersionControlFailedMock{},
				},
				ID:   "test",
				org:  "org-test",
				repo: "repo-test",
			},
			expected: "failed to create branch, branch may exist already: error creating or pushing branch",
		},
	}

	for _, tt := range tests {
		result := tt.params.command.Start(tt.params.ID, tt.params.org, tt.params.repo)
		err := ""
		if result != nil {
			err = result.Error()
		}

		diff := deep.Equal(err, tt.expected)
		if diff != nil {
			t.Error(diff)
		}
	}
}
