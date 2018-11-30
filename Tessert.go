package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "hiiboy.com:6379")
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
	}
	c.Do("AUTH", 810302)
	c.Do("SADD", 123, 222)
	data, _ := c.Do("SMEMBERS", 123)
	var memebers [][]uint8
	get_list(data, &memebers)
	for _, u := range memebers {

		fmt.Println(string(u))

	}

}

func get_list(value interface{}, memebers *[][]uint8) { //redis查询结果不是[]interface{}
	switch v := value.(type) {
	case []interface{}:
		fmt.Println(v, "is array")
		for _, u := range v {
			//*memebers = u.([]uint8)
			//fmt.Println(u)
			*memebers = append(*memebers, u.([]uint8))
		}
		break
	case nil:
		fmt.Println(v, "is nil")
	case string:
		fmt.Println(v, "is string")
	case int:
		fmt.Println(v, "is int")
	case float64:
		fmt.Println(v, "is float64")

	case map[string]interface{}:
		fmt.Println(v, "is an map:")
		//print_map(value)
	default:
		fmt.Println(v, "is unknown type", fmt.Sprintf("%T", v))
	}
}
