package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrorNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("word already exists in dictionary")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]

	if ok {
		return ErrWordExists
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) {
	// return ErrorNotFound
	_, err := d.Search(word)

	if err !
}
