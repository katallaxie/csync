package spec

// Tmux is the configuration for the tmux.
func Tmux() App {
	return App{
		Name: "tmux",
		Files: Files{
			".tmux.conf",
		},
	}
}
