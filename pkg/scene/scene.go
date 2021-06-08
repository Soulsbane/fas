package scene

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/pelletier/go-toml"
)

type FavoriteAnimeScene struct{}

type anime struct {
	Anime  string  `toml:"anime"`
	Scenes []scene `toml:"scene"`
}

type scene struct {
	Name        string
	Season      string
	Episode     string
	Time        string
	Description string
}

func (favorite FavoriteAnimeScene) LoadSceneFiles(path string) {
	files, err := ioutil.ReadDir(".")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileName := file.Name()

		if strings.HasSuffix(fileName, ".toml") {
			favorite.LoadSceneFile(fileName)
		}
	}
}

func (favoriteScene FavoriteAnimeScene) LoadSceneFile(pathname string) {
	var favorite anime
	content, _ := ioutil.ReadFile(pathname)

	if err := toml.Unmarshal(content, &favorite); err != nil {
		log.Fatal(err)
	}

	if len(favorite.Anime) > 0 {
		fmt.Println("Anime: ", favorite.Anime)
	}

	for _, s := range favorite.Scenes {
		fmt.Printf("%s (%s) - %s\n", s.Name, s.Episode, s.Time)
	}
}