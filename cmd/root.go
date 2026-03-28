package cmd

import (
	"fmt"
	"os"
	"strings"

	"exa-cli/client"
	"exa-cli/internal/config"

	"github.com/spf13/cobra"
)

const clientKey = "exa_client"

// resolveEnv returns the value of the first non-empty environment variable from the given names.
func resolveEnv(names ...string) string {
	for _, name := range names {
		if v := os.Getenv(name); v != "" {
			return v
		}
	}
	return ""
}

// isAuthCommand returns true if cmd is a child of the "auth" command.
func isAuthCommand(cmd *cobra.Command) bool {
	if cmd.Name() == "auth" {
		return true
	}
	p := cmd.Parent()
	for p != nil {
		if p.Name() == "auth" {
			return true
		}
		p = p.Parent()
	}
	return false
}

func NewRootCmd() *cobra.Command {
	var (
		apiKey  string
		baseURL string
	)

	root := &cobra.Command{
		Use:   "exa",
		Short: "Exa AI search CLI",
		Long: `exa is a command-line interface for the Exa AI search API.

Token resolution order:
  1. EXA_API_KEY env var (or aliases: EXA_KEY, EXA_API, ...)
  2. Config file (~/.config/exa/config.json via: exa auth set-key)
  3. --api-key flag

Available commands:
  search        Search the web
  find-similar  Find pages similar to a URL
  get-contents  Retrieve full content from URLs
  answer        Get a direct AI answer with citations
  research      Run a deep research task`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if isAuthCommand(cmd) {
				return nil
			}

			key := apiKey
			if key == "" {
				key = resolveEnv(
					"EXA_API_KEY", "EXA_KEY", "EXA_API", "API_KEY_EXA", "API_EXA", "EXA_PK", "EXA_PUBLIC",
					"EXA_API_SECRET", "EXA_SECRET_KEY", "EXA_API_SECRET_KEY", "EXA_SECRET", "SECRET_EXA", "API_SECRET_EXA", "SK_EXA", "EXA_SK",
				)
			}
			if key == "" {
				cfg, err := config.Load()
				if err != nil {
					return fmt.Errorf("failed to load config: %w", err)
				}
				if cfg.APIKey != "" {
					key = cfg.APIKey
				}
			}
			if key == "" {
				return fmt.Errorf("not authenticated — run: exa auth set-key\nor set EXA_API_KEY env var")
			}
			c := client.NewClient(key, baseURL)
			cmd.SetContext(contextWithClient(cmd.Context(), c))
			return nil
		},
		SilenceUsage: true,
	}

	root.PersistentFlags().StringVar(&apiKey, "api-key", "", "Exa API key (or set EXA_API_KEY)")
	root.PersistentFlags().StringVar(&baseURL, "base-url", "", "Exa API base URL")

	root.AddCommand(
		newSearchCmd(),
		newFindSimilarCmd(),
		newGetContentsCmd(),
		newAnswerCmd(),
		newResearchCmd(),
		newUpdateCmd(),
		newAuthCmd(),
	)

	return root
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func joinArgs(args []string) string {
	return strings.Join(args, " ")
}

func buildContents(withText, withSummary, withHighlights bool, livecrawl string, maxAge int) *client.ContentsOptions {
	if !withText && !withSummary && !withHighlights && livecrawl == "" && maxAge == 0 {
		return nil
	}
	c := &client.ContentsOptions{}
	if withText {
		c.Text = &client.TextOptions{}
	}
	if withSummary {
		c.Summary = &client.SummaryOptions{}
	}
	if withHighlights {
		c.Highlights = &client.HighlightOptions{}
	}
	if livecrawl != "" {
		c.Livecrawl = &livecrawl
	}
	if maxAge > 0 {
		c.MaxAgeHours = &maxAge
	}
	return c
}
