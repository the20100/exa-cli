package cmd

import (
	"context"

	"exa-cli/client"
	"github.com/spf13/cobra"
)

type contextKeyType string

const contextClientKey contextKeyType = clientKey

func contextWithClient(ctx context.Context, c *client.Client) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, contextClientKey, c)
}

func clientFromContext(cmd *cobra.Command) *client.Client {
	return cmd.Context().Value(contextClientKey).(*client.Client)
}
