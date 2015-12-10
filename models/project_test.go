package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProject(t *testing.T) {
	project := NewProject()
	assert.Empty(t, project.Name)
}
