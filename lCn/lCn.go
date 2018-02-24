package lCn 


import (
	"net/http"
	// "net/url"
	"io/ioutil"
	"encoding/json"
	//"reflect"  //  typ := reflect.TypeOf(resp).Elem(); fmt.Println(typ) // определить тип элемента  
	// "fmt"
    )


func Connect(resp *http.Response ,err error) (map[string]interface{}, error) {
	if err != nil { 
		return nil,err
	  // panic(err)		
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var dat map[string]interface{}  // 
	if err := json.Unmarshal(body, &dat); err != nil {  // 
		return nil,err
	}
	json := dat["data"].(map[string]interface{})
	return json, nil
}

func GetPriceKucoin(coin string) (map[string]interface{}, error){
	resp, err := http.Get("https://api.kucoin.com/v1/open/tick?symbol=" + coin)
	return Connect(resp, err)	
}

func GetPriceBinance(coin string) (map[string]interface{}, error){
	resp, err := http.Get("https://www.binance.com/api/v3/ticker/price?symbol=" + coin)
	return Connect(resp, err)	
}

func GetPriceCryptopia(coin string) (map[string]interface{}, error){
	resp, err := http.Get("https://www.cryptopia.co.nz/api/GetMarket/" + coin)
	return Connect(resp, err)	
}

func GetPriceGate(coin string) (map[string]interface{}, error){
	resp, err := http.Get("http://data.gate.io/api2/1/ticker/" + coin)
	return Connect(resp, err)	
}