package spec

// Bash is the configuration for the Bash shell.
func Bash() App {
	return App{
		Name: "bash",
		Files: Files{
			".bash_history",
			".bash_profile",
			".bashrc",
			".inputrc",
			".profile",
			".bash_logout",
			".bash_login",
			".inputrc",
			".bash_aliases",
		},
	}
}
