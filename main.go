package main

import (
	"fmt"
	"os"

	"github.com/Soulsbane/goapp/pkg/cli"
)

var args struct {
	Quiet bool `arg:"-q"` // this flag is global to all subcommands
}

func main() {
	app := cli.NewCmdLineApp("Test App", "1.0", &args)

	app.PrintWarning("This is alpha software. Use at your own risk!")

	dir, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
	}

	loadSceneFiles(dir)
}
