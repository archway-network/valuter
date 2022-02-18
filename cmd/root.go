package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "valuter"}

func init() {}

func Execute() {
	initClientCtx := client.Context{}.WithInput(os.Stdin)
	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &initClientCtx)

	if err := rootCmd.ExecuteContext(ctx); err != nil {

		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
