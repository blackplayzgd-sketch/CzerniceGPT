package main

import (
	"bufio"
	"fmt"
	"iter"
	"maps"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	A = " AAA \n" +
		"A   A\n" +
		"AAAAA\n" +
		"A   A\n" +
		"A   A\n"

	B = "BBBB \n" +
		"B   B\n" +
		"BBBB \n" +
		"B   B\n" +
		"BBBB \n"

	C = " CCCC\n" +
		"C    \n" +
		"C    \n" +
		"C    \n" +
		" CCCC\n"

	D = "DDDD \n" +
		"D   D\n" +
		"D   D\n" +
		"D   D\n" +
		"DDDD \n"

	E = "EEEEE\n" +
		"E    \n" +
		"EEEE \n" +
		"E    \n" +
		"EEEEE\n"

	F = "FFFFF\n" +
		"F    \n" +
		"FFF  \n" +
		"F    \n" +
		"F    \n"

	G = " GGG \n" +
		"G    \n" +
		"G  GG\n" +
		"G   G\n" +
		" GGG \n"

	H = "H   H\n" +
		"H   H\n" +
		"HHHHH\n" +
		"H   H\n" +
		"H   H\n"

	I = "IIIII\n" +
		"  I  \n" +
		"  I  \n" +
		"  I  \n" +
		"IIIII\n"

	J = "    J\n" +
		"    J\n" +
		"    J\n" +
		"J   J\n" +
		" JJJ \n"

	K = "K  K \n" +
		"K K  \n" +
		"KK   \n" +
		"K K  \n" +
		"K  K \n"

	L = "L    \n" +
		"L    \n" +
		"L    \n" +
		"L    \n" +
		"LLLLL\n"

	M = "MM MM\n" +
		"M M M\n" +
		"M   M\n" +
		"M   M\n" +
		"M   M\n"

	N = "NN  N\n" +
		"N N N\n" +
		"N N N\n" +
		"N N N\n" +
		"N  NN\n"

	O = " OOO \n" +
		"O   O\n" +
		"O   O\n" +
		"O   O\n" +
		" OOO \n"

	P = "PPPP \n" +
		"P   P\n" +
		"PPPP \n" +
		"P    \n" +
		"P    \n"

	Q = " QQQ \n" +
		"Q   Q\n" +
		"Q   Q\n" +
		"Q   Q\n" +
		" QQQQ\n"

	R = "RRRR \n" +
		"R   R\n" +
		"RRRR \n" +
		"R   R\n" +
		"R   R\n"

	S = " SSSS\n" +
		"S    \n" +
		" SSS \n" +
		"    S\n" +
		"SSSS \n"

	T = "TTTTT\n" +
		"  T  \n" +
		"  T  \n" +
		"  T  \n" +
		"  T  \n"

	U = "U   U\n" +
		"U   U\n" +
		"U   U\n" +
		"U   U\n" +
		" UUU \n"

	V = "V   V\n" +
		"V   V\n" +
		" V V \n" +
		" V V \n" +
		"  V  \n"

	W = "W   W\n" +
		"W   W\n" +
		"W   W\n" +
		"W W W\n" +
		" W W \n"

	X = "X   X\n" +
		" X X \n" +
		"  X  \n" +
		" X X \n" +
		"X   X\n"

	Y = "Y   Y\n" +
		" Y Y \n" +
		"  Y  \n" +
		"  Y  \n" +
		"  Y  \n"

	Z = "ZZZZ \n" +
		"    Z\n" +
		" ZZZ \n" +
		"Z    \n" +
		" ZZZZ\n"

	SPACE = "   \n" +
		"   \n" +
		"   \n" +
		"   \n" +
		"   \n"

	COLON = ":::\n" +
		": :\n" +
		":::\n" +
		": :\n" +
		":::\n"

	EXCLAMATION = "!!!\n" +
		"!!!\n" +
		"!!!\n" +
		"   \n" +
		"!!!\n"

	DOT = "  \n" +
		"  \n" +
		"  \n" +
		"..\n" +
		"..\n"
)

var (
	ALPHABET = []string{A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P, Q, R, S, T, U, V, W, X, Y, Z, SPACE, COLON, EXCLAMATION, DOT}

	letterMap = map[string]string{
		"a": A,
		"b": B,
		"c": C,
		"d": D,
		"e": E,
		"f": F,
		"g": G,
		"h": H,
		"i": I,
		"j": J,
		"k": K,
		"l": L,
		"m": M,
		"n": N,
		"o": O,
		"p": P,
		"q": Q,
		"r": R,
		"s": S,
		"t": T,
		"u": U,
		"v": V,
		"w": W,
		"x": X,
		"y": Y,
		"z": Z,
		" ": SPACE,
		":": COLON,
		"!": EXCLAMATION,
		".": DOT,
	}
)

type RNG struct {
	seed int
}

func toASCIIArtText(text string) string {
	letters := strings.Split(strings.ToLower(text), "")

	var toReturn string

	for i := 0; i < 5; i++ {
		for idx := range letters {
			toReturn += getRowOfLetter(i+1, getASCIIArt(letters[idx]))
		}

		toReturn += "\n"
	}

	return toReturn
}

func getRowOfLetter(row int, letter string) string {
	if !slices.Contains(ALPHABET, letter) {

		return ""
	}

	return strings.Split(letter, "\n")[row-1] + " "
}

func getASCIIArt(letter string) string {
	var inMap bool

	for key := range maps.Keys(letterMap) {
		if key == letter {
			inMap = true
		}
	}

	if !inMap {
		if letter == "ę" {
			return E
		}
		if letter == "ł" {
			return L
		}
		if letter == "ą" {
			return A
		}
		if letter == "ć" {
			return C
		}
		if letter == "ń" {
			return N
		}
		if letter == "ś" {
			return S
		}
		if letter == "ź" || letter == "ż" {
			return Z
		}
		if letter == "ó" {
			return O
		}
		return A
	}

	return letterMap[letter]
}

func newTxtToString(path string) string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	txtPath := filepath.Join(filepath.Dir(exePath), path)

	file, err := os.Open(txtPath)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	r := bufio.NewReader(file)
	var text string

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			text += string(line)
			text += "\n"
		}
		if err != nil {
			break
		}
	}

	return text
}

func (rng RNG) next() int {
	a := 1664525
	c := 1013904223
	m := int(math.Pow(2, 32))

	rng.seed = (a*rng.seed + c) % m
	return rng.seed
}

func (rng RNG) randInt(min int, max int) int {
	raw := rng.next()
	return min + (raw % (max - min + 1))
}

func maxIterValue(seq iter.Seq[int]) int {
	var maximum int

	for num := range seq {
		if num > maximum {
			maximum = num
		}
	}

	return maximum
}

func fToStr(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func between(num float64, lower float64, upper float64) bool {
	return (lower <= num) && (num < upper)
}

func getWindDirectionFromAngle(angle float64) string {
	if between(angle, 0, 22.5) || between(angle, 337.5, 360) {
		return "North"
	} else if between(angle, 22.5, 67.5) {
		return "Northeast"
	} else if between(angle, 67.5, 112.5) {
		return "East"
	} else if between(angle, 112.5, 157.5) {
		return "Southeast"
	} else if between(angle, 157.5, 202.5) {
		return "South"
	} else if between(angle, 202.5, 247.5) {
		return "Southwest"
	} else if between(angle, 247.5, 292.5) {
		return "Wwest"
	} else if between(angle, 292.5, 337.5) {
		return "Northwest"
	} else {
		return "idfk figure it out yourself genius"
	}
}

func fPrint(text string, delay ...time.Duration) {
	def := DefT
	if len(delay) > 0 {
		def = delay[0]
	}
	delRange := int(def * 2)
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(def + time.Duration(rand.Int()%delRange))
	}
}

func fPrintln(text string, delay ...time.Duration) {
	def := DefT
	if len(delay) > 0 {
		def = delay[0]
	}

	fPrint(text, def)
	fmt.Println()
}

func fPrintlnf(text string) {
	fPrintln(text, DefT/3)
}
