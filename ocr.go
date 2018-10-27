package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
var coder = base64.NewEncoding(base64Table)

//"github.com/go-vgo/robotgo"
func main() {
	//	x, y := robotgo.GetMousePos()
	//	fmt.Println("pos:", x, y)
	//	color := robotgo.GetPixelColor(x, y)
	//	robotgo.MouseClick("center", true) //单击
	//fmt.Println("color----", "color")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=M0uwNKrDi9evnZToHUGYPSXU&client_secret=rqfGw2uupQtFrCFvRvTabPQ9xwditSc3",
		nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	result, _ := str2Map(string(body)) //百度返回的字符串转换为json格式
	token := result["access_token"]
	url_token := "https://aip.baidubce.com/rest/2.0/ocr/v1/numbers?access_token="
	url_token += token.(string) //接下来发送post请求所用的URL
	fmt.Println(url_token)

	file, err := os.Open("c:\\20181026220921.jpg")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	httpPost(url_token, Base64Encode(data))

}

func httpPost(myurl string, base64img []byte) {

	postValue := url.Values{
		"image_type": {"BASE64"},
		"group_id":   {"group001"},
		"user_id":    {"123456"},
		"image":      {string(base64img)},
	}

	postString := postValue.Encode()
	req, _ := http.NewRequest("POST", myurl, strings.NewReader(postString))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	nums := getValue(string(body))
	fmt.Println(nums)

}

func str2Map(jsonData string) (result map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(jsonData), &result)
	return result, err
}

func map2Str(mapData map[string]interface{}) (result string, err error) {
	resultByte, errError := json.Marshal(mapData)
	result = string(resultByte)
	err = errError
	return result, err
}

func Base64Encode(encode_byte []byte) []byte {
	return []byte(coder.EncodeToString(encode_byte))
}


func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func getValue(wholestr string) [10]int {
	nums := [10]int{}
	begin_0 := strings.Index(wholestr, "location")
	str_0 := wholestr[begin_0:]
	after_split := strings.Split(str_0, "location")
	for k, v := range after_split {
		if k > 0 {
			begin := strings.Index(v, "words")
			str1 := v[begin:]

			end := strings.Index(str1, "}")
			str2 := str1[:end-1]

			maohao_index := strings.Index(str2, ":")
			res := str2[maohao_index+3:]
			nums[k-1], _ = strconv.Atoi(res)
		}
	}
	return nums
}
