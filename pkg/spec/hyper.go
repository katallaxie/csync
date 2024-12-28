package spec

// Hyper is the configuration for the Hyper terminal.
func Hyper() App {
	return App{
		Name: "hyper",
		Files: Files{
			"~/.hyper.js",
		},
	}
}
