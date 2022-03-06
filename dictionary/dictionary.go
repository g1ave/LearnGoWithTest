package dictionary

type Dictionary map[string]string

const (
	NotFoundError         = DictionaryErr("counld't find the word you are looking for")
	WordExistedError      = DictionaryErr("cannot add word because it already exists")
	WordDoesNotExistError = DictionaryErr("can't update the word because it doesnot exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

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
	case nil:
		return WordExistedError
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case NotFoundError:
		return WordDoesNotExistError
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}
