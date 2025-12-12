package generator

import (
	"bytes"
	_ "embed"
	"fmt"

	"gokit-init/internal/config"
)

//go:embed templates/database/mysql.go.tmpl
var mysqlTemplate string

//go:embed templates/database/postgres.go.tmpl
var postgresTemplate string

//go:embed templates/database/sqlite.go.tmpl
var sqliteTemplate string

// GenerateDatabaseConfig creates database configuration file if database is specified
func GenerateDatabaseConfig(cfg *config.ProjectConfig) error {
	if cfg.Database == "" {
		return nil
	}

	var tmplContent string
	var filePath string

	if cfg.CleanArch {
		filePath = "internal/config/database.go"
	} else {
		filePath = "handler/db.go"
	}

	pkgName := "handler"
	if cfg.CleanArch {
		pkgName = "config"
	}

	switch cfg.Database {
	case "mysql":
		tmplContent = mysqlTemplate
	case "postgres":
		tmplContent = postgresTemplate
	case "sqlite":
		tmplContent = sqliteTemplate
	default:
		return fmt.Errorf("unsupported database: %s", cfg.Database)
	}

	// Render template
	var buf bytes.Buffer
	if err := renderTemplate(&buf, tmplContent, map[string]string{
		"PackageName": pkgName,
	}); err != nil {
		return fmt.Errorf("failed to render database template: %w", err)
	}

	if err := WriteFile(cfg.ProjectName, filePath, buf.String()); err != nil {
		return err
	}

	fmt.Printf("Created database config (%s)\n", cfg.Database)
	return nil
}
