// Thanks Frank
// https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go

// it(struct tags) has no effect on the operation of this Go code.
// To use struct tags to accomplish something, other Go code must be written to examine structs at runtime.

package main

import "fmt"

type User1 struct {
	Name string `example:"name"`
}

func (u *User1) String() string {
	return fmt.Sprintf("Hi! My name is %s", u.Name)
}

func e5() {
	u := &User1{
		Name: "Sammy",
	}

	fmt.Println(u)
}
