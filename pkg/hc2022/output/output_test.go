package output

import (
	"testing"

	"github.com/ismtabo/hashcode-2022-qr/pkg/hc2022/input"
	"github.com/magiconair/properties/assert"
)

func TestSolutionOutput(t *testing.T) {
	solution := Solution{
		Projects: []*Project{
			{
				Name:      "WebServer",
				DaysSpent: 0,
				Assignees: []*input.Contributor{
					{Name: "Bob"},
					{Name: "Anna"},
				},
			},
			{
				Name:      "Logging",
				DaysSpent: 0,
				Assignees: []*input.Contributor{
					{Name: "Anna"},
				},
			},
			{
				Name:      "WebChat",
				DaysSpent: 0,
				Assignees: []*input.Contributor{
					{Name: "Maria"},
					{Name: "Bob"},
				},
			},
		},
	}
	expected := `3
WebServer
Bob Anna
Logging
Anna
WebChat
Maria Bob`
	assert.Matches(t, solution.String(), expected)
}
