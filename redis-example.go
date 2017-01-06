package main

import (
	"gopkg.in/redis.v5"
	"fmt"
)

func NewClient() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    // Output: PONG <nil>
	return client
}

func listOperation(client *redis.Client) {
    client.RPush("fruit", "apple") //在名称为 fruit 的list尾添加一个值为value的元素
    client.RPush("fruit", "orange")//在名称为 fruit 的list尾添加一个值为value的元素
    client.RPush("fruit", "or1")//在名称为 fruit 的list尾添加一个值为value的元素
    client.RPush("fruit", "oran2")//在名称为 fruit 的list尾添加一个值为value的元素
    client.LPush("fruit", "banana") //在名称为 fruit 的list头添加一个值为value的 元素
    length, err := client.LLen("fruit").Result() //返回名称为 fruit 的list的长度
    if err != nil {
        panic(err)
    }
    fmt.Println("length: ", length) // 长度为2

    val, err := client.LIndex("fruit",3).Result() //返回 fruit中， index=3的值
    if err != nil {
        panic(err)
    }
    fmt.Println(" index fruit: ", val)

    value, err := client.LPop("fruit").Result() //返回并删除名称为 fruit 的list中的首元素
    if err != nil {
        panic(err)
    }
    fmt.Println("fruit: ", value)

    value, err = client.RPop("fruit").Result() // 返回并删除名称为 fruit 的list中的尾元素
    if err != nil {
        panic(err)
    }
    fmt.Println("fruit: ", value)
}

func main() {
	client := NewClient()
	defer client.Close()

	listOperation(client)	
}


