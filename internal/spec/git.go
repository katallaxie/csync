package spec

// Git is the configuration for the git version control system.
func Git() App {
	return App{
		Name: "git",
		Files: Files{
			"~/.gitconfig",
		},
	}
}
