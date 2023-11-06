package spec

// MacOS is the configuration for the macOS.
func MacOS() App {
	return App{
		Name: "macOS",
		Files: Files{
			"~/Library/Preferences/com.apple.Terminal.plist",
			"~/Library/PDF Services",
			"~/Library/Preferences/com.apple.symbolichotkeys.plist",
			"~/Library/Scripts",
			"~/Library/Speech/Speakable Items",
			"~/Library/Workflows",
			"~/.MacOSX",
		},
	}
}
