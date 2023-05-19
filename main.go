package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/andersnormal/pkg/utils/files"
	"github.com/katallaxie/csync/pkg/checker"
	"github.com/katallaxie/csync/pkg/config"
	"github.com/katallaxie/csync/pkg/linker"
	"github.com/katallaxie/csync/pkg/spec"
	"mvdan.cc/sh/syntax"

	"github.com/spf13/pflag"
)

var version = ""

const usage = `Usage: csync [-crflvsdpw] [--config] [--restore] [--force] [--verbose] [--unlink] [--dry] [--validate] [--version]

'''
version: 1
storage: icloud	 
'''

Options:
`

// nolint:gocyclo
func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	cfg := config.New()

	err := cfg.InitDefaultConfig()
	if err != nil {
		log.Panic(err)
	}

	pflag.Usage = func() {
		log.Print(usage)
		pflag.PrintDefaults()
	}

	pflag.BoolVarP(&cfg.Flags.Verbose, "verbose", "v", cfg.Flags.Verbose, "verbose output")
	pflag.BoolVarP(&cfg.Flags.Help, "help", "h", cfg.Flags.Help, "show help")
	pflag.BoolVar(&cfg.Flags.Init, "init", cfg.Flags.Init, "init config")
	pflag.BoolVarP(&cfg.Flags.Force, "force", "f", cfg.Flags.Force, "force init")
	pflag.BoolVarP(&cfg.Flags.Dry, "dry", "d", cfg.Flags.Dry, "dry run")
	pflag.StringVarP(&cfg.File, "config", "c", cfg.File, "config file")
	pflag.BoolVarP(&cfg.Flags.Validate, "validate", "V", cfg.Flags.Validate, "validate config")
	pflag.BoolVar(&cfg.Flags.Unlink, "unlink", cfg.Flags.Unlink, "unlink")
	pflag.BoolVar(&cfg.Flags.Version, "version", cfg.Flags.Version, "version")
	pflag.BoolVar(&cfg.Flags.Restore, "restore", cfg.Flags.Version, "restore")
	pflag.BoolVar(&cfg.Flags.Root, "root", cfg.Flags.Root, "run as root")
	pflag.Parse()

	if cfg.Flags.Version {
		fmt.Printf("%s\n", getVersion())
		os.Exit(0)
	}

	if cfg.Flags.Help {
		pflag.Usage()
		os.Exit(0)
	}

	if cfg.Flags.Init {
		if cfg.Flags.Verbose {
			log.Printf("initializing config (%s)", cfg.File)
		}

		if err := spec.Write(spec.Default(), cfg.File, cfg.Flags.Force); err != nil {
			log.Fatal(err)
		}

		if cfg.Flags.Verbose {
			log.Printf("creating config folder (%s)", cfg.Path)
		}

		err = files.MkdirAll(cfg.Path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := checker.New(
		checker.WithChecks(checker.UseableEnv),
		checker.WithChecks(checker.UseSetup),
	)

	if err := c.Ready(ctx, cfg); err != nil {
		log.Panic(err)
	}

	s, err := cfg.LoadSpec()
	if err != nil {
		log.Panic(err)
	}

	if cfg.Flags.Validate {
		err = spec.Validate(s)
		if err != nil {
			log.Panic(err)
		}

		if cfg.Flags.Verbose {
			log.Print("OK")
		}

		return
	}

	_, _ = parseArgs()

	opts := []linker.Opt{linker.WithProvider(s.Provider)}
	if cfg.Flags.Verbose {
		opts = append(opts, linker.WithVerbose())
	}

	l := linker.New(opts...)

	if cfg.Flags.Restore {
		s.Lock()
		defer s.Unlock()

		if cfg.Flags.Verbose {
			log.Print("Restore files...")
		}

		for _, a := range s.Apps {
			a := a
			if cfg.Flags.Verbose {
				log.Printf("Restoring %s", a.Name)
			}

			if err := l.Restore(ctx, &a, cfg.Flags.Force, cfg.Flags.Dry); err != nil {
				log.Panic(err)
			}
		}

		return
	}

	if cfg.Flags.Unlink {
		s.Lock()
		defer s.Unlock()

		if cfg.Flags.Verbose {
			log.Printf("Unlink apps ...")
		}

		for _, a := range s.Apps {
			a := a
			if cfg.Flags.Verbose {
				log.Printf("Unlink '%s'", a.Name)
			}

			if err := l.Unlink(ctx, &a, cfg.Flags.Force, cfg.Flags.Dry); err != nil {
				log.Panic(err)
			}
		}

		return
	}

	if cfg.Flags.Verbose {
		log.Printf("Backup apps ....")
	}

	for _, a := range s.Apps {
		a := a
		s.Lock()
		defer s.Unlock()

		if cfg.Flags.Verbose {
			log.Printf("Backup '%s'", a.Name)
		}

		if err := l.Backup(ctx, &a, cfg.Flags.Force, cfg.Flags.Dry); err != nil {
			log.Panic(err)
		}
	}
}

func parseArgs() ([]string, []string) {
	args := pflag.Args()
	dashPos := pflag.CommandLine.ArgsLenAtDash()

	if dashPos == -1 {
		return args, []string{}
	}

	cliArgs := make([]string, 0)
	for _, arg := range args[dashPos:] {
		arg = syntax.QuotePattern(arg)
		cliArgs = append(cliArgs, arg)
	}

	return args[:dashPos], cliArgs
}

func getVersion() string {
	if version != "" {
		return version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" {
		return "unknown"
	}

	version = info.Main.Version
	if info.Main.Sum != "" {
		version += fmt.Sprintf(" (%s)", info.Main.Sum)
	}

	return version
}
