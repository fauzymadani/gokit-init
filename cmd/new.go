package cmd

import (
	"fmt"

	"gokit-init/internal/banner"
	"gokit-init/internal/config"
	"gokit-init/internal/generator"

	"github.com/spf13/cobra"
)

var (
	dbType     string
	modulePath string
	withDocker bool
	cleanArch  bool
)

var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Generate a new Go project",
	Long: `Generate a new Go web application project with the specified configuration.

Examples:
  gokit-init new myapp
  gokit-init new myapp --db mysql --docker
  gokit-init new myapp --clean-arch --db postgres --module github.com/user/myapp`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		banner.Print()

		// Create configuration
		cfg := &config.ProjectConfig{
			ProjectName: args[0],
			ModulePath:  modulePath,
			Database:    dbType,
			Docker:      withDocker,
			CleanArch:   cleanArch,
		}

		// Validate configuration
		if err := cfg.Validate(); err != nil {
			fmt.Printf("Error: Configuration error: %v\n", err)
			return
		}

		// Generate project
		if err := generator.Generate(cfg); err != nil {
			fmt.Printf("Error: Generation failed: %v\n", err)
			return
		}
	},
}

func init() {
	newCmd.Flags().StringVar(&dbType, "db", "", "Database type (mysql, postgres, sqlite)")
	newCmd.Flags().StringVar(&modulePath, "module", "", "Go module path (default: github.com/user/<project-name>)")
	newCmd.Flags().BoolVar(&withDocker, "docker", false, "Include Dockerfile and docker-compose")
	newCmd.Flags().BoolVar(&cleanArch, "clean-arch", false, "Generate Clean Architecture structure")
}
