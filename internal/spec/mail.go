package spec

// Mail is the configuration for the Mail app.
func Mail() App {
	return App{
		Name: "Mail",
		Files: Files{
			"Library/Preferences/com.apple.mail.plist",
		},
	}
}
