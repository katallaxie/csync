package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/katallaxie/csync/internal/utils"
	"github.com/katallaxie/csync/pkg/proto"
	"github.com/katallaxie/pkg/filex"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

var validate = validator.New()

const (
	// DefaultDirectory is the default directory for the configuration file.
	DefaultDirectory = "csync"
	// DefaultPath is the default path for the configuration file.
	DefaultPath = ".csync"
	// DefaultFilename is the default filename for the configuration file.
	DefaultFilename = ".csync.yml"
)

// Spec is the configuration file for `csync`.
type Spec struct {
	// Version is the version of the configuration file.
	Version int `yaml:"version" validate:"required,eq=1"`
	// Path is the path to the configuration file.
	Path string `yaml:"path,omitempty"`
	// Provider is the configuration for the provider.
	Provider Provider `validate:"required" yaml:"provider"`
	// Apps is a list of apps to sync.
	Apps []App `yaml:"apps,omitempty"`
	// Includes is a list of apps to include.
	Includes []string `yaml:"includes,omitempty" validate:"required_with=Excludes"`
	// Excludes is a list of apps to exclude.
	Excludes []string `yaml:"excludes,omitempty" validate:"required_with=Includes"`
	// Stderr is the standard error output.
	Stderr bool `yaml:"stderr,omitempty"`
	// Stdout is the standard output.
	Stdout bool `yaml:"stdout,omitempty"`
	// FailOnMissing is the flag to fail on missing files.
	FailOnMissing bool `yaml:"fail_on_missing,omitempty"`

	sync.Mutex `yaml:"-"`
}

// UnmarshalYAML overrides the default unmarshaler for the spec.
func (s *Spec) UnmarshalYAML(data []byte) error {
	spec := struct {
		Version  int      `yaml:"version" validate:"required,eq=1"`
		Path     string   `yaml:"path,omitempty"`
		Provider Provider `yaml:"provider" validate:"required"`
		Apps     []App    `yaml:"apps,omitempty"`
		Includes []string `yaml:"includes,omitempty" validate:"required_with=Excludes"`
		Excludes []string `yaml:"excludes,omitempty" validate:"required_with=Includes"`
		Stderr   bool     `yaml:"stderr,omitempty"`
		Stdout   bool     `yaml:"stdout,omitempty"`
	}{
		Stderr: true,
		Stdout: true,
	}

	if err := yaml.Unmarshal(data, &spec); err != nil {
		return errors.WithStack(err)
	}

	s.Version = spec.Version
	s.Path = spec.Path
	s.Provider = spec.Provider
	s.Apps = spec.Apps
	s.Includes = spec.Includes
	s.Excludes = spec.Excludes

	return nil
}

// GetVersion returns the version of the configuration file.
func (s *Spec) GetVersion() int {
	return s.Version
}

// Default is the default configuration.
func Default() *Spec {
	return &Spec{
		Version: 1,
		Provider: Provider{
			Name: "icloud",
		},
	}
}

// Provider is the configuration for the file provider.
// This provider does support local, file-based storages.
type Provider struct {
	// Name ...
	Name string `validate:"required" yaml:"name"`
	// Pathh ...
	Path string `yaml:"path"`
	// Directory ...
	Directory string `yaml:"directory"`
}

// GetPath ...
func (s *Provider) GetPath() string {
	return s.Path
}

// GetDirectory ...
func (s *Provider) GetDirectory() string {
	return s.Directory
}

// GetName ...
func (s *Provider) GetName() string {
	return strings.ToLower(s.Name)
}

// GetApps reutrns the list of apps to sync.
func (s *Spec) GetApps(defaults ...App) []App {
	apps := s.Apps

	if len(s.Includes) == 0 {
		apps = append(apps, defaults...)
	}

	if len(s.Includes) > 0 {
		for _, in := range s.Includes {
			for _, app := range defaults {
				if app.Name == in {
					apps = append(apps, app)
				}
			}
		}
	}

	for _, ex := range s.Excludes {
		for i, app := range apps {
			if app.Name == ex {
				apps = append(apps[:i], apps[i+1:]...)
			}
		}
	}

	return apps
}

// GetFolder ...
func (p *Provider) GetFolder() (string, error) {
	dir := DefaultDirectory
	path := p.GetPath()

	if p.GetDirectory() != "" {
		dir = p.GetDirectory()
	}

	var base string
	var err error
	switch p.GetName() {
	case "file":
	case "dropbox":
		base, err = utils.DropboxFolder()
		if err != nil {
			return "", err
		}
	case "icloud":
		base, err = utils.ICloudFolder()
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("unknown provider")
	}

	// returns a fully resolved path for the backup of the files.
	return filepath.Join(base, path, dir), nil
}

// App is the configuration for the app.
type App struct {
	// Name ...
	Name string `yaml:"name"`
	// Files ...
	Files Files `yaml:"files"`
}

// ToProto ...
func (a *App) ToProto() *proto.Application {
	return &proto.Application{
		Name:  a.Name,
		Files: a.Files,
	}
}

// Files is the list of files to sync.
type Files []string

// Includes is the list of files to include.
type Includes []string

// Excludes is the list of files to exclude.
type Excludes []string

// Validate is the validation function for the spec.
func (s *Spec) Validate() error {
	validate = validator.New()

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("yaml"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(s)
	if err != nil {
		return err
	}

	return validate.Struct(s)
}

// Write is the write function for the spec.
func Write(s *Spec, file string, force bool) error {
	b, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	ok, _ := filex.FileExists(filepath.Clean(file))
	if ok && !force {
		return fmt.Errorf("%s already exists, use --force to overwrite", file)
	}

	f, err := os.Create(filepath.Clean(file))
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}
