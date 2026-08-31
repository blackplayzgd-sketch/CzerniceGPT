package main

import (
	"fmt"
)

func handleHelp(command []string) {
	if len(command) == 0 {
		fPrintln("Available prompts:")

		for _, promptType := range PromptTypes {
			fPrintln(string(promptType), DefT/2)
		}

		fPrintln("\nHELP usage:")
		fPrintln("CzerniceGPT help <PROMPT>")

	} else {
		switch getPromptType(command[0]) {
		case Info:
			handleInfoHelp()
		case Joke:
			handleJokeHelp()
		case WOTD:
			handleWOTDHelp()
		case Stats:
			handleStatsHelp()
		case Music:
			handleMusicHelp()
		case Web:
			handleWebHelp()
		case Weather:
			handleWeatherHelp()
		case "all":
			handleInfoHelp()
			handleJokeHelp()
			handleWOTDHelp()
			handleStatsHelp()
			handleMusicHelp()
			handleWebHelp()
			handleWeatherHelp()
		}
	}
}

func handleInfoHelp() {
	fmt.Println(toASCIIArtText("INFO usage:"))
	fPrintln("CzerniceGPT info <TOPIC>")

	fPrintln("\nDescription:")
	fPrintln("Displays information on a chosen topic.")

	fPrintln("\nAvailable topics:")
	for i, topic := range availableTopics {
		fPrint(topic)
		if i != (len(availableTopics) - 1) {
			fPrint(", ")
		}
	}
}

func handleJokeHelp() {
	fmt.Println(toASCIIArtText("JOKE usage:"))
	fPrintln("CzerniceGPT joke")

	fPrintln("\nDescription:")
	fPrintln("Tells a random joke!")
}

func handleWOTDHelp() {
	fmt.Println(toASCIIArtText("WOTD usage:"))
	fPrintln("CzerniceGPT wotd")

	fPrintln("\nDescription:")
	fPrintln("Displays a random word, a different one everyday.")
}

func handleStatsHelp() {
	fmt.Println(toASCIIArtText("STATS usage:"))
	fPrintlnf("CzerniceGPT stats <COMMAND> <ARGS>")

	fPrintlnf("\nA Statistics calculator module for CzerniceGPT!")
	fPrintlnf("Operates on lists of 64-bit floating point numbers.")

	fPrintlnf("\nAvailable commands:")

	fPrintlnf("\nstats create <LIST_NAME>")
	fPrintlnf("Creates a list with the given name (can't contain spaces).")

	fPrintlnf("\nstats add <LIST_NAME> <ARGS>")
	fPrintlnf("Adds the given elements to the target list.")
	fPrintlnf("Supports many elements at once, separated by spaces (eg. add example 1 2.3 -4 5)")

	fPrintlnf("\nstats remove <LIST_NAME> <INDEX>")
	fPrintlnf("Removes the element at the given index (starting at 0) from the target list.")
	fPrintlnf("Removing all elements does NOT actually delete the list, use stats delete instead.")

	fPrintlnf("\nstats display <LIST_NAME>")
	fPrintlnf("Displays all elements of the target list.")

	fPrintlnf("\nstats delete <LIST_NAME>")
	fPrintlnf("Deletes the target list.")

	fPrintlnf("\nstats all_stats <LIST_NAME>")
	fPrintlnf("Calculates and displays various statistics of the target list.")
}

func handleMusicHelp() {
	fmt.Println(toASCIIArtText("MUSIC usage:"))
	fPrintlnf("CzerniceGPT music <COMMAND>")

	fPrintlnf("\nDescription:")
	fPrintlnf("An Audio module for CzerniceGPT! (not very robust because I hate working with audio)")

	fPrintlnf("\nAvailable commands:")

	fPrintlnf("\nmusic play")
	fPrintlnf("Plays all audio files from the songs directory (must be MP3) in randomized order.")
}

func handleWebHelp() {
	fmt.Println(toASCIIArtText("WEB usage:"))
	fPrintlnf("CzerniceGPT web <COMMAND>")

	fPrintlnf("\nDescription:")
	fPrintlnf("A miscellaneous module for web-related stuff for CzerniceGPT!")

	fPrintlnf("\nAvailable commands:")

	fPrintlnf("\nweb get <URL>")
	fPrintlnf("Issues a GET to the specified URL and prints the response.")
}

func handleWeatherHelp() {
	fmt.Println(toASCIIArtText("WEATHER usage:"))
	fPrintlnf("CzerniceGPT weather <LOCATION>")

	fPrintlnf("\nDescription:")
	fPrintlnf("Displays the current weather and 7-day forecast for the provided location.")
	fPrintlnf("Location name can't contain spaces.")
}
