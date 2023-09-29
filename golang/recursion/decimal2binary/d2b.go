package decimal2binary

import "fmt"

func d2b(input int, result string) string {
	if input == 0 && result == "" {
		return "0"
	}

	if input == 0 && result != "" {
		return result
	}

	result = fmt.Sprintf("%d", input%2) + result

	return d2b(input/2, result)
}
