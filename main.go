package main

import (
	"os"
	"path/filepath"
	"time"
)

type PromptType string

const (
	Info    PromptType = "info"
	Joke    PromptType = "joke"
	WOTD    PromptType = "wotd"
	Stats   PromptType = "stats"
	Music   PromptType = "music"
	Web     PromptType = "web"
	Weather PromptType = "weather"
	Help    PromptType = "help"

	DefT time.Duration = 20 * time.Millisecond
)

var (
	PromptTypes []PromptType
)

func main() {
	prompt := os.Args[1:]
	getOutput(prompt)
}

func init() {
	loadInfo()
	loadStats()
	loadMusic()
	loadWeb()
	loadWeather()
	loadMain()
}

func getPromptType(prompt string) PromptType {
	return PromptType(prompt)
}

func getOutput(prompt []string) {
	promptType := getPromptType(prompt[0])

	switch promptType {
	case Info:
		handleInfo(prompt[1])
	case Joke:
		handleJoke()
	case WOTD:
		handleWOTD()
	case Stats:
		handleStats(prompt[1:])
	case Music:
		handleMusic(prompt[1:])
	case Web:
		handleWeb(prompt[1:])
	case Weather:
		handleWeather(prompt[1])
	case Help:
		handleHelp(prompt[1:])
	default:
		fPrintln("Prompt type invalid", DefT)
	}

}

func getNicePath(filename string) string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(filepath.Dir(exePath), filename)

	return path
}

func loadMain() {
	PromptTypes = append(PromptTypes, Info, Joke, WOTD, Stats, Music, Web, Weather, Help)
}
