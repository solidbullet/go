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
