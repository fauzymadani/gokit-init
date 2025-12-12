package generator

import (
	"fmt"
	"os"
	"path/filepath"

	"gokit-init/internal/config"
)

// CreateDirectoryStructure creates all necessary directories for the project
func CreateDirectoryStructure(cfg *config.ProjectConfig) error {
	// Check if project directory already exists
	if _, err := os.Stat(cfg.ProjectName); !os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' already exists", cfg.ProjectName)
	}

	// Create base project directory
	if err := os.MkdirAll(cfg.ProjectName, 0755); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	var dirs []string

	if cfg.CleanArch {
		// Clean Architecture structure
		dirs = []string{
			"cmd/app",
			"internal/config",
			"internal/handler",
			"internal/service",
			"internal/repository",
			"internal/domain",
			"pkg",
		}
	} else {
		// Simple structure
		dirs = []string{
			"handler",
			"domain",
		}
	}

	// Create all directories
	for _, dir := range dirs {
		path := filepath.Join(cfg.ProjectName, dir)
		if err := os.MkdirAll(path, 0755); err != nil {
			return fmt.Errorf("failed to create directory '%s': %w", path, err)
		}
	}

	fmt.Printf("Created directory structure for '%s'\n", cfg.ProjectName)
	return nil
}
