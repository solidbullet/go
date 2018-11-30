package main

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

type Member struct {
	Ticket     int
	Symbol     string
	Entry      int
	T_type     int
	Lots       float64
	Sl         float64
	Tp         float64
	Magic      int
	PositionID int
	AccoundID  int
	Localtime  int64
}

type MemberSet struct {
	Slice []Member
}

func main() {
	var mm MemberSet
	c, err := redis.Dial("tcp", "hiiboy.com:6379")
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
	}
	c.Do("AUTH", "810302")
	//c.Do("SADD", "mt5", post)
	//len, _ := c.Do("SCARD", "mt5")

	data, _ := redis.Values(c.Do("SMEMBERS", "mt4"))
	mm.SetData(&data)
	fmt.Println(mm.Slice)
	defer c.Close()

}

func (m *MemberSet) SetData(redisData *[]interface{}) {
	var person Member
	for _, v := range *redisData {
		json.Unmarshal([]byte(v.([]byte)), &person)
		m.Slice = append(m.Slice, person)
	}
}

func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}
