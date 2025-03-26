package singleton1

import "sync"

type OrderService struct{}

var instance *OrderService
var once sync.Once

// not thread safe
func NewOrderService1() *OrderService {
	if instance == nil {
		instance = &OrderService{}
	}
	return instance
}

func NewOrderService2() *OrderService {
	once.Do(func() {
		instance = &OrderService{}
	})

	return instance
}
