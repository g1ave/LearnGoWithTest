package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("couldn't find the word you are looking for")
	}
	return definition, nil

}
