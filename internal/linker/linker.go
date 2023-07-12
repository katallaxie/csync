package linker

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/katallaxie/csync/internal/spec"

	"github.com/katallaxie/pkg/utils/files"
)

type linker struct {
	opts *Opts
}

// Linker ...
type Linker interface {
	Backup(context.Context, *spec.App, bool, bool) error
	Restore(context.Context, *spec.App, bool, bool) error
	Unlink(context.Context, *spec.App, bool, bool) error
}

// Opt ...
type Opt func(*Opts)

// Opts ...
type Opts struct {
	HomeDir  string
	Provider spec.Provider
	Verbose  bool
}

// Configure ...
func (o *Opts) Configure(opts ...Opt) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithProvider ...
func WithProvider(p spec.Provider) Opt {
	return func(o *Opts) {
		o.Provider = p
	}
}

// WithHomedir ...
func WithHomedir(dir string) Opt {
	return func(o *Opts) {
		o.HomeDir = dir
	}
}

// WithVerbose ...
func WithVerbose() Opt {
	return func(o *Opts) {
		o.Verbose = true
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
//
//nolint:gocyclo
func (l *linker) Backup(ctx context.Context, app *spec.App, force bool, dry bool) error {
	for _, src := range app.Files {
		dst, err := l.opts.Provider.GetFilePath(src)
		if err != nil {
			return err
		}

		dstfi, err := os.Lstat(dst)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}

		src, err := files.ExpandHomeFolder(src)
		if err != nil {
			return err
		}

		fi, err := os.Lstat(src)
		if err != nil {
			return err
		}

		if os.SameFile(dstfi, fi) {
			continue
		}

		if fi.Mode()&os.ModeSymlink == os.ModeSymlink && !force {
			continue // already is a symlink, needs force
		}

		if l.opts.Verbose {
			log.Printf("Link '%s' => '%s'", src, dst)
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

// Unlink ...
func (l *linker) Unlink(ctx context.Context, app *spec.App, force bool, dry bool) error {
	for _, dst := range app.Files {
		dst, err := files.ExpandHomeFolder(dst)
		if err != nil {
			return err
		}

		src, err := l.opts.Provider.GetFilePath(dst)
		if err != nil {
			return err
		}

		not, err := files.FileNotExists(src)
		if err != nil {
			return err
		}

		if not {
			continue
		}

		if l.opts.Verbose {
			log.Printf("Unlink %s from %s", dst, src)
		}

		// try to delete and ignore any error
		_ = os.Remove(dst)

		_, err = files.CopyFile(src, dst, true)
		if err != nil {
			return err
		}
	}

	return nil
}

// Restore ...
func (l *linker) Restore(ctx context.Context, app *spec.App, force bool, dry bool) error {
	for _, src := range app.Files {
		dst, err := files.PathTransform(src, files.ExpandHomeFolder, files.ExpandHomeFolder)
		if err != nil {
			return err
		}

		ok, err := files.FileNotExists(dst)
		if err != nil {
			return err
		}

		if !ok {
			continue
		}

		if l.opts.Verbose {
			log.Printf("restore %s from %s", src, dst)
		}

		// Create symlink ...
		err = os.Symlink(src, dst)
		if err != nil {
			return err
		}
	}

	return nil
}
