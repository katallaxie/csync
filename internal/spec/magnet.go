package spec

// Magnet is the configuration for the Magnet app.
func Magnet() App {
	return App{
		Name: "Magnet",
		Files: Files{
			"~/Library/Preferences/com.crowdcafe.windowmagnet.plist",
		},
	}
}
