package spec

// Zed is the configuration for the Zed Editor.
// See: https://github.com/zed-industries/zed
func Zed() App {
	return App{
		Name: "zed",
		Files: Files{
			"~/.config/zed",
			"~/.zed",
		},
	}
}
