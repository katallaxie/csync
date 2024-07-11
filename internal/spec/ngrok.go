package spec

// Ngrok is the configuration for `ngrok“.
func Ngrok() App {
	return App{
		Name: "ngrok",
		Files: []string{
			"~/.ngrok",
			"~/.ngrok2",
		},
	}
}
