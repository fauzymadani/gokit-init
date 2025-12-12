package generator

import (
	"bytes"
	_ "embed"
	"fmt"

	"gokit-init/internal/config"
)

//go:embed templates/base/main.go.tmpl
var mainTemplate string

//go:embed templates/base/main-cleanarch.go.tmpl
var mainCleanArchTemplate string

// GenerateMainFile creates the main.go file based on architecture choice
func GenerateMainFile(cfg *config.ProjectConfig) error {
	var filePath string
	var tmplContent string

	if cfg.CleanArch {
		filePath = "cmd/app/main.go"
		tmplContent = mainCleanArchTemplate
	} else {
		filePath = "main.go"
		tmplContent = mainTemplate
	}

	// Render template with project name and module path
	var buf bytes.Buffer
	if err := renderTemplate(&buf, tmplContent, map[string]string{
		"ProjectName": cfg.ProjectName,
		"ModulePath":  cfg.ModulePath,
	}); err != nil {
		return fmt.Errorf("failed to render main template: %w", err)
	}

	if err := WriteFile(cfg.ProjectName, filePath, buf.String()); err != nil {
		return err
	}

	fmt.Printf("Created %s\n", filePath)
	return nil
}
