package spec

// Kubectl is the configuration for the kubectl.
func Kubectl() App {
	return App{
		Name: "kubectl",
		Files: Files{
			"~/.kube/config",
		},
	}
}
