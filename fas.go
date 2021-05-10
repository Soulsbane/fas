package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/Soulsbane/goapp/pkg/cli"
	"github.com/pelletier/go-toml"
)

type anime struct {
	Anime  string  `toml:"anime"`
	Scenes []Scene `toml:"scene"`
}

type Scene struct {
	Name        string
	Season      string
	Episode     string
	Time        string
	Description string
}

var args struct {
	Quiet bool `arg:"-q"` // this flag is global to all subcommands
}

func main() {
	var favorites anime
	app := cli.NewCmdLineApp("Test App", "1.0", &args)

	app.PrintWarning("This is alpha software. Use at your own risk!")
	content, _ := ioutil.ReadFile("test.toml")

	if err := toml.Unmarshal(content, &favorites); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Anime: ", favorites.Anime)
	for _, s := range favorites.Scenes {
		fmt.Printf("%s (%s) - %s\n", s.Name, s.Episode, s.Time)
	}
}
