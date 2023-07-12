package cmd

import (
	"context"
	"log"

	"github.com/katallaxie/csync/internal/checker"
	"github.com/katallaxie/csync/internal/config"
	"github.com/katallaxie/csync/internal/linker"
	"github.com/spf13/cobra"
)

var cfg = config.New()

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(RestoreCmd)
	RootCmd.AddCommand(UnlinkCmd)
	RootCmd.AddCommand(ValidateCmd)
	RootCmd.AddCommand(BackupCmd)

	RootCmd.PersistentFlags().StringVarP(&cfg.File, "config", "c", cfg.File, "config file")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Verbose, "verbose", "v", cfg.Flags.Verbose, "verbose output")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Dry, "dry", "d", cfg.Flags.Dry, "dry run")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Root, "root", "r", cfg.Flags.Root, "run as root")
	RootCmd.PersistentFlags().BoolVarP(&cfg.Flags.Force, "force", "f", cfg.Flags.Force, "force init")

	RootCmd.SilenceErrors = true
}

func initConfig() {
	err := cfg.InitDefaultConfig()
	if err != nil {
		log.Fatal(err)
	}
}

var RootCmd = &cobra.Command{
	Use:   "csync",
	Short: "csync",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return checkEnv(cmd.Context())
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRoot(cmd.Context())
	},
}

func checkEnv(ctx context.Context) error {
	c := checker.New(
		checker.WithChecks(checker.UseableEnv),
		checker.WithChecks(checker.UseSetup),
	)

	if err := c.Check(ctx, cfg); err != nil {
		return err
	}

	return nil
}

func runRoot(ctx context.Context) error {
	s, err := cfg.LoadSpec()
	if err != nil {
		return err
	}

	s.Lock()
	defer s.Unlock()

	if cfg.Flags.Verbose {
		log.Printf("Backup apps ...")
	}

	opts := []linker.Opt{linker.WithProvider(s.Provider)}
	if cfg.Flags.Verbose {
		opts = append(opts, linker.WithVerbose())
	}

	l := linker.New(opts...)

	for _, a := range s.Apps {
		a := a

		if cfg.Flags.Verbose {
			log.Printf("Backup '%s'", a.Name)
		}

		if err := l.Backup(ctx, &a, cfg.Flags.Force, cfg.Flags.Dry); err != nil {
			return err
		}
	}

	return nil
}
