package main

import (
	"context"
	"fmt"
	"runtime"

	"github.com/redis/go-redis/v9"
)

// https://redis.io/docs/latest/develop/data-types/probabilistic/bloom-filter/
// possible false positive
// imposible false negative

// brew services stop redis
// docker run -p 6379:6379 -it --rm redis/redis-stack-server:latest
// go run redis-bloom-filter/main.go
func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()

	re, err := client.Ping(ctx).Result()
	boom(err)

	fmt.Println("ping", re)

	res1, err := client.BFReserve(ctx, "bikes:models", 0.01, 1000).Result()
	if err != nil && err.Error() != "ERR item exists" {
		boom(err)
	}

	fmt.Println("res1", res1) // >>> OK

	res2, err := client.BFAdd(ctx, "bikes:models", "Smoky Mountain Striker").Result()
	boom(err)

	fmt.Println("res2", res2) // >>> true

	res3, err := client.BFExists(ctx, "bikes:models", "Smoky Mountain Striker").Result()
	boom(err)

	fmt.Println("res3", res3) // >>> true

	res4, err := client.BFMAdd(ctx, "bikes:models",
		"Rocky Mountain Racer",
		"Cloudy City Cruiser",
		"Windy City Wippet",
	).Result()
	boom(err)

	fmt.Println("res4", res4) // >>> [true true true]

	res5, err := client.BFMExists(ctx, "bikes:models",
		"Rocky Mountain Racer",
		"Cloudy City Cruiser",
		"Windy City Wippet",
	).Result()

	boom(err)

	fmt.Println("res5", res5) // >>> [true true true]
}

func boom(err error) {
	if err != nil {
		_, filename, line, _ := runtime.Caller(1)
		fmt.Println(fmt.Sprintf("raised from: %s:%d", filename, line))
		panic(err)
	}
}
