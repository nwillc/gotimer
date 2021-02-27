package efs

import (
	"flag"
	"fmt"
	"os"
)

// EFlagSet is an extended flag.FlagSet.
type EFlagSet struct {
	*flag.FlagSet
	memberSets map[string][]string
}

// NewEFlagSet instantiate an EFlagSet with flag.ContinueOnError and default usage that enumerates MemberString choices.
func NewEFlagSet(name string) *EFlagSet {
	efs := EFlagSet{
		FlagSet:    flag.NewFlagSet(name, flag.ContinueOnError),
		memberSets: map[string][]string{},
	}
	efs.Usage = efs.defaultUsage
	return &efs
}

// ParseOsArgs runs Parse(string) on the targeted EFlagSet against the program arguments.
func (efs *EFlagSet) ParseOsArgs() {
	err := efs.Parse(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		os.Exit(2)
	}
}

// MemberString defines a string flag with specified name, default value, and usage string.
// It additionally provides a set of allowable values for that string to be a member of.
func (efs *EFlagSet) MemberString(name string, value *string, usage string, values []string, setterFn func(string)) *string {
	if !contains(values, *value) {
		panic(fmt.Sprintf("illegal default %s for %s", *value, name))
	}
	efs.memberSets[name] = values
	efs.Func(name, fmt.Sprintf("%s (default \"%s\")", usage, *value), func(s string) error {
		if !contains(values, s) {
			return fmt.Errorf("unkown %s %s", name, s)
		}
		setterFn(s)
		return nil
	})
	return value
}

func (efs *EFlagSet) defaultUsage() {
	if efs.Name() == "" {
		_, _ = fmt.Fprintf(efs.Output(), "Usage:\n")
	} else {
		_, _ = fmt.Fprintf(efs.Output(), "Usage of %s:\n", efs.Name())
	}
	efs.PrintDefaults()
	for k, v := range efs.memberSets {
		_, _ = fmt.Fprintf(efs.Output(), "\n  Allowable values for %s:\n", k)
		for _, s := range v {
			_, _ = fmt.Fprintln(efs.Output(), "\t", s)
		}
	}
}

func contains(values []string, member string) bool {
	for _, m := range values {
		if m == member {
			return true
		}
	}
	return false
}
