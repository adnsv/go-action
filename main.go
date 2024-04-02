package main

import "github.com/alecthomas/kong"

var cli struct {
	Gitstat gitstat          `cmd:"" help:"Collect statistics on git repository."`
	Version kong.VersionFlag `short:"v" help:"Print version information and quit."`
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("go-action"),
		kong.Description("Git action helper utility."),
		kong.UsageOnError(),
		kong.Vars{"version": app_version()},
	)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
