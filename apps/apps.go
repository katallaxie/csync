package apps

import "github.com/katallaxie/csync/internal/spec"

// List of apps.
func List() []spec.App {
	return []spec.App{
		Docker(),
		AWS(),
	}
}
