package apps

import "github.com/katallaxie/csync/internal/spec"

// AWS is the configuration for the AWS provider.
func AWS() spec.App {
	return spec.App{
		Name: "aws",
		Files: spec.Files{
			".aws",
		},
	}
}
