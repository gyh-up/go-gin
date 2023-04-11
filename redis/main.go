package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)
var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "name", "123", 0).Err()
	if err != nil {
		panic(any(err))
	}

	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(any(err))
	}
	fmt.Println("name", val)

	val2, err := rdb.Get(ctx, "age").Result()
	if err == redis.Nil {
		fmt.Println("age does not exist")
	} else if err != nil {
		panic(any(err))
	} else {
		fmt.Println("age", val2)
	}
	// Output: key value
	// key2 does not exist
}

type Course struct {
	Cid int
	Cname string
}
