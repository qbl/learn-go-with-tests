package main

// Dictionary is a custom type with base type a map of string
type Dictionary map[string]string

// Search accepts a map of string and key to be searched, and returns value of particular key in the map
func (d Dictionary) Search(word string) string {
	return d[word]
}

func main() {}
