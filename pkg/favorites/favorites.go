package favorites

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/pelletier/go-toml"
)

type Favorites struct {
	animes []Anime
}

type Anime struct {
	Name   string  `toml:"anime"`
	Scenes []scene `toml:"scene"`
}

type scene struct {
	Name        string
	Season      string
	Episode     string
	Time        string
	Description string
}

func (favorites *Favorites) LoadSceneFiles(path string) {
	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()

		if strings.HasSuffix(fileName, ".toml") {
			scene := favorites.LoadSceneFile(fileName)
			favorites.animes = append(favorites.animes, scene)
		}
	}
}

func (favorites Favorites) LoadSceneFile(pathname string) Anime {
	var favorite Anime
	content, _ := ioutil.ReadFile(pathname)

	if err := toml.Unmarshal(content, &favorite); err != nil {
		log.Fatal(err)
	}

	return favorite
}

func (favorites Favorites) OutputScenes() {
	for _, a := range favorites.animes {
		if len(a.Name) > 0 {
			fmt.Println("Name: ", a.Name)
		}

		for _, s := range a.Scenes {
			fmt.Printf("%s (%s) - %s\n", s.Name, s.Episode, s.Time)
		}
	}
}
