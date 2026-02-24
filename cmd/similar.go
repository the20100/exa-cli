package cmd

import (
	"fmt"

	"exa-cli/exa"
	"github.com/spf13/cobra"
)

func newFindSimilarCmd() *cobra.Command {
	var (
		numResults          int
		excludeSourceDomain bool
		includeDomains      []string
		excludeDomains      []string
		startDate           string
		endDate             string
		withText            bool
		withSummary         bool
		withHighlights      bool
		livecrawl           string
		maxAge              int
		jsonOut             bool
	)

	cmd := &cobra.Command{
		Use:   "find-similar <url>",
		Short: "Find pages similar to a URL",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := clientFromContext(cmd)

			req := exa.FindSimilarRequest{
				URL:                 args[0],
				ExcludeSourceDomain: excludeSourceDomain,
				IncludeDomains:      includeDomains,
				ExcludeDomains:      excludeDomains,
			}
			if numResults > 0 {
				req.NumResults = &numResults
			}
			if startDate != "" {
				req.StartPublishedDate = &startDate
			}
			if endDate != "" {
				req.EndPublishedDate = &endDate
			}

			contents := buildContents(withText, withSummary, withHighlights, livecrawl, maxAge)
			if contents != nil {
				req.Contents = contents
			}

			resp, err := client.FindSimilar(req)
			if err != nil {
				return fmt.Errorf("find-similar failed: %w", err)
			}

			fmt.Printf("Found %d similar pages\n", len(resp.Results))
			printResults(resp.Results, jsonOut)
			return nil
		},
	}

	cmd.Flags().IntVarP(&numResults, "num-results", "n", 10, "Number of results")
	cmd.Flags().BoolVar(&excludeSourceDomain, "exclude-source-domain", false, "Exclude results from the source domain")
	cmd.Flags().StringSliceVar(&includeDomains, "include-domains", nil, "Only include these domains")
	cmd.Flags().StringSliceVar(&excludeDomains, "exclude-domains", nil, "Exclude these domains")
	cmd.Flags().StringVar(&startDate, "start-date", "", "Filter results after date (YYYY-MM-DD)")
	cmd.Flags().StringVar(&endDate, "end-date", "", "Filter results before date (YYYY-MM-DD)")
	cmd.Flags().BoolVar(&withText, "text", false, "Include page text in results")
	cmd.Flags().BoolVar(&withSummary, "summary", false, "Include AI summary in results")
	cmd.Flags().BoolVar(&withHighlights, "highlights", false, "Include highlights in results")
	cmd.Flags().StringVar(&livecrawl, "livecrawl", "", "Live crawl mode (always, fallback, never, auto)")
	cmd.Flags().IntVar(&maxAge, "max-age", 0, "Maximum content age in hours")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "Output as JSON")

	return cmd
}
