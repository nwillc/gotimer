package setup

import (
	"flag"
	"fmt"
	"os"
)

type CliFlagSet struct {
	*flag.FlagSet
}

func NewCliFlagSet(name string) *CliFlagSet {
	return &CliFlagSet{
		FlagSet: flag.NewFlagSet(name, flag.ContinueOnError),
	}
}

func (cfs *CliFlagSet) ParseOsArgs() {
	err := cfs.Parse(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		os.Exit(2)
	}
}

func (cfs *CliFlagSet) MemberString(name string, value *string, usage string, memberFn func(string) bool, setterFn func(string)) *string {
	cfs.Func(name, fmt.Sprintf("%s (default \"%s\")", usage, *value), func(s string) error {
		if !memberFn(s) {
			return fmt.Errorf("unkown %s", name)
		}
		setterFn(s)
		return nil
	})
	if ok := memberFn(*value); !ok {
		panic(fmt.Sprintf("illegal default %s for %s", *value, name))
	}
	return value
}
