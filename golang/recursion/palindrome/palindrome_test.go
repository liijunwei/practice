package palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	type test struct {
		input string
		want  bool
	}

	tests := []test{
		{
			input: "",
			want:  true,
		},
		{
			input: "a",
			want:  true,
		},
		{
			input: "ab",
			want:  false,
		},
		{
			input: "abc",
			want:  false,
		},
		{
			input: "aba",
			want:  true,
		},
		{
			input: "kayak",
			want:  true,
		},
		{
			input: "hello world",
			want:  false,
		},
		{
			input: "racecar",
			want:  true,
		},
	}

	for _, tt := range tests {
		tt := tt

		if got := isPalindrome(tt.input); tt.want != got {
			t.Fatalf("input: %s, want: %v, got: %v", tt.input, tt.want, got)
		}
	}
}
