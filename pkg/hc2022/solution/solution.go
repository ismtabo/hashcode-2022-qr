package solution

import (
	"github.com/ismtabo/hashcode-2022-qr/pkg/hc2022/input"
	"github.com/ismtabo/hashcode-2022-qr/pkg/hc2022/output"
	"github.com/juju/collections/set"
)

type ProjectQueue []*input.Project

func (pq ProjectQueue) Len() int {
	return len(pq)
}

func (pq ProjectQueue) Less(i, j int) bool {
	a := pq[i]
	b := pq[j]
	return a.BestBfDay < b.BestBfDay ||
		a.Score > b.Score ||
		a.NRoles < b.NRoles
}

func (pq ProjectQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *ProjectQueue) Push(x interface{}) {
	item := x.(*input.Project)
	*pq = append(*pq, item)
}

func (pq *ProjectQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func Resolve(problem input.Problem) (solution output.Solution) {
	availableProjects := set.NewStrings()
	for _, project := range problem.Projects {
		availableProjects.Add(project.Name)
	}
	solution.Projects = []*output.Project{}
	for availableProjects.Size() > 0 {
		targetProject := findSuitableProject(problem.Projects, availableProjects, problem.Contributors)
		solution.Projects = append(solution.Projects, targetProject)
		availableProjects.Remove(targetProject.Name)
	}
	return solution
}

func findContributor(contributors []input.Contributor, name string) *input.Contributor {
	for _, contrib := range contributors {
		if contrib.Name == name {
			return &contrib
		}
	}
	return nil
}

func findProject(projects []input.Project, name string) *input.Project {
	for _, project := range projects {
		if project.Name == name {
			return &project
		}
	}
	return nil
}

func findSuitableProject(projects []input.Project, availableProjects set.Strings, contribs []input.Contributor) *output.Project {
	for _, projectName := range availableProjects.Values() {
		project := findProject(projects, projectName)
		selectedContribs := []*input.Contributor{}
		availableContribs := set.NewStrings()
		for _, contrib := range contribs {
			availableContribs.Add(contrib.Name)
		}
		suitable := true
		for _, role := range project.Roles {
			findContrib := false
		contribLoop:
			for _, contribName := range availableContribs.Values() {
				contrib := findContributor(contribs, contribName)
				for skillIdx, skill := range contrib.Skills {
					if role.Name == skill.Name && role.Level <= skill.Level {
						if role.Level == skill.Level {
							contrib.Skills[skillIdx].Level += 1
						}
						findContrib = true
						availableContribs.Remove(contrib.Name)
						selectedContribs = append(selectedContribs, contrib)
						break contribLoop
					}
				}
			}
			if findContrib == false {
				suitable = false
				break
			}
		}
		if suitable {
			projectCpy := output.Project{
				Name:      projectName,
				Assignees: selectedContribs,
			}
			return &projectCpy
		}
	}
	return nil
}
