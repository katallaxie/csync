package spec

// Bartender is the configuration for the Bartender.
func Bartender() App {
	return App{
		Name: "bartender",
		Files: Files{
			"Library/Preferences/com.surteesstudios.Bartender.plist",
			"Library/Preferences/com.surteesstudios.Bartender-setapp.plist",
			"Library/Application Support/Bartender/Bartender.BartenderPreferences",
		},
	}
}
