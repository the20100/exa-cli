package cmd

import (
	"fmt"

	"exa-cli/client"
	"github.com/spf13/cobra"
)

func newAnswerCmd() *cobra.Command {
	var (
		model        string
		systemPrompt string
		userLocation string
		withText     bool
		jsonOut      bool
	)

	cmd := &cobra.Command{
		Use:   "answer <query>",
		Short: "Get a direct answer to a question with citations",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c := clientFromContext(cmd)
			query := joinArgs(args)

			req := client.AnswerRequest{
				Query: query,
			}
			if model != "" {
				req.Model = &model
			}
			if systemPrompt != "" {
				req.SystemPrompt = &systemPrompt
			}
			if userLocation != "" {
				req.UserLocation = &userLocation
			}
			if withText {
				t := true
				req.Text = &t
			}

			resp, err := c.Answer(req)
			if err != nil {
				return fmt.Errorf("answer failed: %w", err)
			}

			printAnswer(resp, jsonOut)
			return nil
		},
	}

	cmd.Flags().StringVarP(&model, "model", "m", "", "Model to use (exa, exa-pro)")
	cmd.Flags().StringVar(&systemPrompt, "system-prompt", "", "Custom system prompt")
	cmd.Flags().StringVar(&userLocation, "location", "", "User location for localized results")
	cmd.Flags().BoolVar(&withText, "text", false, "Include full text of sources")
	cmd.Flags().BoolVar(&jsonOut, "json", false, "Output as JSON")

	return cmd
}
