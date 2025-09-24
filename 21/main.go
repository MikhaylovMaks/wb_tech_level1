package main

import "fmt"

type Music interface {
	Play()
}

type MP3 struct{}

func (m *MP3) Play() {
	fmt.Println("Playing MP3")
}

type Cassette struct{}

func (c *Cassette) PlayCassette() {
	fmt.Println("Playing cassette")
}

type CassetteAdapter struct {
	cassette *Cassette
}

func (a *CassetteAdapter) Play() {
	fmt.Println("Addapter take cassette...")
	a.cassette.PlayCassette()
}

func main() {
	var player Music

	// Используем MP3 напрямую
	player = &MP3{}
	player.Play()

	// Используем кассету через адаптер
	cassette := &Cassette{}
	player = &CassetteAdapter{cassette: cassette}
	player.Play()
}
