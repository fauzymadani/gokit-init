package generator

import (
	"bytes"
	_ "embed"
	"fmt"

	"gokit-init/internal/config"
)

//go:embed templates/docker/Dockerfile.tmpl
var dockerfileTemplate string

//go:embed templates/docker/docker-compose.yml.tmpl
var dockerComposeTemplate string

// GenerateDockerFiles creates Dockerfile and docker-compose.yml if Docker flag is set
func GenerateDockerFiles(cfg *config.ProjectConfig) error {
	if !cfg.Docker {
		return nil
	}

	// Generate Dockerfile
	if err := generateDockerfile(cfg); err != nil {
		return err
	}

	// Generate docker-compose.yml
	if err := generateDockerCompose(cfg); err != nil {
		return err
	}

	fmt.Println("Created Docker files")
	return nil
}

func generateDockerfile(cfg *config.ProjectConfig) error {
	templateData := map[string]interface{}{
		"CleanArch": cfg.CleanArch,
	}

	// Render template
	var buf bytes.Buffer
	if err := renderTemplate(&buf, dockerfileTemplate, templateData); err != nil {
		return fmt.Errorf("failed to render Dockerfile template: %w", err)
	}

	return WriteFile(cfg.ProjectName, "Dockerfile", buf.String())
}

func generateDockerCompose(cfg *config.ProjectConfig) error {
	templateData := map[string]interface{}{
		"ProjectName": cfg.ProjectName,
		"Database":    cfg.Database,
		"CleanArch":   cfg.CleanArch,
	}

	var buf bytes.Buffer
	if err := renderTemplate(&buf, dockerComposeTemplate, templateData); err != nil {
		return fmt.Errorf("failed to render docker-compose template: %w", err)
	}

	return WriteFile(cfg.ProjectName, "docker-compose.yml", buf.String())
}
