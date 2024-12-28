package spec

// Terminal is the configuration for the terminal.
func Terminal() App {
	return App{
		Name: "Terminal",
		Files: Files{
			"~/Library/Preferences/com.apple.Terminal.plist",
		},
	}
}
