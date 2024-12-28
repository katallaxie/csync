package spec

// Raycast represents for the raycast config.
func Raycast() App {
	return App{
		Name: "Raycast",
		Files: []string{
			"~/Library/Preferences/com.raycast.macos.plist",
		},
	}
}
