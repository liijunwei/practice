package main

// https://go.dev/tour/methods/19
import (
	"fmt"
	"testing"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() (int, error) {
	// return 100, nil
	return 100, &MyError{
		time.Now(),
		"it didn't work",
	}

}

func TestI8(t *testing.T) {
	value, err := run()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(value)
}
