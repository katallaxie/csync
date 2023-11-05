package spec

// VSCode is the configuration for the VSCode.
func VSCode() App {
	return App{
		Name: "vscode",
		Files: Files{
			"Library/Application Support/Code/User/settings.json",
			"Library/Application Support/Code/User/keybindings.json",
			"Library/Application Support/Code/User/snippets",
		},
	}
}
