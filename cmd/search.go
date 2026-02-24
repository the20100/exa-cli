package cmd

import (
	"fmt"

	"exa-cli/client"
	"github.com/spf13/cobra"
)

func newSearchCmd() *cobra.Command {
	var (
		numResults     int
		startDate      string
		endDate        string
		includeDomains []string
		excludeDomains []string
		includeText    []string
		excludeText    []string
		category       string
		searchType     string
		withText       bool
		withSummary    bool
		withHighlights bool
		livecrawl      string
		maxAge         int
		jsonOut        bool
	)

	cmd := &cobra.Command{
		Use:   "search <query>",
		Short: "Search the web with Exa",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c := clientFromContext(cmd)
			query := joinArgs(args)

			req := client.SearchRequest{
				Query:          query,
				IncludeDomains: includeDomains,
				ExcludeDomains: excludeDomains,
				IncludeText:    includeText,
				ExcludeText:    excludeText,
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
			if category != "" {
				req.Category = &category
			}
			if searchType != "" {
				req.Type = &searchType
			}

			contents := buildContents(withText, withSummary, withHighlights, livecrawl, maxAge)
			if contents != nil {
				req.Contents = contents
			}

			resp, err := c.Search(req)
			if err != nil {
				return fmt.Errorf("search failed: %w", err)
			}

			if resp.ResolvedSearchType != nil {
				fmt.Printf("Search type: %s\n", *resp.ResolvedSearchType)
			}
			fmt.Printf("Found %d results\n", len(resp.Results))
			printResults(resp.Results, jsonOut)
			return nil
		},
	}

	cmd.Flags().IntVarP(&numResults, "num-results", "n", 10, "Number of results")
	cmd.Flags().StringVar(&startDate, "start-date", "", "Filter results after date (YYYY-MM-DD)")
	cmd.Flags().StringVar(&endDate, "end-date", "", "Filter results before date (YYYY-MM-DD)")
	cmd.Flags().StringSliceVar(&includeDomains, "include-domains", nil, "Only include these domains")
	cmd.Flags().StringSliceVar(&excludeDomains, "exclude-domains", nil, "Exclude these domains")
	cmd.Flags().StringSliceVar(&includeText, "include-text", nil, "Results must include this text")
	cmd.Flags().StringSliceVar(&excludeText, "exclude-text", nil, "Results must not include this text")
	cmd.Flags().StringVar(&category, "category", "", "Filter by category (news, research paper, company, pdf, tweet, etc.)")
	cmd.Flags().StringVar(&searchType, "type", "", "Search type (auto, neural, fast, deep)")
	cmd.Flags().BoolVar(&withText, "text", false, "Include page text in results")
	cmd.Flags().BoolVar(&withSummary, "summary", false, "Include AI summary in results")
	cmd.Flags().BoolVar(&withHighlights, "highlights", false, "Include highlights in results")
	cmd.Flags().StringVar(&livecrawl, "livecrawl", "", "Live crawl mode (always, fallback, never, auto)")
	cmd.Flags().IntVar(&maxAge, "max-age", 0, "Maximum content age in hours")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "Output as JSON")

	return cmd
}
