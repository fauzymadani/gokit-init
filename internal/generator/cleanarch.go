package generator

import (
	"bytes"
	_ "embed"
	"fmt"

	"gokit-init/internal/config"
)

//go:embed templates/cleanarch/domain.go.tmpl
var domainTemplate string

//go:embed templates/cleanarch/repository.go.tmpl
var repositoryTemplate string

//go:embed templates/cleanarch/service.go.tmpl
var serviceTemplate string

//go:embed templates/cleanarch/handler.go.tmpl
var handlerTemplate string

// GenerateCleanArchFiles creates additional files for Clean Architecture pattern
func GenerateCleanArchFiles(cfg *config.ProjectConfig) error {
	if !cfg.CleanArch {
		return nil
	}

	templateData := map[string]string{
		"ModulePath": cfg.ModulePath,
	}

	// Generate sample domain entity
	if err := WriteFile(cfg.ProjectName, "internal/domain/user.go", domainTemplate); err != nil {
		return err
	}

	// Generate sample repository
	var buf bytes.Buffer
	if err := renderTemplate(&buf, repositoryTemplate, templateData); err != nil {
		return fmt.Errorf("failed to render repository template: %w", err)
	}
	if err := WriteFile(cfg.ProjectName, "internal/repository/user_repository.go", buf.String()); err != nil {
		return err
	}

	// Generate sample service
	buf.Reset()
	if err := renderTemplate(&buf, serviceTemplate, templateData); err != nil {
		return fmt.Errorf("failed to render service template: %w", err)
	}
	if err := WriteFile(cfg.ProjectName, "internal/service/user_service.go", buf.String()); err != nil {
		return err
	}

	// Generate sample handler
	buf.Reset()
	if err := renderTemplate(&buf, handlerTemplate, templateData); err != nil {
		return fmt.Errorf("failed to render handler template: %w", err)
	}
	if err := WriteFile(cfg.ProjectName, "internal/handler/user_handler.go", buf.String()); err != nil {
		return err
	}

	fmt.Println("Created Clean Architecture files")
	return nil
}
