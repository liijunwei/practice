package singleton1

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderService1(t *testing.T) {
	os1 := NewOrderService1()
	os2 := NewOrderService1()
	assert.Equal(t, os1, os2, "Expected the same instance")
}

func TestNewOrderService2(t *testing.T) {
	wg := sync.WaitGroup{}
	count := 10
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("%p\n", NewOrderService1())
		}()
	}
	wg.Wait()
}
