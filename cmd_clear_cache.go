package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/shimberger/gohls/hls"
)

type clearCmd struct{}

func (*clearCmd) Name() string     { return "clear" }
func (*clearCmd) Synopsis() string { return "Clears all caches and temporary files" }
func (*clearCmd) Usage() string {
	return `clear:
  Clears the caches
`
}

func (p *clearCmd) SetFlags(f *flag.FlagSet) {}

func (p *clearCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	hls.ClearCache()
	return subcommands.ExitSuccess
}
