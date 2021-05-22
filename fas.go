package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
	"github.com/saracen/walker"
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

func loadScenes(dir string) {
	walkFn := func(pathname string, fi os.FileInfo) error {
		if strings.HasSuffix(pathname, ".toml") {
			processTOML(pathname)
		}

		return nil
	}

	errorCallbackOption := walker.WithErrorCallback(func(pathname string, err error) error {
		if os.IsPermission(err) {
			return nil
		}

		return err
	})

	walker.Walk(dir, walkFn, errorCallbackOption)
}

func processTOML(pathname string) {
	var favorite anime
	content, _ := ioutil.ReadFile("test.toml")

	if err := toml.Unmarshal(content, &favorite); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Anime: ", favorite.Anime)

	for _, s := range favorite.Scenes {
		fmt.Printf("%s (%s) - %s\n", s.Name, s.Episode, s.Time)
	}
}
