package maps

import (
	"errors"
)

type Dictionary map[string]string

var ErrWordNotInDictionary = errors.New("word not found")
var ErrWordAlreadyExists = errors.New("word already in dictionary")

func (d Dictionary) Search(w string) (string, error) {
	v, ok := d[w]

	if !ok {
		return "", ErrWordNotInDictionary
	}

	return v, nil
}

func (d Dictionary) Add(k, v string) error {
	// If there's no error, the word is already in the dict
	if _, err := d.Search(k); err == nil {
		return ErrWordAlreadyExists
	}

	d[k] = v

	return nil
}

func (d Dictionary) Update(k, v string) error {
	if _, err := d.Search(k); err != nil {
		return ErrWordNotInDictionary
	}

	d[k] = v
	return nil
}

func (d Dictionary) Delete(k string) {
	delete(d, k)
}
