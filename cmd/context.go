package cmd

import (
	"context"

	"exa-cli/exa"
	"github.com/spf13/cobra"
)

type contextKeyType string

const contextClientKey contextKeyType = clientKey

func contextWithClient(ctx context.Context, client *exa.Client) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, contextClientKey, client)
}

func clientFromContext(cmd *cobra.Command) *exa.Client {
	return cmd.Context().Value(contextClientKey).(*exa.Client)
}
