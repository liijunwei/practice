package main

import (
	"fmt"
	"time"
)

func e3() {
	m, _ := time.ParseDuration("1m31s")
	fmt.Printf("Take off in t-%.0f seconds.", m.Seconds())
}
