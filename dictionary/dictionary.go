package dictionary

import "errors"

type Dictionary map[string]string

var NotFoundError = errors.New("counld't find the word you are looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", NotFoundError
	}
	return definition, nil

}
