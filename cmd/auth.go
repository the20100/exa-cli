package cmd

import (
	"fmt"
	"os"

	"exa-cli/internal/config"

	"github.com/spf13/cobra"
)

func newAuthCmd() *cobra.Command {
	authCmd := &cobra.Command{
		Use:   "auth",
		Short: "Manage Exa authentication",
	}

	authCmd.AddCommand(newAuthSetKeyCmd(), newAuthStatusCmd(), newAuthLogoutCmd())
	return authCmd
}

func newAuthSetKeyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "set-key <api-key>",
		Short: "Save an Exa API key to the config file",
		Long: `Save an Exa API key to the local config file.

Get your API key from: https://dashboard.exa.ai/api-keys

The key is stored at:
  macOS:   ~/Library/Application Support/exa/config.json
  Linux:   ~/.config/exa/config.json
  Windows: %AppData%\exa\config.json

You can also set the EXA_API_KEY env var instead of using this command.`,
		Args:    cobra.ExactArgs(1),
		Example: "  exa auth set-key your_api_key_here",
		RunE: func(cmd *cobra.Command, args []string) error {
			key := args[0]
			if len(key) < 8 {
				return fmt.Errorf("API key looks too short — check your key at https://dashboard.exa.ai/api-keys")
			}

			c := &config.Config{APIKey: key}
			if err := config.Save(c); err != nil {
				return fmt.Errorf("saving config: %w", err)
			}

			fmt.Printf("API key saved to %s\n", config.Path())
			fmt.Printf("Key: %s\n", maskOrEmpty(key))
			return nil
		},
	}
}

func newAuthStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show current authentication status",
		RunE: func(cmd *cobra.Command, args []string) error {
			c, err := config.Load()
			if err != nil {
				return fmt.Errorf("loading config: %w", err)
			}

			fmt.Printf("Config: %s\n", config.Path())
			fmt.Println()

			if envKey := os.Getenv("EXA_API_KEY"); envKey != "" {
				fmt.Println("Key source: EXA_API_KEY env var (takes priority over config)")
				fmt.Printf("Key:        %s\n", maskOrEmpty(envKey))
			} else if c.APIKey != "" {
				fmt.Println("Key source: config file")
				fmt.Printf("Key:        %s\n", maskOrEmpty(c.APIKey))
			} else {
				fmt.Println("Status: not authenticated")
				fmt.Println()
				fmt.Println("Run: exa auth set-key <your-api-key>")
				fmt.Println("Or:  export EXA_API_KEY=<your-api-key>")
			}
			return nil
		},
	}
}

func newAuthLogoutCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Remove the saved API key from the config file",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := config.Clear(); err != nil {
				return fmt.Errorf("removing config: %w", err)
			}
			fmt.Println("API key removed from config.")
			fmt.Println("Set EXA_API_KEY env var if you still need access.")
			return nil
		},
	}
}

func maskOrEmpty(v string) string {
	if v == "" {
		return "(not set)"
	}
	if len(v) <= 8 {
		return "***"
	}
	return v[:4] + "..." + v[len(v)-4:]
}
