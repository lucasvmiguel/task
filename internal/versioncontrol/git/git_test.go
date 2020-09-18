package git

import (
	"testing"

	"github.com/lucasvmiguel/task/internal/versioncontrol"
	"github.com/pkg/errors"

	"github.com/go-test/deep"
)

func TestSplitOriginURL(t *testing.T) {
	var tests = []struct {
		params        string
		expected      *versioncontrol.Origin
		expectedError error
	}{
		{
			"git@github.com:lucasvmiguel/task.git\n",
			&versioncontrol.Origin{Org: "lucasvmiguel", Repository: "task"},
			nil,
		},
		{
			"lucasvmiguel/task.git",
			&versioncontrol.Origin{},
			errors.New("invalid remote url"),
		},
		{
			"git@github.com:bla",
			&versioncontrol.Origin{},
			errors.New("invalid remote org and repository"),
		},
		{
			"",
			&versioncontrol.Origin{},
			errors.New("invalid remote url"),
		},
	}

	for _, tt := range tests {
		result, err := splitOriginURL(tt.params)

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
