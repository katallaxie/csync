package files

import (
	"context"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/katallaxie/csync/pkg/homedir"
	p "github.com/katallaxie/csync/pkg/provider"
	"github.com/katallaxie/csync/pkg/spec"
	"github.com/katallaxie/pkg/utils/files"

	cp "github.com/otiai10/copy"
)

type provider struct {
	folder  string
	homedir string

	p.Unimplemented
}

var _ p.Provider = (*provider)(nil)

// Opt is the functional option for the provider.
type Opt func(*provider)

// WithFolder is configuring a specific folder for the provider.
func WithFolder(folder string) Opt {
	return func(p *provider) {
		p.folder = folder
	}
}

// WithHomeDir is configuring a specific home directory for the provider.
func WithHomeDir(homedir string) Opt {
	return func(p *provider) {
		p.homedir = homedir
	}
}

// Configure is configuring a set of options of the provider.
func (p *provider) Configure(opts ...Opt) {
	for _, o := range opts {
		o(p)
	}
}

// New ...
func New(opts ...Opt) *provider {
	p := new(provider)
	p.Configure(opts...)

	return p
}

// Backup a file.
// nolint:gocyclo
func (p *provider) Backup(ctx context.Context, app *spec.App, opts *p.Opts) error {
	for _, src := range app.Files {
		dst, err := FilePath(src, p.folder)
		if err != nil {
			return err
		}

		dstfi, err := os.Lstat(dst)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}

		src, err := homedir.Expand(src)
		if err != nil {
			return err
		}

		if ok, _ := files.FileNotExists(src); ok {
			continue
		}

		fi, err := os.Lstat(src)
		if err != nil {
			return err
		}

		if os.SameFile(dstfi, fi) {
			continue
		}

		if fi.Mode()&os.ModeSymlink == os.ModeSymlink && !opts.Force {
			continue // already is a symlink, needs force
		}

		log.Printf("Link '%s' => '%s'", src, dst)

		if opts.Dry {
			continue
		}

		if fi.Mode().IsDir() {
			if ok, _ := files.FileExists(dst); ok && opts.Force {
				err = os.RemoveAll(dst)
				if err != nil {
					return err
				}
			}

			err = cp.Copy(src, dst)
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

// Restore a file.
func (p *provider) Restore(ctx context.Context, app *spec.App, opts *p.Opts) error {
	for _, src := range app.Files {
		if ok, _ := files.FileNotExists(src); !ok {
			continue
		}

		dst, err := files.PathTransform(src, files.ExpandHomeFolder, files.ExpandHomeFolder)
		if err != nil {
			return err
		}

		if ok, _ := files.FileNotExists(dst); ok {
			continue
		}

		log.Printf("Restore %s from %s", src, dst)

		if opts.Dry {
			continue
		}

		// Create symlink ...
		err = os.Symlink(src, dst)
		if err != nil {
			return err
		}
	}

	return nil
}

// Unlink is unlinking files from the backup folder.
// nolint:gocyclo
func (p *provider) Unlink(ctx context.Context, app *spec.App, opts *p.Opts) error {
	for _, dst := range app.Files {
		dstfi, err := homedir.Expand(dst)
		if err != nil {
			return err
		}

		src, err := FilePath(dst, p.folder)
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

		log.Printf("Unlink %s from %s", dstfi, src)

		if opts.Dry {
			continue
		}

		fi, err := os.Lstat(src)
		if err != nil {
			return err
		}

		// try to delete and ignore any error
		_ = os.Remove(dstfi)

		if fi.Mode().IsDir() {
			err := cp.Copy(src, dstfi)
			if err != nil {
				return err
			}
		} else {
			_, err = files.CopyFile(src, dstfi, true)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Link ...
func (p *provider) Link(ctx context.Context, app *spec.App, opts *p.Opts) error {
	// this is not implemented with the file provider right now,
	// because the file provider does this in the backup phase.
	return nil
}

// FilePath ...
func FilePath(src, folder string) (string, error) {
	src = filepath.Clean(src)

	if strings.HasPrefix(src, "~/") {
		src = filepath.Join(folder, src[2:])
	}

	return src, nil
}
