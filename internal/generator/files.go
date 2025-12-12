package generator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	"gokit-init/internal/config"
)

// renderTemplate is a helper to render template strings
func renderTemplate(w io.Writer, tmplContent string, data interface{}) error {
	tmpl, err := template.New("tmpl").Parse(tmplContent)
	if err != nil {
		return err
	}
	return tmpl.Execute(w, data)
}

// WriteFile writes content to a file in the project directory
func WriteFile(projectName, filePath, content string) error {
	fullPath := filepath.Join(projectName, filePath)

	// Ensure parent directory exists
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory for file '%s': %w", filePath, err)
	}

	// Write file
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file '%s': %w", filePath, err)
	}

	return nil
}

// WriteFileFromTemplate renders a template and writes it to a file
func WriteFileFromTemplate(projectName, filePath, tmplContent string, data interface{}) error {
	tmpl, err := template.New(filePath).Parse(tmplContent)
	if err != nil {
		return fmt.Errorf("failed to parse template for '%s': %w", filePath, err)
	}

	fullPath := filepath.Join(projectName, filePath)

	// Ensure parent directory exists
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory for file '%s': %w", filePath, err)
	}

	// Create file
	f, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("failed to create file '%s': %w", filePath, err)
	}
	defer f.Close()

	// Execute template
	if err := tmpl.Execute(f, data); err != nil {
		return fmt.Errorf("failed to execute template for '%s': %w", filePath, err)
	}

	return nil
}

// GenerateGoMod creates the go.mod file
func GenerateGoMod(cfg *config.ProjectConfig) error {
	content := fmt.Sprintf(`module %s

go 1.21

`, cfg.ModulePath)

	if err := WriteFile(cfg.ProjectName, "go.mod", content); err != nil {
		return err
	}

	fmt.Println("Created go.mod")
	return nil
}

// GenerateEnvExample creates the .env.example file
func GenerateEnvExample(cfg *config.ProjectConfig) error {
	content := `APP_PORT=8080
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=
DB_NAME=app
`

	if err := WriteFile(cfg.ProjectName, ".env.example", content); err != nil {
		return err
	}

	fmt.Println("Created .env.example")
	return nil
}
