package spec

// Homebrew is the configuration for the Homebrew package manager.
func Homebrew() App {
	return App{
		Name: "homebrew",
		Files: Files{
			".Bewfile",
		},
	}
}
