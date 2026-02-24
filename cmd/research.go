package cmd

import (
	"encoding/json"
	"fmt"

	"exa-cli/client"
	"github.com/spf13/cobra"
)

func newResearchCmd() *cobra.Command {
	var jsonOut bool

	cmd := &cobra.Command{
		Use:   "research <instructions>",
		Short: "Run a deep research task",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c := clientFromContext(cmd)
			instructions := joinArgs(args)

			req := client.ResearchRequest{
				Instructions: instructions,
			}

			resp, err := c.Research(req)
			if err != nil {
				return fmt.Errorf("research failed: %w", err)
			}

			if jsonOut {
				b, _ := json.MarshalIndent(resp, "", "  ")
				fmt.Println(string(b))
				return nil
			}

			fmt.Println("\n=== Research Result ===")
			switch v := resp.Data.(type) {
			case string:
				fmt.Println(wrap(v, 80))
			default:
				b, _ := json.MarshalIndent(v, "", "  ")
				fmt.Println(string(b))
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&jsonOut, "json", false, "Output as JSON")
	return cmd
}
