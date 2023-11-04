package spec

// Zsh is the configuration for the Zsh shell.
func Zsh() App {
	return App{
		Name: "zsh",
		Files: Files{
			".zshrc",
			".zshenv",
			".zprofile",
			".zlogin",
			".zlogout",
		},
	}
}
