package spec

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/andersnormal/pkg/utils/files"
	s "github.com/andersnormal/pkg/utils/strings"
	"github.com/go-playground/validator/v10"
	"github.com/katallaxie/csync/pkg/utils"
	"gopkg.in/yaml.v3"
)

const (
	DefaultDirectory = "csync"
	DefaultPath      = ".csync"
	DefaultFilename  = ".csync.yml"
)

var allowedExt = []string{"yml", "yaml"}

// Spec ...
type Spec struct {
	// Version ...
	Version int `validate:"required" yaml:"version"`
	// Path ...
	Path string `yaml:"path,omitempty"`
	// Provider ...
	Provider Provider `validate:"required" yaml:"provider"`
	// Apps ...
	Apps []App `yaml:"apps,omitempty"`
	// Includes ...
	Includes []string `yaml:"includes,omitempty"`
	// Excludes ...
	Excludes []string `yaml:"excludes,omitempty"`

	sync.Mutex
}

// Default ...
func Default() *Spec {
	return &Spec{
		Version: 1,
		Provider: Provider{
			Name: "icloud",
		},
	}
}

// Provider ...
type Provider struct {
	// Name ...
	Name string `validate:"required" yaml:"name"`
	// Pathh ...
	Path string `yaml:"path"`
	// Directory ...
	Directory string `yaml:"directory"`
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
		base, err = utils.DropboxFodler()
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

// App ...
type App struct {
	// Name ...
	Name string `yaml:"name"`
	// Files ...
	Files Files `yaml:"files"`
}

// Files ...
type Files []string

// Includes ...
type Includes []string

// Excludes ...
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
