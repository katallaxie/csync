package spec

// Ice is the configuration for the Ice menu bar.
// See: https://github.com/jordanbaird/Ice
func IceMenuBar() App {
	return App{
		Name: "icemenubar",
		Files: Files{
			"~/Library/Preferences/com.jordanbaird.Ice.plist",
		},
	}
}
