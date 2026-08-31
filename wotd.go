package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func handleWOTD() {
	wotd := getWOTD()
	fPrintln("Today's Word of The Day is...")
	time.Sleep(time.Second)
	fPrintlnf(toASCIIArtText(wotd))
}

func getWOTD() string {
	date := time.Now().Format("2006-01-02")
	dateInt, _ := strconv.Atoi(strings.ReplaceAll(date, "-", ""))
	wordList := getWordList()
	rng := RNG{seed: dateInt}
	return wordList[rng.randInt(0, len(wordList)-1)]
}

func getWordList() []string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	txtPath := filepath.Join(filepath.Dir(exePath), "words_alpha.txt")

	file, err := os.Open(txtPath)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	words := make([]string, 0, 300000)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return words
}
