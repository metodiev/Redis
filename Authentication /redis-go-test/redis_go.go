package main

import (
    "context"
    "fmt"
    "github.com/redis/go-redis/v9"
)

func main() {
    ctx := context.Background()
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    // Strings
    rdb.Set(ctx, "name", "Miroslav", 0)
    name, _ := rdb.Get(ctx, "name").Result()
    fmt.Println(name)  // Miroslav

    // Hash
    rdb.HSet(ctx, "user:1", map[string]interface{}{"name": "Alice", "age": 30})
    user, _ := rdb.HGetAll(ctx, "user:1").Result()
    fmt.Println(user)  // map[name:Alice age:30]

    // List
    rdb.LPush(ctx, "mylist", "a", "b", "c")
    list, _ := rdb.LRange(ctx, "mylist", 0, -1).Result()
    fmt.Println(list)  // [c b a]
}
