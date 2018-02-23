package lCn 


import (
	"net/http"
	// "net/url"
	"io/ioutil"
	"encoding/json"
	"fmt"
    )

func GetPriceKucoin(coin string)map[string]interface{}{
	resp, err := http.Get("https://api.kucoin.com/v1/open/tick?symbol=" + coin)
	if err != nil { 
	  fmt.Println( "error PostForm" )
	  panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)


	var dat map[string]interface{}  // 
	if err := json.Unmarshal(body, &dat); err != nil {  // 
		fmt.Println( "error Unmarshal" )
	    panic(err)
	}
	json := dat["data"].(map[string]interface{})
	// fmt.Println( json["lastDealPrice"] )      // получаем значение buy: 0.11
	return json
}