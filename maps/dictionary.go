package main

// Various error types
const (
	ErrNotFound      = DictionaryErr("could not find the word you were looking for")
	ErrWordExists    = DictionaryErr("can not add word because it already exists")
	ErrWordNotExists = DictionaryErr("can not update word because it does not exists")
)

// DictionaryErr is a custom type to describe errors
type DictionaryErr string

// Dictionary is a custom type with base type a map of string
type Dictionary map[string]string

func (e DictionaryErr) Error() string {
	return string(e)
}

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

// Update accepts word and definition, and updates word's definition accordingly
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = definition
	case ErrNotFound:
		return ErrWordNotExists
	default:
		return err
	}

	return nil
}

// Delete accepts word, and deletes word from dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func main() {}
