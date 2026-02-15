package spec

// K9s is the configuration for the k9s Kubernetes CLI.
// See: https://k9scli.io/
func K9s() App {
	return App{
		Name: "k9s",
		Files: Files{
			"~/Library/Application Support/k9s/config.yml",
		},
	}
}
