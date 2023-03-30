package concurrency_two

import "testing"

func Test_generator(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator()
		})
	}
}

func BenchmarkGenerator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generator()
	}
}
