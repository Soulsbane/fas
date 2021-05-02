package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/Soulsbane/goapp/pkg/cli"
)

type duration struct {
	time.Duration
}
type scene struct {
	Name     string
	Duration duration
}
type songs struct {
	Scenes []scene `toml:"scene"`
}

type Scene struct {
	Name        string
	Season      string
	Episode     string
	Time        string
	Description string
}

func (d *duration) UnmarshalText(text []byte) error {
	var err error
	d.Duration, err = time.ParseDuration(string(text))
	return err
}

var args struct {
	Quiet bool `arg:"-q"` // this flag is global to all subcommands
}

func main() {
	app := cli.NewCmdLineApp("Test App", "1.0", &args)
	app.PrintWarning("blah")
	content, _ := ioutil.ReadFile("test.toml")

	var favorites songs
	if _, err := toml.Decode(string(content), &favorites); err != nil {
		log.Fatal(err)
	}

	for _, s := range favorites.Scenes {
		fmt.Printf("%s (%s)\n", s.Name, s.Duration)
	}
}
