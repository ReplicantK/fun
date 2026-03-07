package main

import "testing"

func TestSearch(t *testing.T) {
	t.Run("testing with a valid value inside dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"

		if got != want {
			t.Error("expected:", want, "got:", got)
		}
	})

	t.Run("testing with a word that is not in the dictionary", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("sup")

		if err == nil {
			t.Fatal("expected and error, but got nil")
		}

		if err.Error() != Error_Not_Found_Dictionary.Error() {
			t.Error("expected:", Error_Not_Found_Dictionary, "got:", err.Error()) 
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding to a dictionary", func(t *testing.T) {
		dic := Dictionary{}

		dic.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dic.Search("test")

		if err != nil {
			t.Fatal("there should be this word in the dictionary:", want)
		}

		if want != got {
			t.Error("expecting:", want, "got:", got)
		}
	})

	t.Run("adding when a value already exists", func(t *testing.T) {
		dic := Dictionary{"test": "this is just a test"}

		err := dic.Add("test", "this is just a test")

		if err == nil {
			t.Fatal("there should have been an error returned")
		}

		if err.Error() != "cannot add key value pair, it already exists" {
			t.Error("expecting: cannot add key value pair, it already exists, got: ", err.Error())
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("testing when a key already exists", func(t *testing.T ) {
		key := "test"
		value := "this is just a test"

		dic := Dictionary{key: value}

		new_value := "new definition"

		dic.Update(key, new_value)

		got, err := dic.Search(key)

		if err != nil {
			t.Fatal("there should not be an error returned when updating a key value pair")
		}

		if got != new_value {
			t.Error("expected:", new_value, "got:", got)
		}
	})

	t.Run("testing when a key does not exist", func(t *testing.T) {
		word := "test"
		def := "this is just a test"

		dic := Dictionary{}

		err := dic.Update(word, def)

		expected := "cannot update key, as it does not exist!"

		if err == nil {
			t.Fatal("there should be an error returned when updating a key value pair that does not exist")
		}

		if err.Error() != expected {
			t.Error("expected the following error:", expected, "got", err.Error())
		}
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete an existing word test", func(t *testing.T) {
		dic := Dictionary{"test": "this is a test"}

		err := dic.Delete("test")

		if err != nil {
			t.Fatal("there should not have been an error when deleting a key")
		}

		_, err2 := dic.Search("test")
		
		if err2 == nil {
			t.Fatal("there should have been an error when searching for a key that was deleted")
		}

		if err2.Error() != Error_Not_Found_Dictionary.Error() {
			t.Error("expected error:", Error_Not_Found_Dictionary, "got:", err2.Error())
		}
	})

	t.Run("delete a nonexistent word", func(t *testing.T) {
		dic := Dictionary{}

		err := dic.Delete("test")

		if err == nil {
			t.Fatal("there should have been an error when deleting a key because it does not exist")
		}
	})
}
