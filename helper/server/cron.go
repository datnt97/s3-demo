package server

import (
	"context"
	"os"

	"github.com/spf13/cobra"
)

type CronHandler interface {
	GetRoot() *cobra.Command
	Command() *cobra.Command
}

const codeFailure = 1

func Cron(ctx context.Context, register func(cmd *cobra.Command)) {
	cmd := &cobra.Command{
		Use:   "cron",
		Args:  cobra.NoArgs,
		Short: "exec cron job",
	}

	Initialize()
	register(cmd)

	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(codeFailure)
	}
}

func AddCommand(cli CronHandler) {
	cli.GetRoot().AddCommand(cli.Command())
}

func AddCommands(clis ...CronHandler) {
	for _, i := range clis {
		AddCommand(i)
	}
}

// func CronContext[T any](cmd *cobra.Command) (*T, error) {
// 	reg, ok := cmd.Context().Value(define.CronContextKey).(*T)
// 	if !ok {
// 		return nil, fmt.Errorf("%s: missing cron context", cmd.Name())
// 	}
// 	return reg, nil
// }
