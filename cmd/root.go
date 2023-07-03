package cmd

import (
	"context"

	"github.com/katallaxie/csync/pkg/config"
	"github.com/spf13/cobra"
)

var cfg = config.New()

func init() {
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(RestoreCmd)
	RootCmd.AddCommand(BackupCmd)
	RootCmd.AddCommand(UnlinkCmd)
	RootCmd.AddCommand(ValidateCmd)

	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Verbose, "verbose", "v", cfg.Flags.Verbose, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Dry, "dry", "d", cfg.Flags.Dry, "dry run")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Root, "root", "r", cfg.Flags.Root, "run as root")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Force, "force", "f", cfg.Flags.Force, "force init")
}

var RootCmd = &cobra.Command{
	Use:   "csync",
	Short: "csync",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context())
	},
}

func runRoot(ctx context.Context) error {
	return nil
}
