package spec

// List of apps.
func List() []App {
	return []App{
		Docker(),
		AWS(),
	}
}