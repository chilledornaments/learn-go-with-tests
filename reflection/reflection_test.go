package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func Test_walk(t *testing.T) {
	tests := []struct {
		name          string
		input         interface{}
		expectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Mitch"},
			[]string{"Mitch"},
		},
		{
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Mitch", "Boulder"},
			[]string{"Mitch", "Boulder"},
		},
		{
			"struct with string and int fields",
			struct {
				Name string
				Age  int
			}{"Mitch", 100},
			[]string{"Mitch"},
		},
		{
			"struct with nested fields",
			Person{
				Name: "Mitch",
				Profile: Profile{
					Age:  100,
					City: "Boulder",
				},
			},
			[]string{"Mitch", "Boulder"},
		},
		{
			"pointer to struct",
			&Person{
				Name: "Mitch",
				Profile: Profile{
					Age:  100,
					City: "Boulder",
				},
			},
			[]string{"Mitch", "Boulder"},
		},
		{
			"slices",
			[]Profile{
				{99, "Hello"},
				{100, "world"},
			},
			[]string{"Hello", "world"},
		},
		{
			"arrays",
			[2]Profile{
				{99, "Hello"},
				{100, "world"},
			},
			[]string{"Hello", "world"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			f := func(x string) {
				got = append(got, x)
			}
			walk(tt.input, f)

			if !reflect.DeepEqual(got, tt.expectedCalls) {
				t.Errorf("got %v want %v", got, tt.expectedCalls)
			}
		})
	}

	// Maps don't preserve order, so map tests are broken into their own run
	t.Run(
		"with maps",
		func(t *testing.T) {
			m := map[string]string{
				"hello": "world",
				"foo":   "bar",
			}

			var got []string

			walk(m, func(s string) {
				got = append(got, s)
			})

			assertContains(t, got, "world")
			assertContains(t, got, "bar")
		},
	)

	t.Run(
		"channel",
		func(t *testing.T) {
			c := make(chan Profile)

			go func() {
				c <- Profile{100, "Boulder"}
				c <- Profile{99, "Denver"}
				close(c)
			}()

			var got []string
			want := []string{"Boulder", "Denver"}

			walk(c, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)

	t.Run(
		"function",
		func(t *testing.T) {
			f := func() (Profile, Profile) {
				return Profile{100, "Boulder"}, Profile{99, "Denver"}
			}

			var got []string
			want := []string{"Boulder", "Denver"}

			walk(f, func(s string) {
				got = append(got, s)
			})

			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		},
	)
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
