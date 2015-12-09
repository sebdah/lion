package models

import (
	"errors"
	"fmt"
	"sort"

	"github.com/sebdah/lion/config"
)

type Projects struct {
	Projects []Project `json:"projects"`
}

type Project struct {
	Name string `json:"name"`
}

func NewProject() *Project {
	return &Project{}
}

func PopulateAllProjectsFromConfig() []Project {
	var projects []Project
	var projectNames []string

	// Ensure that the projects are sorted by name
	for p, _ := range config.Config.GetStringMapString("projects") {
		projectNames = append(projectNames, p)
	}
	sort.Strings(projectNames)

	for _, p := range projectNames {
		project := NewProject()
		_ = project.PopulateFromConfig(p)
		projects = append(projects, *project)
	}

	return projects
}

func (p *Project) PopulateFromConfig(name string) error {
	var found = false

	for project, _ := range config.Config.GetStringMapString("projects") {
		if project == name {
			p.Name = project
			found = true
			break
		}
	}

	if !found {
		return errors.New(fmt.Sprintf("Project %s not found", name))
	}

	return nil
}
