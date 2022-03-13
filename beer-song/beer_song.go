// package beer implements routines for the beer song
package beer

import "fmt"

// Song returns the entire beer song
func Song() string {
	verses, _ := Verses(99, 0)
	return verses
}

// Verses returns part of the song from start to stop
func Verses(start, stop int) (string, error) {
	if start < stop {
		return "", fmt.Errorf("invalid stop")
	}
	verses := ""
	for i := start; i >= stop; i-- {
		if verse, err := Verse(i); err == nil {
			verses += verse + "\n"
		} else {
			return "", err
		}
	}
	return verses, nil
}

// Verse returns the nth verse
func Verse(n int) (string, error) {
	verse := ""
	if n > 99 || n < 0 {
		return verse, fmt.Errorf("invalid number of bottles")
	}
	switch {
	case n > 1:
		verse += fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\n", n, n)
		if n == 2 {
			verse += "Take one down and pass it around, 1 bottle of beer on the wall.\n"
		} else {
			verse += fmt.Sprintf("Take one down and pass it around, %d bottles of beer on the wall.\n", n-1)
		}
	case n == 1:
		verse += "1 bottle of beer on the wall, 1 bottle of beer.\n"
		verse += "Take it down and pass it around, no more bottles of beer on the wall.\n"
	case n == 0:
		verse += "No more bottles of beer on the wall, no more bottles of beer.\n"
		verse += "Go to the store and buy some more, 99 bottles of beer on the wall.\n"
	}
	return verse, nil
}
