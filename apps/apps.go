package apps

import "github.com/katallaxie/csync/internal/spec"

// List of apps.
func List() Apps {
	return []spec.App{
		Docker(),
		AWS(),
	}
}

// Apps is a list of apps.
type Apps []spec.App

// Len returns the length of the list.
func (a Apps) Len() int {
	return len(a)
}

// Filter returns a list of apps that match the given name.
func (a Apps) Filter(name string) Apps {
	apps := make(Apps, 0)

	for _, app := range List() {
		if app.Name == name {
			continue
		}

		apps = append(apps, app)
	}

	return apps
}
