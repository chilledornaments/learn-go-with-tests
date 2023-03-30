package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	type args struct {
		value string
		times int
	}
	testCases := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name:     "Test single character",
			args:     args{value: "a", times: 5},
			expected: "aaaaa",
		},
		{
			name:     "Test multiple characters",
			args:     args{value: "AA", times: 5},
			expected: "AAAAAAAAAA",
		},
		{
			name:     "Test single character 50 times",
			args:     args{value: "a", times: 50},
			expected: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			repeated := Repeat(tc.args.value, tc.args.times)
			if repeated != tc.expected {
				t.Errorf("got %q but expected %q", repeated, tc.expected)

			}
		})
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	v := "Y"
	n := 10
	r := Repeat(v, n)
	fmt.Println(r)
	// Output: YYYYYYYYYY

}
