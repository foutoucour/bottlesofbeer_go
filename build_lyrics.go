package main

import (
	"fmt"
	"strings"
)
//
//func main() {
//	fmt.Printf(lyrics())
//}

// Build a single line of the song
// Take in count the special cases of each numbers of bottle
func lineLyrics(number int) string {
	var lyrics []string
	switch number {
	default:
		lyrics = append(
			lyrics, fmt.Sprintf("%[1]d bottles of beer on the wall, %[1]d bottles of beer.", number))
	case 1:
		lyrics = append(lyrics, "1 bottle of beer on the wall, 1 bottle of beer.")
	case 0:
		lyrics = append(lyrics, "No more bottles of beer on the wall, no more bottles of beer.")
	}

	switch number {
	default:
		lyrics = append(
			lyrics, fmt.Sprintf("Take one down and pass it around, %d bottles of beer on the wall.", number-1))
	case 2:
		lyrics = append(lyrics, "Take one down and pass it around, 1 bottle of beer on the wall.")
	case 1:
		lyrics = append(lyrics, "Take one down and pass it around, no more bottles of beer on the wall.")
	case 0:
		lyrics = append(lyrics, "Go to the store and buy some more, 99 bottles of beer on the wall.")
	}

	return strings.Join(lyrics, "\n")
}

// Build the set of lyrics of the song
func lyrics() string {
	var lyrics []string
	for number := 99; number >= 0; number-- {
		lyrics = append(lyrics, lineLyrics(number))
		// No additional line for the last line of the text
		if number != 0 {
			lyrics = append(lyrics, "")
		}
	}
	return strings.Join(lyrics, string('\n'))
}
