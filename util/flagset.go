package util

import (
	"flag"
	"fmt"
	"os"
)

type FlagSet struct {
	*flag.FlagSet
}

func NewFlagSet(name string) *FlagSet {
	flags := &FlagSet{
		FlagSet: flag.NewFlagSet("ghbin-mgr "+name, flag.ExitOnError),
	}
	return flags
}

func (f *FlagSet) ValidateStringNotEmpty(str, errMsg string) {
	if str == "" {
		fmt.Printf("err: %s\n\n", errMsg)
		f.Usage()
		os.Exit(1)
	}
}

func (f *FlagSet) ParseCmdFlags() {
	if err := f.Parse(os.Args[2:]); err != nil {
		panic(err)
	}
}
