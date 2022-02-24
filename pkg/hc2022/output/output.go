package output

import (
	"fmt"
	"strings"

	"github.com/ismtabo/hashcode-2022-qr/pkg/hc2022/input"
)

type Solution struct {
	Projects []*Project
}

func (s Solution) String() string {
	projects := []string{}
	for _, project := range s.Projects {
		projects = append(projects, project.String())
	}
	return strings.Join(
		[]string{
			fmt.Sprint(len(s.Projects)),
			strings.Join(projects, "\n"),
		},
		"\n",
	)
}

type Project struct {
	Name      string
	DaysSpent int
	Assignees []*input.Contributor
}

func (s Project) String() string {
	assignees := []string{}
	for _, assignee := range s.Assignees {
		assignees = append(assignees, assignee.Name)
	}
	return strings.Join(
		[]string{
			s.Name,
			strings.Join(assignees, " "),
		},
		"\n",
	)
}
