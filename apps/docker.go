package apps

import "github.com/katallaxie/csync/internal/spec"

// Docker ...
func Docker() spec.App {
	return spec.App{
		Name: "docker",
		Files: spec.Files{
			".docker/config.json",
			".docker/daemon.json",
		},
	}
}
