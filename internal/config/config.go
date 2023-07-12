package config

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/katallaxie/csync/internal/spec"
)

// Flags ...
type Flags struct {
	Dry     bool
	Force   bool
	Root    bool
	Verbose bool
	Version bool
}

// Config ...
type Config struct {
	// Verbose toggles the verbosity
	Verbose bool
	// File...
	File string
	// Path ...
	Path string
	// FileMode ...
	FileMode os.FileMode
	// Flags ...
	Flags Flags
	// Stdin ...
	Stdin *os.File
	// Stdout ...
	Stdout *os.File
	// Stderr ...
	Stderr *os.File
}

// New ...
func New() *Config {
	return &Config{
		File:   ".csync.yml",
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
}

// InitDefaultConfig() ...
func (c *Config) InitDefaultConfig() error {
	cwd, err := c.Cwd()
	if err != nil {
		return err
	}
	c.File = filepath.Join(cwd, c.File)

	usr, err := user.Current()
	if err == nil {
		c.Path = filepath.Join(usr.HomeDir, spec.DefaultPath)
	}

	return nil
}

// HomeDir ...
func (c *Config) HomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, err
}

// Cwd ...
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}

// SpecFile ...
func (c *Config) LoadSpec() (*spec.Spec, error) {
	s, err := spec.Load(c.File)
	if err != nil {
		return nil, err
	}

	return s, nil
}
