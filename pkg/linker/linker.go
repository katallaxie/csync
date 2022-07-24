package linker

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/andersnormal/pkg/utils/files"
	"github.com/katallaxie/csync/pkg/spec"
)

type linker struct {
	opts *Opts
}

// Linker ...
type Linker interface {
	Backup(context.Context, *spec.App, bool, bool) error
	Restore(context.Context, *spec.App, bool, bool) error
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	Provider *spec.Provider
}

// Configure ...
func (o *Opts) Configure(opts ...Opt) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithProvider ...
func WithProvider(p *spec.Provider) Opt {
	return func(o *Opts) {
		o.Provider = p
	}
}

// New ...
func New(opts ...Opt) Linker {
	options := new(Opts)
	options.Configure(opts...)

	return &linker{
		opts: options,
	}
}

// Backup ...
func (l *linker) Backup(ctx context.Context, app *spec.App, force bool, dry bool) error {
	for _, src := range app.Files {
		dst, err := l.FilePath(src)
		if err != nil {
			return err
		}

		dstfi, err := os.Lstat(dst)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}

		fi, err := os.Lstat(src)
		if err != nil {
			return err
		}

		if os.SameFile(dstfi, fi) {
			continue
		}

		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			continue // already is a symlink, needs force
		}

		// Copy file ...
		_, err = files.CopyFile(src, dst, true)
		if err != nil {
			return err
		}

		// Delete source file ...
		err = os.Remove(src)
		if err != nil {
			return err
		}

		// Create symlink ...
		err = os.Symlink(dst, src)
		if err != nil {
			return err
		}
	}

	return nil
}

// Restore ...
func (l *linker) Restore(ctx context.Context, app *spec.App, force bool, dry bool) error {
	for _, dst := range app.Files {
		_ = spec.FilePathFromProvider(l.opts.Provider, dst)
	}

	return nil
}

// FilePath ...
func (l *linker) FilePath(src string) (string, error) {
	name := strings.ToLower(l.opts.Provider.Name)
	path := "csync"

	if l.opts.Provider.Path != "" {
		path = l.opts.Provider.Path
	}

	var home string
	var err error
	switch name {
	case "icloud":
		home, err = ICloudFolder()
		if err != nil {
			return "", err
		}
	default:
		return "", fmt.Errorf("unknown provider")
	}

	return filepath.Join(home, path, src), nil
}
