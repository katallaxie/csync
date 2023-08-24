package spec

// AWS is the configuration for the AWS provider.
func AWS() App {
	return App{
		Name: "aws",
		Files: Files{
			".aws",
		},
	}
}
