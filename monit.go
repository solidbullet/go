package main

import (
	"bytes"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//"github.com/garyburd/redigo/redis"
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

var currency, gold, oil string

func monit(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("read body err,%v %v\n", err, body)
		return
	}
	post := bytes.TrimRight(body, "\x00")
	fmt.Println(string(post))
	w.Write([]byte(body))
}
func getorders(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	accountid := req.Form.Get("accountid")
	currency = "waihui"
	gold = "huangjin"
	oil = "yuanyou"
	//w.Header().Add("Access-Control-Allow-Origin", "*")
	if accountid == "1" {
		w.Write([]byte(currency))
		//fmt.Println(currency)
	} else if accountid == "2" {
		w.Write([]byte(gold))
	} else if accountid == "3" {
		w.Write([]byte(oil))
	} else {
		w.Write([]byte("accountid no exist"))
	}
}
func main() {

	//	fmt.Println("---------------------------------------------------------------------------")
	//	var mm MemberSet
	//	c, err := redis.Dial("tcp", "hiiboy.com:6379")
	//	if err != nil {
	//		fmt.Println("connect to redis err", err.Error())
	//	}
	//	c.Do("AUTH", "810302")
	//c.Do("SADD", "mt5", post)
	//len, _ := c.Do("SCARD", "mt5")

	//	ticker := time.NewTicker(time.Millisecond * 100000)
	//	go func() {
	//		for t := range ticker.C {
	//			//以下每隔一段时间读取redis
	//			data, _ := redis.Values(c.Do("SMEMBERS", "mt4"))
	//			mm.SetData(&data)
	//			fmt.Println(mm.Slice, t)
	//		}
	//	}()
	//	defer c.Close()
	http.HandleFunc("/getorders", getorders)
	http.HandleFunc("/monit", monit)
	http.ListenAndServe(":80", nil)

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
