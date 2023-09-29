package reverse

import "testing"

func Test_reverse(t *testing.T) {
	type test struct {
		input string
		want  string
	}

	tests := []test{
		{
			input: "",
			want:  "",
		},
		{
			input: "a",
			want:  "a",
		},
		{
			input: "ab",
			want:  "ba",
		},
		{
			input: "abc",
			want:  "cba",
		},
		{
			input: "hello",
			want:  "olleh",
		},
		{
			input: "hello world",
			want:  "dlrow olleh",
		},
	}

	for _, tt := range tests {
		tt := tt

		if got := reverse(tt.input); tt.want != got {
			t.Fatalf("input: %s, want: %s, got: %s", tt.input, tt.want, got)
		}
	}
}
