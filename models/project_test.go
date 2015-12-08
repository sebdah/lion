package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	project := NewProject()
	assert.Empty(t, project.Name)
}

func TestPopulateAllProjectsFromConfig(t *testing.T) {
	projects := PopulateAllProjectsFromConfig()
	assert.Equal(t, 3, len(projects))
}

func TestProjectPopulateFromConfig(t *testing.T) {
	project := NewProject()
	project.PopulateFromConfig("ios")
	assert.Equal(t, "ios", project.Name)
}
