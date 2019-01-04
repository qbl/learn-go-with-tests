package main

import (
	"errors"
)

// Dictionary is a custom type with base type a map of string
type Dictionary map[string]string

// ErrNotFound returned when word is not found in the dictionary
var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

// Search accepts the word to be searched, and definition of the word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add accepts word and definition, and adds them to dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func main() {}
