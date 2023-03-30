package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Parallel()

	t.Run(
		"existing key",
		func(t *testing.T) {
			d := Dictionary{"test": "this is a test"}

			got, err := d.Search("test")

			want := "this is a test"

			assertNoError(t, err)

			assertDefinition(t, got, want)
		},
	)

	t.Run("missing key", func(t *testing.T) {
		d := Dictionary{"test": "this is a test"}

		got, err := d.Search("NONE")

		want := ""

		assertError(t, err, ErrWordNotInDictionary)

		assertDefinition(t, got, want)
	})

}

func TestAdd(t *testing.T) {
	t.Run(
		"add new word",
		func(t *testing.T) {
			word := "hello"
			definition := "world"

			d := Dictionary{}
			err := d.Add(word, definition)

			assertNoError(t, err)

			def, err := d.Search(word)

			assertNoError(t, err)

			assertDefinition(t, def, definition)

		})
}

func TestUpdate(t *testing.T) {
	t.Parallel()

	t.Run("update existing word", func(t *testing.T) {
		word := "foo"
		d := Dictionary{word: "bar"}
		want := "test"
		err := d.Update(word, want)
		assertNoError(t, err)

		v, err := d.Search(word)

		assertNoError(t, err)
		assertDefinition(t, v, want)

	})

	t.Run(
		"update new word",
		func(t *testing.T) {
			word := "foo"
			d := Dictionary{}
			want := ErrWordNotInDictionary
			err := d.Update(word, "bar")
			assertError(t, err, want)

			_, err = d.Search(word)
			assertError(t, err, want)

		})
}

func TestDelete(t *testing.T) {
	t.Parallel()

	t.Run(
		"delete",
		func(t *testing.T) {
			word := "foo"
			def := "bar"
			d := Dictionary{word: def}

			d.Delete(word)

			_, err := d.Search(word)

			assertError(t, err, ErrWordNotInDictionary)
		})
}

var assertNoError = func(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatalf("got error but expected none %q", got.Error())
	}
}

var assertError = func(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected error but got none")
	}

	if got != want {
		t.Errorf("got %q but wanted %q", got.Error(), want.Error())
	}
}

var assertDefinition = func(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}
}
