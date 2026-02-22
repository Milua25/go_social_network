package maps

import (
	"errors"
)

var (
	ErrNotFound      = errors.New("could not find the word you were looking for")
	ErrAlreadyExists = errors.New("the word already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(val string) (string, error) {
	if _, ok := d[val]; !ok {
		return "", ErrNotFound
	}
	return d[val], nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(value)
	switch {
	case errors.Is(err, ErrNotFound):
		d[key] = value
	case err == nil:
		return ErrAlreadyExists
	default:
		return err
	}
	return nil
}
