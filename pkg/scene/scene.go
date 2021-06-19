package scene

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/pelletier/go-toml"
)

type FavoriteAnimeScene struct {
	animes []anime
}

type anime struct {
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

func (favorite *FavoriteAnimeScene) LoadSceneFiles(path string) {
	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()

		if strings.HasSuffix(fileName, ".toml") {
			scene := favorite.LoadSceneFile(fileName)
			favorite.animes = append(favorite.animes, scene)
		}
	}
}

func (favoriteScene FavoriteAnimeScene) LoadSceneFile(pathname string) anime {
	var favorite anime
	content, _ := ioutil.ReadFile(pathname)

	if err := toml.Unmarshal(content, &favorite); err != nil {
		log.Fatal(err)
	}

	return favorite
}

func (favoriteScene FavoriteAnimeScene) OutputScenes() {
	for _, a := range favoriteScene.animes {
		if len(a.Name) > 0 {
			fmt.Println("Name: ", a.Name)
		}

		for _, s := range a.Scenes {
			fmt.Printf("%s (%s) - %s\n", s.Name, s.Episode, s.Time)
		}
	}
}
