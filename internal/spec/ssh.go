package spec

// SSH is the configuration for the SSH client.
func SSH() App {
	return App{
		Name: "ssh",
		Files: Files{
			"~/.ssh/config",
			"~/.ssh/authorized_keys",
		},
	}
}
