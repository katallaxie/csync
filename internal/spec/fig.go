package spec

// Fig represents for the fig config.
func Fig() App {
	return App{
		Name: "fig",
		Files: []string{
			"~/.fig/settings.json",
		},
	}
}
