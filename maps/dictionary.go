package main

// Search accepts a map of string and key to be searched, and returns value of particular key in the map
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}

func main() {}
