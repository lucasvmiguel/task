package command

import (
	"errors"
	"testing"

	"github.com/lucasvmiguel/task/internal/gitrepo"
	"github.com/lucasvmiguel/task/internal/issuetracker"
	"github.com/lucasvmiguel/task/internal/versioncontrol"

	"github.com/go-test/deep"
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

func (i *VersionControlMock) Origin() (*versioncontrol.Origin, error) {
	return &versioncontrol.Origin{Org: "test-org", Repository: "test-repo"}, nil
}

type VersionControlFailedMock struct{}

func (i *VersionControlFailedMock) CreateBranchAndPush(name string) error {
	return errors.New("error creating or pushing branch")
}

func (i *VersionControlFailedMock) Origin() (*versioncontrol.Origin, error) {
	return nil, errors.New("error getting origin")
}

func TestStart(t *testing.T) {
	var tests = []struct {
		params struct {
			command     Command
			startParams StartParams
		}
		expected string
	}{
		{
			params: struct {
				command     Command
				startParams StartParams
			}{
				command: Command{
					IssueTracker:   &IssueTrackerMock{},
					GitRepo:        &GitRepoMock{},
					VersionControl: &VersionControlMock{},
				},
				startParams: StartParams{
					ID:                  "test",
					TitleTemplate:       "testing: {{ISSUE_TRACKER.ID}}",
					DescriptionTemplate: "# Description\n\n {{ISSUE_TRACKER.TITLE}}",
				},
			},
			expected: "",
		},
		{
			params: struct {
				command     Command
				startParams StartParams
			}{
				command: Command{
					IssueTracker:   &IssueTrackerFailedMock{},
					GitRepo:        &GitRepoMock{},
					VersionControl: &VersionControlMock{},
				},
				startParams: StartParams{
					ID:                  "test",
					TitleTemplate:       "testing: {{ISSUE_TRACKER.ID}}",
					DescriptionTemplate: "# Description\n\n {{ISSUE_TRACKER.TITLE}}",
				},
			},
			expected: "failed to fetch issue: error fetching issue",
		},
		{
			params: struct {
				command     Command
				startParams StartParams
			}{
				command: Command{
					IssueTracker:   &IssueTrackerMock{},
					GitRepo:        &GitRepoFailedMock{},
					VersionControl: &VersionControlMock{},
				},
				startParams: StartParams{
					ID:                  "test",
					TitleTemplate:       "testing: {{ISSUE_TRACKER.ID}}",
					DescriptionTemplate: "# Description\n\n {{ISSUE_TRACKER.TITLE}}",
				},
			},
			expected: "failed to create pull request: error creating pr",
		},
		{
			params: struct {
				command     Command
				startParams StartParams
			}{
				command: Command{
					IssueTracker:   &IssueTrackerMock{},
					GitRepo:        &GitRepoMock{},
					VersionControl: &VersionControlFailedMock{},
				},
				startParams: StartParams{
					ID:                  "test",
					TitleTemplate:       "testing: {{ISSUE_TRACKER.ID}}",
					DescriptionTemplate: "# Description\n\n {{ISSUE_TRACKER.TITLE}}",
				},
			},
			expected: "failed to get origin from the repository: error getting origin",
		},
	}

	for _, tt := range tests {
		result := tt.params.command.Start(tt.params.startParams)
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
