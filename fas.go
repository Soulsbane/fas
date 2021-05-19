package main

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
