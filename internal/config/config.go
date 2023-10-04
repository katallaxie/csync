package config

import (
	"os"
	"os/user"
	"path/filepath"
	"sync"

	"github.com/katallaxie/csync/internal/spec"
)

// Flags contains the command line flags.
type Flags struct {
	// Dry toggles the dry run mode.
	Dry bool
	// Force toggles the force mode.
	Force bool
	// Root runs the command as root.
	Root bool
	// Verbose toggles the verbosity.
	Verbose bool
	// Version toggles the version.
	Version bool
}

// NewFlags returns a new flags.
func NewFlags() *Flags {
	return &Flags{}
}

// Config contains the configuration.
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
	Flags *Flags
	// Stdin ...
	Stdin *os.File
	// Stdout ...
	Stdout *os.File
	// Stderr ...
	Stderr *os.File
	// Spec ...
	Spec *spec.Spec

	sync.RWMutex
}

// New returns a new config.
func New() *Config {
	return &Config{
		File:   ".csync.yml",
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Flags:  &Flags{},
		Spec:   spec.Default(),
	}
}

// InitDefaultConfig initializes the default configuration.
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

// HomeDir returns the home directory.
func (c *Config) HomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	return usr.HomeDir, err
}

// Cwd returns the current working directory.
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}

// LoadSpec is a helper to load the spec from the config file.
func (c *Config) LoadSpec() error {
	f, err := os.ReadFile(filepath.Clean(c.File))
	if err != nil {
		return err
	}

	return c.Spec.UnmarshalYAML(f)
}
