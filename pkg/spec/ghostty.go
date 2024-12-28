package spec

// Ghostty is the configuration for the Ghostty terminal.
//
// See: https://github.com/ghostty-org/ghostty
func Ghostty() App {
	return App{
		Name: "ghostty",
		Files: Files{
			"~/.config/ghostty",
		},
	}
}
