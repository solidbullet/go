package main

import (
	"encoding/json"
	"fmt"
)

var slice []map[string]interface{}

func main() {

	ocr := `{"log_id": 7864068706914929404, "words_result_num": 2, "words_result": [{"location": {"width": 57, "top": 5, "left": 0, "height": 14}, "words": "2592000"}, {"location": {"width": 65, "top": 36, "left": 0, "height": 15}, "words": "14575555"}]}`
	ocr1, _ := str2Map(ocr)
	print_map(ocr1)
	for _, v := range slice {
		fmt.Println(v["words"])
	}
}

//解析 map[string]interface{} 数据格式
func print_map(m map[string]interface{}) {
	slice = make([]map[string]interface{}, 0, 100)
	for k, v := range m {
		switch value := v.(type) {
		case nil:
			fmt.Println(k, "is nil", "null")
		case string:
			fmt.Println(k, "is string", value)
		case int:
			fmt.Println(k, "is int", value)
		case float64:
			fmt.Println(k, "is float64", value)
		case []interface{}:
			fmt.Println(k, "is an array:")
			//make(map[string]string, 10)
			for i, u := range value {
				fmt.Println(i, u)
				slice = append(slice, u.(map[string]interface{}))
			}
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			print_map(value)
		default:
			fmt.Println(k, "is unknown type", fmt.Sprintf("%T", v))
		}
	}
}

func str2Map(jsonData string) (result map[string]interface{}, err error) {
	err = json.Unmarshal([]byte(jsonData), &result)
	return result, err
}
