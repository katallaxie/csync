package spec

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"

	"github.com/andersnormal/pkg/utils/files"
	s "github.com/andersnormal/pkg/utils/strings"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

var (
	allowedExt = []string{"yml", "yaml"}
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

	for _, in := range spec.Includes {
		ext := filepath.Ext(in)
		ok := s.Contains(allowedExt, ext)

		if !ok {
			continue
		}

		a, err := ioutil.ReadFile(in)
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

// FilePathFromProvider ...
func FilePathFromProvider(p *Provider, f string) string {
	return f
}
