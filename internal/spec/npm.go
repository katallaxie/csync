package spec

// Npm is the configuration for the npm.
func Npm() App {
	return App{
		Name: "npm",
		Files: Files{
			"~/.npmrc",
		},
	}
}
