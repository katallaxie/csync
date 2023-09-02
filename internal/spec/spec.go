package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/katallaxie/csync/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/katallaxie/pkg/utils/files"
	s "github.com/katallaxie/pkg/utils/strings"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

const (
	DefaultDirectory = "csync"
	DefaultPath      = ".csync"
	DefaultFilename  = ".csync.yml"
)

var allowedExt = []string{"yml", "yaml"}

// Spec is the configuration file for `csync`.
type Spec struct {
	// Version is the version of the configuration file.
	Version int `yaml:"version" validate:"required"`
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

	sync.Mutex `yaml:"-"`
}

// UnmarshalYAML overrides the default unmarshaler for the spec.
func (s *Spec) UnmarshalYAML(data []byte) error {
	spec := struct {
		Version  int      `yaml:"version" validate:"required"`
		Path     string   `yaml:"path,omitempty"`
		Provider Provider `yaml:"provider" validate:"required"`
		Apps     []App    `yaml:"apps,omitempty"`
		Includes []string `yaml:"includes,omitempty" validate:"required_with=Excludes"`
		Excludes []string `yaml:"excludes,omitempty" validate:"required_with=Includes"`
	}{}

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

// Default is the default configuration.
func Default() *Spec {
	return &Spec{
		Version: 1,
		Provider: Provider{
			Name: "icloud",
		},
		Apps: List(),
	}
}

// Provider is the configuration for the provider.
type Provider struct {
	// Name ...
	Name string `validate:"required" yaml:"name"`
	// Pathh ...
	Path string `yaml:"path"`
	// Directory ...
	Directory string `yaml:"directory"`
}

// GetVersion returns the version of the configuration file.
func (s *Spec) GetVersion() int {
	return s.Version
}

// GetApps reutrns the list of apps to sync.
func (s *Spec) GetApps(defaults ...App) []App {
	apps := make([]App, 0)

	if len(s.Includes) == 0 {
		apps = append(apps, defaults...)
	}

	for _, in := range s.Includes {
		for _, app := range defaults {
			if app.Name == in {
				apps = append(apps, app)
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

	return s.Apps
}

// GetName ...
func (p Provider) GetName() string {
	return strings.ToLower(p.Name)
}

// GetPath ...
func (p Provider) GetPath() string {
	return p.Path
}

// GetDirectory ...
func (p Provider) GetDirectory() string {
	return p.Directory
}

// GetFilePath ...
func (p Provider) GetFilePath(f string) (string, error) {
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

	return filepath.Join(base, path, dir, f), nil
}

// App is the configuration for the app.
type App struct {
	// Name ...
	Name string `yaml:"name"`
	// Files ...
	Files Files `yaml:"files"`
}

// Files is the list of files to sync.
type Files []string

// Includes is the list of files to include.
type Includes []string

// Excludes is the list of files to exclude.
type Excludes []string

// Validate ..
func Validate(s *Spec) error {
	v := validator.New()

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("yaml"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := v.Struct(s)
	if err != nil {
		return err
	}

	return v.Struct(s)
}

// Write ...
func Write(s *Spec, file string, force bool) error {
	b, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	ok, _ := files.FileExists(filepath.Clean(file))
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

// Load ...
func Load(file string) (*Spec, error) {
	f, err := os.ReadFile(filepath.Clean(file))
	if err != nil {
		return nil, err
	}

	var spec Spec
	err = yaml.Unmarshal(f, &spec)
	if err != nil {
		return nil, err
	}

	err = Validate(&spec)
	if err != nil {
		return nil, err
	}

	for _, in := range spec.Includes {
		ext := filepath.Ext(in)
		ok := s.Contains(allowedExt, ext)

		if !ok {
			continue
		}

		a, err := os.ReadFile(filepath.Clean(in))
		if err != nil {
			return nil, err
		}

		var app App
		err = yaml.Unmarshal(a, &app)
		if err != nil {
			return nil, err
		}

		spec.Apps = append(spec.Apps, app)
	}

	return &spec, nil
}
