package spec

// Azure is the configuration for the Azure CLI.
func Azure() App {
	return App{
		Name: "azure",
		Files: Files{
			".azure",
		},
	}
}
