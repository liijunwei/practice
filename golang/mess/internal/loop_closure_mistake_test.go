package internal

import (
	"fmt"
	"testing"
	"time"
)

// In Go, a loop closure refers to a common mistake that can occur when using closures inside loops.
// loop variable capture

var colors = []string{"red", "green", "blue"}

func TestLoopVariableCapture(t *testing.T) {
	fmt.Println("serial running")
	for _, c := range colors {
		fmt.Println(c)
	}

	fmt.Println("parallel running")
	// fix by either declare new local variable in loop or pass variable to anonymous function
	for _, c := range colors {
		go func() {
			fmt.Println(c)
		}()
	}

	time.Sleep(3 * time.Second)
}
