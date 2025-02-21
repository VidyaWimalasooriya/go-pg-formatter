package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/pflag"
)

func getConfig(flags *pflag.FlagSet) []string {
	config := configurationType{
		Anonymize:     setDefault(flags.Changed("anonymize"), false),
		CommaBreak:    setDefault(flags.Changed("comma-break"), false),
		FunctionCase:  setDefault(getInt(flags, "function-case"), 0),
		KeywordCase:   setDefault(getInt(flags, "keyword-case"), 2),
		NoRcFile:      setDefault(flags.Changed("no-rcfile"), false),
		Placeholder:   setDefault(getString(flags, "placeholder"), ""),
		Spaces:        setDefault(getInt(flags, "spaces"), 4),
		StripComments: setDefault(flags.Changed("nocomment"), false),
		Tabs:          setDefault(flags.Changed("tabs"), false),
	}

	return mapArgs(config)
}

func mapArgs(config configurationType) []string {
	var args []string

	if config.Anonymize {
		args = append(args, "--anonymize")
	}

	if config.FunctionCase < 4 && config.FunctionCase >= 0 {
		args = append(args, "--function-case", strconv.Itoa(config.FunctionCase))
	}

	if config.KeywordCase < 4 && config.KeywordCase >= 0 {
		args = append(args, "--keyword-case", strconv.Itoa(config.KeywordCase))
	}

	if config.NoRcFile {
		args = append(args, "--no-rcfile")
	}

	if config.Placeholder != "" {
		args = append(args, "--placeholder", config.Placeholder)
	}

	if config.Spaces > 0 {
		args = append(args, "--spaces", strconv.Itoa(config.Spaces))
	}

	if config.StripComments {
		args = append(args, "--nocomment")
	}

	if config.Tabs {
		args = append(args, "--tabs")
	}

	if config.CommaBreak {
		args = append(args, "--comma-break")
	}

	return args
}

func setDefault[T comparable](value, defaultValue T) T {
	var zero T
	if value != zero {
		return value
	}

	return defaultValue
}

func getInt(flags *pflag.FlagSet, name string) int {
	val, err := flags.GetInt(name)
	if err != nil {
		log.Fatalf("Error retrieving %s flag: %v", name, err)
	}

	return val
}

func getString(flags *pflag.FlagSet, name string) string {
	val, err := flags.GetString(name)
	if err != nil {
		log.Fatalf("Error retrieving %s flag: %v", name, err)
	}

	return val
}
