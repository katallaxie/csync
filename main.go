package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/debug"
	"time"

	"github.com/katallaxie/csync/pkg/config"
	"mvdan.cc/sh/syntax"

	"github.com/spf13/pflag"
)

var (
	version = ""
)

const usage = `Usage: csync [-cflvsdpw] [--config] [--force] [--verbose] [--dry] [--validate] [--var] [--version]

'''
version: 1
storage: icloud	 
'''

Options:
`

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	cfg := config.New()

	err := cfg.InitDefaultConfig()
	if err != nil {
		log.Fatal(err)
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
	pflag.BoolVar(&cfg.Flags.Version, "version", cfg.Flags.Version, "version")
	pflag.Parse()

	if cfg.Flags.Version {
		fmt.Printf("%s\n", getVersion())
		return
	}

	_, _, err = parseArgs()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

func parseArgs() ([]string, []string, error) {
	args := pflag.Args()
	dashPos := pflag.CommandLine.ArgsLenAtDash()

	if dashPos == -1 {
		return args, []string{}, nil
	}

	cliArgs := make([]string, 0)
	for _, arg := range args[dashPos:] {
		arg = syntax.QuotePattern(arg)
		cliArgs = append(cliArgs, arg)
	}

	return args[:dashPos], cliArgs, nil
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
