package checker

import (
	"context"
	"fmt"
	"os/user"

	"github.com/katallaxie/csync/pkg/config"
)

// UsuableEnv ...
func UsuableEnv(ctx context.Context, cfg *config.Config) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	if user.Uid == "0" && !cfg.Flags.Root {
		return fmt.Errorf("running as 'root' is dangerous. You can overwrite this behavior with '--root'")
	}

	return nil
}

// HomeFolder ...
func HomeFolder(ctx context.Context, cfg *config.Config) error {
	return nil
}
