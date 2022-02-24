package input

import (
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetInput() string {
	var input []byte
	if len(os.Args) < 2 {
		var err error
		input, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
	} else {
		var err error
		input, err = ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic(err)
		}
	}
	return string(input)
}

type Problem struct {
	NContributors int
	Contributors  []Contributor
	NProjects     int
	Projects      []Project
}

type Contributor struct {
	Name    string
	NSkills int
	Skills  []Skill
}

type Skill struct {
	Name  string
	Level int
}

type Project struct {
	Name      string
	NDays     int
	Score     int
	BestBfDay int
	NRoles    int
	Roles     []Role
}

type Role struct {
	Name  string
	Level int
}

var NAME_WITH_NUMBER_RE = regexp.MustCompile(`([\w+-]+) (\d+)`)
var PROJECT_RE = regexp.MustCompile(`(\w+) (\d+) (\d+) (\d+) (\d+)`)

func ParseInput(input string) (problem Problem) {
	var err error
	lines := strings.Split(input, "\n")
	info := strings.Split(lines[0], " ")
	problem.NContributors, err = strconv.Atoi(info[0])
	logPanic(err)
	problem.NProjects, err = strconv.Atoi(info[1])
	logPanic(err)
	lineIdx := 1
	// Read contributors
	problem.Contributors = []Contributor{}
	for contribIdx := 0; contribIdx < problem.NContributors; contribIdx++ {
		contributor := Contributor{}
		cSubMatchs := NAME_WITH_NUMBER_RE.FindStringSubmatch(lines[lineIdx])
		lineIdx++
		contributor.Name = cSubMatchs[1]
		contributor.NSkills, err = strconv.Atoi(cSubMatchs[2])
		logPanic(err)
		contributor.Skills = []Skill{}
		for skillIdx := 0; skillIdx < contributor.NSkills; skillIdx++ {
			skill := Skill{}
			sSubMatchs := NAME_WITH_NUMBER_RE.FindStringSubmatch(lines[lineIdx])
			lineIdx++
			skill.Name = sSubMatchs[1]
			skill.Level, err = strconv.Atoi(sSubMatchs[2])
			logPanic(err)
			contributor.Skills = append(contributor.Skills, skill)
		}
		problem.Contributors = append(problem.Contributors, contributor)
	}
	// Read projects
	problem.Projects = []Project{}
	for projIdx := 0; projIdx < problem.NProjects; projIdx++ {
		project := Project{}
		pSubMatches := PROJECT_RE.FindStringSubmatch(lines[lineIdx])
		lineIdx++
		project.Name = pSubMatches[1]
		project.NDays, err = strconv.Atoi(pSubMatches[2])
		logPanic(err)
		project.Score, err = strconv.Atoi(pSubMatches[3])
		logPanic(err)
		project.BestBfDay, err = strconv.Atoi(pSubMatches[4])
		logPanic(err)
		project.NRoles, err = strconv.Atoi(pSubMatches[5])
		logPanic(err)
		for roleIdx := 0; roleIdx < project.NRoles; roleIdx++ {
			role := Role{}
			roleSubMatches := NAME_WITH_NUMBER_RE.FindStringSubmatch(lines[lineIdx])
			lineIdx++
			role.Name = roleSubMatches[1]
			role.Level, err = strconv.Atoi(roleSubMatches[2])
			logPanic(err)
			project.Roles = append(project.Roles, role)
		}
		problem.Projects = append(problem.Projects, project)
	}
	return
}

func logPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
