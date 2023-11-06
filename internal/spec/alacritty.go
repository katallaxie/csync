package spec

// Alacritty is the configuration for the Alacritty terminal.
func Alacritty() App {
	return App{
		Name: "alacritty",
		Files: Files{
			"~/.config/alacritty",
			"~/.alacritty.yml",
		},
	}
}
