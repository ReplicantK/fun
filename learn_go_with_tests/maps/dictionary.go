package main

import "errors"

type Dictionary map[string]string

var Error_Not_Found_Dictionary error = errors.New("the provided word is not found in the dictionary")

func (d Dictionary) Search(word string) (string, error) {
	found, err := d[word]

	if !err {
		return "", Error_Not_Found_Dictionary
	}

	return found, nil
}

func Search(dic map[string]string, word string) string {
	return dic[word]
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]

	if ok {
		return errors.New("cannot add key value pair, it already exists")
	}

	d[key] = value

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, ok := d[key]

	if !ok {
		return errors.New("cannot update key, as it does not exist!")
	}

	d[key] = value

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, ok := d[key]	

	if !ok {
		return errors.New("cannot delete key because it does not exist")
	}

	delete(d, key)

	return nil
}
