package aas

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Parallel()
	
	type args struct {
		value []int
	}

	testCases := []struct {
		name     string
		args     args
		expected int
	}{
		{
			name:     "Test one",
			args:     args{value: []int{1, 2, 3}},
			expected: 6,
		},
		{
			name:     "Test two",
			args:     args{value: []int{1, 2, 3, 11, 15, 104}},
			expected: 136,
		},
		{
			name:     "Negative",
			args:     args{value: []int{1, -2}},
			expected: -1,
		},
		{
			name:     "Empty",
			args:     args{value: []int{}},
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Sum(tc.args.value)

			if got != tc.expected {
				t.Errorf("got '%d' but expected '%d'", got, tc.expected)
			}
		})
	}
}

func ExampleSum() {
	a := []int{1, 2, 3, 4, 5}
	s := Sum(a)
	fmt.Println(s)
	// Output: 15
}

func TestSumAll(t *testing.T) {
	t.Parallel()

	type args struct {
		value [][]int
	}
	testCases := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name:     "Simple",
			args:     args{value: [][]int{{1, 2}, {3, 4}}},
			expected: []int{3, 7},
		},
		{
			name:     "Negative",
			args:     args{value: [][]int{{1, -2}, {3, 44}}},
			expected: []int{-1, 47},
		},
		{
			name:     "Empty",
			args:     args{value: [][]int{{}}},
			expected: []int{0},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := SumAll(tc.args.value...)

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got '%d' but expected '%d'", got, tc.expected)
			}

		})
	}
}

func TestSumAllTails(t *testing.T) {
	t.Parallel()

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	type args struct {
		value [][]int
	}

	testCases := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name:     "simple",
			args:     args{value: [][]int{{1, 2, 3}, {3, 4, 5}}},
			expected: []int{5, 9},
		},
		{
			name:     "negative",
			args:     args{value: [][]int{{1, 2, 3}, {1, 1, -4}}},
			expected: []int{5, -3},
		},
		{
			name:     "empty",
			args:     args{value: [][]int{{}}},
			expected: []int{0},
		},
	}

	for _, tc := range testCases {
		t.Run(
			tc.name,
			func(t *testing.T) {
				got := SumAllTails(tc.args.value...)
				checkSums(t, got, tc.expected)
			})
	}
}
