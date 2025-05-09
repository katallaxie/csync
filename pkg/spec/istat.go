package spec

// Istat is the configuration for the iStat Menus app.
func Istat() App {
	return App{
		Name: "iStat Menus",
		Files: Files{
			"Library/Preferences/com.bjango.istatmenus.plist",
			"Library/Preferences/com.bjango.istatmenus.status.plist",
			"Library/Preferences/com.bjango.istatmenus5.extras.plist",
			"Library/Preferences/com.bjango.istatmenus6.extras.plist",
		},
	}
}
