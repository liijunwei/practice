package main

// https://go.dev/tour/methods/20
import (
	"fmt"
	"math"
	"testing"
)

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func TestI9(t *testing.T) {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
