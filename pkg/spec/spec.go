package spec

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/andersnormal/pkg/utils"
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
	Apps []App `yaml:"apps,omitempty"`
	// Includes ...
	Includes []string `yaml:"includes,omitempty"`
	// Excludes ...
	Excludes []string `yaml:"excludes,omitempty"`

	sync.Mutex
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

	ok, _ := utils.FileExists(file)
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
