package spec

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/andersnormal/pkg/utils/files"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

// Spec ...
type Spec struct {
	// Version ...
	Version int `validate:"required" yaml:"version"`
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
	Name string `validate:"required" yaml:"name"`
	Path string `yaml:"path"`
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
func (s *Spec) Validate() error {
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

	ok, _ := files.FileExists(file)
	if ok && !force {
		return fmt.Errorf("%s already exists, use --force to overwrite", file)
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}

// Load ...
func Load(file string) (*Spec, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var spec Spec
	err = yaml.Unmarshal(f, &spec)
	if err != nil {
		return nil, err
	}

	return &spec, nil
}

// FilePathFromProvider ...
func FilePathFromProvider(p *Provider, f string) string {
	return f
}
