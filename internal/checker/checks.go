package checker

import (
	"context"
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/katallaxie/csync/internal/config"
	"github.com/katallaxie/pkg/utils/files"
)

// UseableEnv ...
func UseableEnv(ctx context.Context, cfg *config.Config) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	if user.Uid == "0" && !cfg.Flags.Root {
		return fmt.Errorf("running as 'root' is dangerous. You can overwrite this behavior with '--root'")
	}

	return nil
}

// UseSetup ...
func UseSetup(ctx context.Context, cfg *config.Config) error {
	ok, _ := files.FileExists(cfg.File)

	if !ok {
		return fmt.Errorf("%s does not exists. You can create a new config with 'init'", cfg.File)
	}

	path := filepath.Join(filepath.Dir(cfg.File), ".csync")
	ok, _ = files.FileExists(path)

	if !ok {
		return fmt.Errorf("%s does not exists to store config files. You can create a new config with 'init'", path)
	}

	return nil
}
