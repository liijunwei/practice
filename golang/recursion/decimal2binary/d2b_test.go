package decimal2binary

import "testing"

func Test_d2b(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{
			input: "0",
			want:  "0",
		},
		{
			input: "1",
			want:  "1",
		},
	}

	for _, tt := range tests {
		tt := tt

		if got := d2b(tt.input); tt.want != got {
			t.Fatalf("input: %s, want: %v, got: %v", tt.input, tt.want, got)
		}
	}
}
