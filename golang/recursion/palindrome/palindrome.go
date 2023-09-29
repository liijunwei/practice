package palindrome

func isPalindrome(input string) bool {
	if len(input) == 0 || len(input) == 1 {
		return true
	}

	lastIndex := len(input) - 1
	if input[0] != input[lastIndex] {
		return false
	}

	return isPalindrome(input[1:lastIndex])
}
