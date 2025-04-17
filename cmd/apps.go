package cmd

import (
	"context"
	"log"

	"github.com/katallaxie/csync/pkg/spec"
	"github.com/spf13/cobra"
)

func init() {
	AppsCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "List apps",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runAppsList(cmd.Context())
		},
	})
}

var AppsCmd = &cobra.Command{
	Use:   "apps",
	Short: "Manage apps",
}

func runAppsList(_ context.Context) error {
	apps := cfg.Spec.GetApps(spec.List()...)

	for _, app := range apps {
		log.Printf("%s", app.Name)
	}

	return nil
}
