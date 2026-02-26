package cmd

import (
	"fmt"
	"os"
	"strings"

	"exa-cli/client"
	"github.com/spf13/cobra"
)

const clientKey = "exa_client"

func NewRootCmd() *cobra.Command {
	var (
		apiKey  string
		baseURL string
	)

	root := &cobra.Command{
		Use:   "exa",
		Short: "Exa AI search CLI",
		Long: `exa is a command-line interface for the Exa AI search API.

Set your API key via the EXA_API_KEY environment variable or --api-key flag.

Available commands:
  search        Search the web
  find-similar  Find pages similar to a URL
  get-contents  Retrieve full content from URLs
  answer        Get a direct AI answer with citations
  research      Run a deep research task`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			key := apiKey
			if key == "" {
				key = os.Getenv("EXA_API_KEY")
			}
			if key == "" {
				return fmt.Errorf("API key required: set EXA_API_KEY or use --api-key")
			}
			c := client.NewClient(key, baseURL)
			cmd.SetContext(contextWithClient(cmd.Context(), c))
			return nil
		},
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
