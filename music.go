package main

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/gopxl/beep/v2/speaker"
)

type MusicCommand string

const (
	Play MusicCommand = "play"
)

var (
	playlist     []string
	PlaylistPath string = getNicePath("songs")
)

func handleMusic(command []string) {
	switch MusicCommand(command[0]) {
	case Play:
		playPlaylist()
	}
}

func playPlaylist() {
	err := speaker.Init(sr, sr.N(time.Second/2))
	if err != nil {
		panic(err)
	}

	for range len(playlist) {
		rng := RNG{seed: rand.Int()}
		toPlayIdx := rng.randInt(0, len(playlist)-1)

		songDuration := playSongAndWaitForItToFinish(playlist[toPlayIdx])
		time.Sleep(songDuration)
	}
}

func loadMusic() {
	files, err := os.ReadDir(PlaylistPath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		playlist = append(playlist, filepath.Join(PlaylistPath, file.Name()))
	}
}
