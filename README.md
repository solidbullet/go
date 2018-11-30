# redis-class中可以把redis查询出来的数据变为对象   
#要用redis.value：data, _ := redis.Values(c.Do("SMEMBERS", "mt4")) 
#生成json：  
	var s Serverslice  
	b, err := json.Marshal(s)   
#解析成json：  
	var s Serverslice   
	str := `{"servers":[{"serverName":"SH_VPN","serverIP":"127.0.0.1"},{"serverName":"BJ_VPN","serverIP":"127.0.0.2"}]}`    
	json.Unmarshal([]byte(str), &s)     
	fmt.Println(s)     

#定时器:
	ticker := time.NewTicker(time.Millisecond * 1000)  
	go func() {  
		for t := range ticker.C {  
			//以下每隔一段时间读取redis  
			fmt.Println(t)  
		}  
	}()  	
	
#解析get及post请求  
	err := req.ParseForm()  
	if err != nil {  
		panic(err)  
	}  
	accountid := req.Form.Get("accountid")  
	  
	body, err := ioutil.ReadAll(req.Body)  //post请求  
