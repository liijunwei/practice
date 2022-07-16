package main

import (
	_ "github.com/go-redis/redis/v8" // _ 为空导入
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("Hello, go module mode")
	logrus.Println(uuid.NewString())
}
