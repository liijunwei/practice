package main

// https://go.dev/tour/methods/20
import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, error) {
	// if x < 0
	// 	return 0, &ErrNegativeSqrt(x).Error
	// else
		return math.Sqrt(x), nil
	// end

}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
