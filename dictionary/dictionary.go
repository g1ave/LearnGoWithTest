package dictionary

import "errors"

type Dictionary map[string]string

var (
	NotFoundError    = errors.New("counld't find the word you are looking for")
	WordExistedError = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", NotFoundError
	}
	return definition, nil

}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case NotFoundError:
		d[word] = definition
		return nil
	case nil:
		return WordExistedError
	default:
		return err
	}
}
