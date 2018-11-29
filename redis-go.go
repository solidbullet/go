package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"time"

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

//type MemberSet struct {
//	MemberSet []Member
//}

var memebers [][]uint8

func mt5(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("read body err,%v %v\n", err, body)
		return
	}
	post := bytes.TrimRight(body, "\x00")
	fmt.Println(post)
	w.Write([]byte("ok"))
}

func mt4(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("read body err,%v %v\n", err, body)
		return
	}
	//post := bytes.TrimRight(body, "\x00")
	c, err := redis.Dial("tcp", "hiiboy.com:6379")
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
	}
	c.Do("AUTH", "810302")
	//c.Do("SADD", "mt5", post)
	//len, _ := c.Do("SCARD", "mt5")

	data, _ := c.Do("SMEMBERS", "mt5")
	get_list(data)

	var person Member
	for _, v := range memebers {
		//fmt.Println(byteString(v))
		json.Unmarshal([]byte(byteString(v)), &person)
		difftime := time.Now().Unix() + 8*60*60 - person.Localtime
		if difftime > 3600 {
			//SREM KEY MEMBER1
			res, _ := c.Do("SREM", "mt5", v)
			fmt.Println(res)
		}

	}
	//fmt.Println(memeber, time.Now().Unix(), len)
	defer c.Close()
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/mt5", mt5)
	http.HandleFunc("/mt4", mt4)
	http.ListenAndServe(":80", nil)
}

func get_list(value interface{}) { //redis查询结果不是[]interface{}
	switch v := value.(type) {
	case []interface{}:
		for _, u := range v {
			memebers = append(memebers, u.([]uint8))
		}
		break
	case nil:
		fmt.Println(v, "is nil", "null")
	case string:
		fmt.Println(v, "is string", v)
	case int:
		fmt.Println(v, "is int", v)
	case float64:
		fmt.Println(v, "is float64", v)

	case map[string]interface{}:
		fmt.Println(v, "is an map:")
		//print_map(value)
	default:
		fmt.Println(v, "is unknown type", fmt.Sprintf("%T", v))
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
