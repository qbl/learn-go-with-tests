package main

import (
	"errors"
)

// Dictionary is a custom type with base type a map of string
type Dictionary map[string]string

// Search accepts a map of string and key to be searched, and returns value of particular key in the map
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}

func main() {}
