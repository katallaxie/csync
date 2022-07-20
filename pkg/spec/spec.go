package spec

import (
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

// Spec ...
type Spec struct {
	// Version ...
	Version int `validate:"required" yaml:"version"`
	// Storage ...
	Storage string `validate:"required" yaml:"storage"`
	// Apps ...
	Apps []App `yaml:"apps"`
	// Includes ...
	Includes []string `yaml:"includes"`
	// Excludes ...
	Excludes []string `yaml:"excludes"`
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
