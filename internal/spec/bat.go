package spec

// Bat is the configuration for the bat pager.
func Bat() App {
	return App{
		Name: "bat",
		Files: Files{
			"bat/config",
			"bat/themes",
			"bat/themes",
		},
	}
}
