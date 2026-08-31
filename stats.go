package main

import (
	"slices"
	"strconv"
)

type UtilCommand string
type StatCommand string

const (
	Create   UtilCommand = "create"
	Add      UtilCommand = "add"
	Display  UtilCommand = "display"
	Remove   UtilCommand = "remove"
	Delete   UtilCommand = "delete"
	AllStats UtilCommand = "all_stats"

	ArithMean StatCommand = "ari_mean"
	GeoMean   StatCommand = "geo_mean"
	Median    StatCommand = "median"
	Mode      StatCommand = "mode"
	Min       StatCommand = "min"
	Max       StatCommand = "max"
	IQR       StatCommand = "iqr"
	Variance  StatCommand = "variance"
	StdDev    StatCommand = "std_dev"
	Skewness  StatCommand = "skewness"
)

var (
	UtilCommands []UtilCommand
	StatCommands []StatCommand
)

func handleStats(command []string) {
	inputType := command[0]

	if slices.Contains(UtilCommands, UtilCommand(inputType)) {
		handleUtilCommand(command)
	} else if slices.Contains(StatCommands, StatCommand(command[0])) {
		handleStatCommand(command)
	} else {
		fPrintln("i wil murder yu")
	}

}

func handleUtilCommand(command []string) {
	switch UtilCommand(command[0]) {
	case Create:
		handleCreate(command[1])
	case Add:
		handleAdd(command[1], command[2:])
	case Display:
		handleDisplay(command[1])
	case Remove:
		handleRemove(command[1], command[2])
	case Delete:
		handleDelete(command[1])
	case AllStats:
		handleAllStats(command[1])
	}
}

func handleStatCommand(command []string) {
	switch StatCommand(command[0]) {
	case ArithMean:
		handleArithMean(command[1])
	case GeoMean:
		handleGeoMean(command[1])
	case Median:
		handleMedian(command[1])
	case Mode:
		handleMode(command[1])
	case Min:
		handleMin(command[1])
	case Max:
		handleMax(command[1])
	case IQR:
		handleIQR(command[1])
	case Variance:
		handleVariance(command[1])
	case StdDev:
		handleStdDev(command[1])
	case Skewness:
		handleSkewness(command[1])
	}
}

func handleCreate(listName string) {
	lists := loadStatsData()
	_, is := lists[listName]
	if is {
		fPrintln("A list with this name already exists! Either choose a different name or expunge that bih")
	} else {
		lists[listName] = []float64{}
	}

	saveStatsData(lists)
}

func handleAdd(listName string, numStrings []string) {
	lists := loadStatsData()

	list, is := lists[listName]

	var nums []float64

	for _, str := range numStrings {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
	}

	if !is {
		fPrintln("A list with this name does NOT already existing. Nope. Do not even THINK about adding ANYTHING to this list. Either choose a different one or make one with this name idioute")
	} else {
		lists[listName] = append(list, nums...)
	}

	saveStatsData(lists)
}

func handleDisplay(listName string) {
	lists := loadStatsData()

	list, is := lists[listName]
	if !is {
		fPrintln("A list with this name does NOT already existing. Nope. Do not even THINK about displaying ANYTHING from this list. Either choose a different one or make one with this name imbcil")
	} else if len(list) == 0 {
		fPrintln("At this very particular moment, the list colloquially referred to as " + listName + " contains the following (presumably real) numbers:")
		fPrintln("None!!!1!!1!! Go add sum numbers you SG (Silly gouse)!")
	} else {
		fPrintln(listName + "has:")
		for i, num := range list {
			fPrintln(strconv.FormatFloat(num, 'f', -1, 64))
			if i < (len(list) - 1) {
				fPrintln(", ")
			}
		}
	}
}

func handleRemove(listName string, idxString string) {
	lists := loadStatsData()

	list, is := lists[listName]
	idx, err := strconv.Atoi(idxString)
	if err != nil {
		panic(err)
	}

	if !is {
		fPrintln("A list with this name does NOT already existing. Nope. You CREUL bastard!!! You tried to hurt a list that doesn't even is!! Either choose a different one or make one with this name murederer")
	} else if idx > len(list)-1 {
		fPrintln("Oh you CREUL CREUL bastard!!! You thought you can HURT this list that LOW but you CANET!!!!!1!1!1!1 If you want be CREUL chuse a smaller indecks!!!")
	} else {
		lists[listName] = append(list[:idx], list[idx+1:]...)
	}

	saveStatsData(lists)
}

func handleDelete(listName string) {
	lists := loadStatsData()

	_, is := lists[listName]
	if !is {
		fPrintln("A list with this name does NOT already existing. Nope. You MUREDERUR You tried to KILL a list that doesn't even is!! Either choose a different one or make one with this name could harted keeler!!!")
	} else {
		delete(lists, listName)
	}

	saveStatsData(lists)
}

func loadStats() {
	UtilCommands = []UtilCommand{Create, Add, Display, Remove, Delete, AllStats}
	StatCommands = []StatCommand{ArithMean, GeoMean, Median, Mode, Min, Max, IQR, Variance, StdDev, Skewness}
}

func handleArithMean(listName string) {
	lists := loadStatsData()
	list := lists[listName]

	fPrintln(getOutputMessageOfStat("Arithmetic Mean", calcArithMean(list), listName))
}

func handleGeoMean(listName string) {
	lists := loadStatsData()
	list := lists[listName]

	fPrintln(getOutputMessageOfStat("Geometric Mean", calcGeoMean(list), listName))
}

func handleMedian(listName string) {
	lists := loadStatsData()
	list := lists[listName]

	fPrintln(getOutputMessageOfStat("Median", calcMedian(list), listName))
}

func handleMode(listName string) {
	lists := loadStatsData()
	list := lists[listName]
	mode := calcMode(list)

	if mode == nil {
		fPrint("List " + listName + " has no mode!\n")
	} else {
		fPrintln("The Mode(s) of list " + listName + " is/are " + getNiceStringYk(calcMode(list)))
	}
}

func handleMin(listName string) {
	lists := loadStatsData()
	list := lists[listName]

	fPrintln(getOutputMessageOfStat("Minimum", slices.Min(list), listName))
}

func handleMax(listName string) {
	lists := loadStatsData()
	list := lists[listName]

	fPrintln(getOutputMessageOfStat("Maximum", slices.Max(list), listName))
}

func handleIQR(listName string) {
	lists := loadStatsData()
	list := lists[listName]

	fPrintln(getOutputMessageOfStat("IQR (Interquartile Range)", calcIQR(list), listName))
}

func handleVariance(listName string) {
	lists := loadStatsData()
	list := lists[listName]

	fPrintln(getOutputMessageOfStat("Variance", calcVariance(list), listName))
}

func handleStdDev(listName string) {
	lists := loadStatsData()
	list := lists[listName]

	fPrintln(getOutputMessageOfStat("Standard Deviation", calcStdDev(list), listName))
}

func handleSkewness(listName string) {
	lists := loadStatsData()
	list := lists[listName]

	fPrintln(getOutputMessageOfStat("Skewness", calcSkewness(list), listName))
}

func handleAllStats(listName string) {
	handleDisplay(listName)
	handleArithMean(listName)
	handleGeoMean(listName)
	handleMedian(listName)
	handleMode(listName)
	handleMin(listName)
	handleMax(listName)
	handleIQR(listName)
	handleVariance(listName)
	handleStdDev(listName)
	handleSkewness(listName)
}

func getOutputMessageOfStat(statName string, stat float64, listName string) string {
	return "The " + statName + " of " + listName + " is " + strconv.FormatFloat(stat, 'f', -1, 64)
}

func getNiceStringYk(nums []float64) string {
	var output string

	for _, num := range nums {
		output += strconv.FormatFloat(num, 'f', -1, 64)
	}

	return output
}
