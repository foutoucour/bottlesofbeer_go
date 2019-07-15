package main

import (
	"testing"
)

func Test99bottles(t *testing.T) {
	expected := `99 bottles of beer on the wall, 99 bottles of beer.
Take one down and pass it around, 98 bottles of beer on the wall.`
	number := 99
	result := lineLyrics(number)
	if expected != result {
		errorMessage(t, number, expected, result)
	}
}

func Test2bottles(t *testing.T) {
	expected := `2 bottles of beer on the wall, 2 bottles of beer.
Take one down and pass it around, 1 bottle of beer on the wall.`
	number := 2
	result := lineLyrics(number)
	if expected != result {
		errorMessage(t, number, expected, result)
	}
}

func Test1bottle(t *testing.T) {
	expected := `1 bottle of beer on the wall, 1 bottle of beer.
Take one down and pass it around, no more bottles of beer on the wall.`
	number := 1
	result := lineLyrics(number)
	if expected != result {
		errorMessage(t, number, expected, result)
	}
}

func TestNoMoreBottles(t *testing.T) {
	expected := `No more bottles of beer on the wall, no more bottles of beer.
Go to the store and buy some more, 99 bottles of beer on the wall.`
	number := 0
	result := lineLyrics(number)
	if expected != result {
		errorMessage(t, number, expected, result)
	}

}

func errorMessage(t *testing.T, number int, expected string, result string) {
	t.Errorf("not the same for number: %d.\nexpexted:\n%s\nresult:\n%s", number, expected, result)
}