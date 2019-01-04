package main

import (
	"errors"
)

// Dictionary is a custom type with base type a map of string
type Dictionary map[string]string

// ErrNotFound returned when word is not found in the dictionary
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search accepts the word to be searched, and definition of the word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add accepts word and definition, and adds them to dictionary
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}

func main() {}
