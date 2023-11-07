package spec

// Nano is the configuration file for `nano`.
func Nano() App {
	return App{
		Name: "nano",
		Files: []string{
			"~/.nanorc",
		},
	}
}
