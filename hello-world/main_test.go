package main_test

import (
	"testing"

	hw "github.com/chilledornaments/learn-go-with-tests/hello-world"
)

func TestGreet(t *testing.T) {
	t.Parallel()

	type args struct {
		name string
		lang string
	}

	type cases struct {
		name string
		args args
		want string
	}

	testCases := []cases{
		{
			name: "Default hello world",
			args: args{
				name: "world",
				lang: "",
			},
			want: "Hello, world",
		},
		{
			name: "Empty string",
			args: args{
				name: "",
				lang: "",
			},
			want: "Hello, world",
		},
		{
			name: "Spanish with name",
			args: args{
				name: "Test",
				lang: "sp",
			},
			want: "Hola, Test",
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				v := hw.Greet(tc.args.name, tc.args.lang)

				if v != tc.want {
					t.Errorf("got %q want %q", v, tc.want)
				}
			},
		)
	}
}
