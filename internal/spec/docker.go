package spec

// Docker ...
func Docker() App {
	return App{
		Name: "docker",
		Files: Files{
			".docker/config.json",
			".docker/daemon.json",
		},
	}
}
