package config

import (
	"fmt"
	"regexp"
	"strings"
)

// ProjectConfig holds all configuration for project generation
type ProjectConfig struct {
	ProjectName string
	ModulePath  string
	Database    string
	Docker      bool
	CleanArch   bool
}

// Validate checks if the configuration is valid
func (c *ProjectConfig) Validate() error {
	// Validate project name
	if c.ProjectName == "" {
		return fmt.Errorf("project name cannot be empty")
	}

	// Check if project name is valid for directories
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, c.ProjectName)
	if !matched {
		return fmt.Errorf("project name can only contain letters, numbers, hyphens, and underscores")
	}

	// Validate database type if specified
	if c.Database != "" {
		validDBs := []string{"mysql", "postgres", "sqlite"}
		valid := false
		for _, db := range validDBs {
			if strings.ToLower(c.Database) == db {
				valid = true
				c.Database = db // normalize to lowercase
				break
			}
		}
		if !valid {
			return fmt.Errorf("invalid database type: %s (valid options: mysql, postgres, sqlite)", c.Database)
		}
	}

	// Set default module path if not provided
	if c.ModulePath == "" {
		c.ModulePath = "github.com/user/" + c.ProjectName
	}

	return nil
}
