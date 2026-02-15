package spec

// GH is the configuration for the GitHub CLI.
// See: https://cli.github.com/
func GH() App {
	return App{
		Name: "gh",
		Files: Files{
			"~/.config/gh/hosts.yml",
			"~/.config/gh/config.yml",
		},
	}
}
