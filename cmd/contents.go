package cmd

import (
	"fmt"

	"exa-cli/exa"
	"github.com/spf13/cobra"
)

func newGetContentsCmd() *cobra.Command {
	var (
		withText       bool
		withSummary    bool
		withHighlights bool
		livecrawl      string
		maxAge         int
		jsonOut        bool
	)

	cmd := &cobra.Command{
		Use:   "get-contents <url> [url...]",
		Short: "Retrieve content from URLs",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := clientFromContext(cmd)

			req := exa.GetContentsRequest{
				IDs: args,
			}

			contents := buildContents(withText, withSummary, withHighlights, livecrawl, maxAge)
			if contents == nil {
				// Default: include text
				t := true
				_ = t
				contents = &exa.ContentsOptions{
					Text: &exa.TextOptions{},
				}
			}
			req.Contents = contents

			resp, err := client.GetContents(req)
			if err != nil {
				return fmt.Errorf("get-contents failed: %w", err)
			}

			fmt.Printf("Retrieved %d pages\n", len(resp.Results))
			printResults(resp.Results, jsonOut)
			return nil
		},
	}

	cmd.Flags().BoolVar(&withText, "text", true, "Include page text")
	cmd.Flags().BoolVar(&withSummary, "summary", false, "Include AI summary")
	cmd.Flags().BoolVar(&withHighlights, "highlights", false, "Include highlights")
	cmd.Flags().StringVar(&livecrawl, "livecrawl", "", "Live crawl mode (always, fallback, never, auto)")
	cmd.Flags().IntVar(&maxAge, "max-age", 0, "Maximum content age in hours")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "Output as JSON")

	return cmd
}
