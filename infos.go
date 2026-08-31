package main

import (
	"fmt"
	"maps"
	"path/filepath"
	"strings"
)

var (
	infoConfig      = loadInfoConfig()
	availableTopics []string
)

func handleInfo(prompt string) {
	topic := matchClosestString(prompt, availableTopics)
	bigText := infoConfig[topic]["wText"]
	text := infoConfig[topic]["text"]

	var toReturn string

	if strings.Contains(text, ".txt") {
		toReturn = newTxtToString(filepath.Join("longThingies", text))
	} else {
		toReturn = text
	}

	fmt.Println(toASCIIArtText(bigText))
	fPrintlnf("\n" + toReturn)
}

func loadInfo() {
	getAvailableTopics()
}

func getAvailableTopics() {
	for topic := range maps.Keys(infoConfig) {
		availableTopics = append(availableTopics, topic)
	}
}
