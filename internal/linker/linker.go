package linker

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/katallaxie/csync/internal/spec"

	"github.com/katallaxie/pkg/utils/files"
	cp "github.com/otiai10/copy"
)

type linker struct {
	opts *Opts
}

// Linker is the interface for linking, backing up and restoring files.
type Linker interface {
	// Backup creates a backup of the files in the app spec.
	Backup(context.Context, *spec.App, bool, bool) error
	// Restore restores the files in the app spec.
	Restore(context.Context, *spec.App, bool, bool) error
	// Unlink unlinks the files in the app spec.
	Unlink(context.Context, *spec.App, bool, bool) error
}

// Opt is the functional option for the linker.
type Opt func(*Opts)

// Opts ...
type Opts struct {
	HomeDir  string
	Provider *spec.Provider
	Verbose  bool
}

// Configure is configuring the linker.
func (o *Opts) Configure(opts ...Opt) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithProvider is setting the storage provider.
func WithProvider(p *spec.Provider) Opt {
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

// Backup is copying the files from the app spec to the backup directory
// and then linking the files from the app spec to the original location.
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

		if fi.Mode().IsDir() {
			err := cp.Copy(src, dst)
			if err != nil {
				return err
			}

			err = os.RemoveAll(src)
			if err != nil {
				return err
			}
		} else {
			// Copy file to backup directory ...
			_, err = files.CopyFile(src, dst, true)
			if err != nil {
				return err
			}

			// Delete source file
			err = os.Remove(src)
			if err != nil {
				return err
			}
		}

		// Create symlink from destination to source
		err = os.Symlink(dst, src)
		if err != nil {
			return err
		}
	}

	return nil
}

// Unlink is copying the files from the app spec to the original location
// and then unlinking the files from the backup directory.
//
// nolint:gocyclo
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

		fi, err := os.Lstat(src)
		if err != nil {
			return err
		}

		// try to delete and ignore any error
		_ = os.Remove(dst)

		if fi.Mode().IsDir() {
			err := cp.Copy(src, dst)
			if err != nil {
				return err
			}
		} else {
			_, err = files.CopyFile(src, dst, true)
			if err != nil {
				return err
			}
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
