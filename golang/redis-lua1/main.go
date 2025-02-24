package main

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
)

const script = `
redis.call('SET', KEYS[1], ARGV[1])
return redis.call('GET', KEYS[1])
`

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()

	sha, err := client.ScriptLoad(ctx, script).Result()
	boom(err)

	fmt.Println("script sha1(script load):", sha)
	fmt.Println("script sha1(sha1sum):", sha1sum(script))

	key := "mykey"
	value := "myvalue"

	// execute script
	{
		result, err := executeScript(ctx, client, sha, []string{key}, []any{value})
		boom(err)
		fmt.Println("script execution result:", result)
	}

	// simulate Redis "SCRIPT FLUSH" cmd
	{
		err := client.ScriptFlush(ctx).Err()
		boom(err)
		fmt.Println("script cache flushed")
	}

	// execute script again, will trigger reload
	{
		result, err := executeScript(ctx, client, sha, []string{key}, []any{value})
		boom(err)
		fmt.Println("script execution result:", result)

	}
}

// handle NOSCRIPT error
func executeScript(
	ctx context.Context,
	client *redis.Client,
	sha string,
	keys []string,
	args []any,
) (any, error) {
	result, err := client.EvalSha(ctx, sha, keys, args...).Result()
	if err != nil {
		if err.Error() == "NOSCRIPT No matching script. Please use EVAL." {
			// reload script on NOSCRIPT error
			newSha, err := client.ScriptLoad(ctx, script).Result()
			boom(err)
			fmt.Println("script reloaded, new sha1sum:", newSha)

			return client.EvalSha(ctx, newSha, keys, args...).Result()
		}

		return nil, err
	}

	return result, nil
}

func boom(err error, msg ...string) {
	if err != nil {
		fmt.Println(strings.Join(msg, " "))
		panic(err)
	}
}

func assert(ok bool, msg ...string) {
	if !ok {
		panic(strings.Join(msg, " "))
	}
}

func sha1sum(data string) string {
	h := sha1.New()
	h.Write([]byte(data))

	return hex.EncodeToString(h.Sum(nil))
}
