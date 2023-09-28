package reverse

/*
reverse("hello")
reverse("ello") + "h"
reverse("llo") + "e"
reverse("lo") + "l"
reverse("o") + "l"
reverse("") + "o"
"" + "o"
*/
func reverse(input string) string {
	if input == "" {
		return ""
	}

	// runtime.Breakpoint()
	// fmt.Printf("input[0:1]: %v\n", input[0:1])
	return reverse(input[1:]) + input[0:1]
}
