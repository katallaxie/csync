package checker

import (
	"context"
	"fmt"
	"os/user"

	"github.com/katallaxie/csync/internal/config"

	"github.com/katallaxie/pkg/filex"
)

// UseableEnv is a check to see if the environment is useable.
func UseableEnv(_ context.Context, cfg *config.Config) error {
	user, err := user.Current()
	if err != nil {
		return err
	}

	if user.Uid == "0" && !cfg.Flags.Root {
		return fmt.Errorf("running as 'root' is dangerous. You can overwrite this behavior with '--root'")
	}

	return nil
}

// UseSetup is a check to see if the setup is useable.
func UseSetup(_ context.Context, cfg *config.Config) error {
	ok, _ := filex.FileExists(cfg.File)

	if !ok {
		return fmt.Errorf("%s does not exists. You can create a new config with 'init'", cfg.File)
	}

	// TODO (@katallaxie): disable this test, because there is not a clear use case here.
	// We may need to make this configurable in the spec file.
	//
	// path := filepath.Join(filepath.Dir(cfg.File), ".csync")
	// ok, _ = files.FileExists(path)

	// if !ok {
	// 	return fmt.Errorf("%s does not exists to store config files. You can create a new config with 'init'", path)
	// }

	return nil
}
