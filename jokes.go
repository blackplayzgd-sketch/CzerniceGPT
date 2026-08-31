package main

import (
	"math/rand"
	"time"
)

var (
	jokeConfig = loadJokeConfig()
)

func handleJoke() {
	jokeMap := jokeConfig[rand.Intn(len(jokeConfig))]
	var toReturn string

	switch jokeMap["type"] {
	case "real":
		toReturn = jokeMap["senchline"]
		fPrintln(toReturn)

	case "complex":
		toReturn = jokeMap["setup"] + jokeMap["punchline"]
		fPrintln(jokeMap["setup"])
		time.Sleep(time.Second)
		fPrintln(jokeMap["punchline"], 50*time.Millisecond)
	default:
		toReturn = "yeah no joke like that exists lol"
		fPrintln(toReturn)
	}
}
