package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/Soulsbane/goapp/pkg/cli"
)

/* Only useful if we store Duration in a value
type duration struct {
	time.Duration
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}*/

type scenes struct {
	Scenes []Scene `toml:"scene"`
}

type Scene struct {
	Name        string
	Season      string
	Episode     string
	Time        string // TODO: Make a Duration type instead?
	Description string
}

var args struct {
	Quiet bool `arg:"-q"` // this flag is global to all subcommands
}

func main() {
	app := cli.NewCmdLineApp("Test App", "1.0", &args)
	app.PrintWarning("blah")
	content, _ := ioutil.ReadFile("test.toml")

	var favorites scenes
	if _, err := toml.Decode(string(content), &favorites); err != nil {
		log.Fatal(err)
	}

	for _, s := range favorites.Scenes {
		fmt.Printf("%s (%s)\n", s.Name, s.Season)
	}
}
