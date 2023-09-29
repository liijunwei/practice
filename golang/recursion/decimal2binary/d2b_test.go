package decimal2binary

import (
	"testing"
)

func Test_d2b(t *testing.T) {
	type test struct {
		input int
		want  string
	}

	tests := []test{
		{
			input: 0,
			want:  "0",
		},
		{
			input: 1,
			want:  "1",
		},
		{
			input: 2,
			want:  "10",
		},
		{
			input: 3,
			want:  "11",
		},
		{
			input: 100,
			want:  "1100100",
		},
		{
			input: 99999,
			want:  "11000011010011111",
		},
	}

	for _, tt := range tests {
		tt := tt

		if got := d2b(tt.input, ""); got != tt.want {
			t.Fatalf("input: %v, want: %v, got: %v", tt.input, tt.want, got)
		}
	}
}
