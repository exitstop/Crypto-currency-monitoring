package lCn 


import (
	"net/http"
	"errors"
	// "net/url"
	"io/ioutil"
	"encoding/json"
	//"reflect"  //  typ := reflect.TypeOf(resp).Elem(); fmt.Println(typ) // определить тип элемента  
	// "fmt"
	"strconv"
	// "time"
    )


func Connect(resp *http.Response) (map[string]interface{}, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{ return nil, err }
	var dat map[string]interface{}  // 
	if err := json.Unmarshal(body, &dat); err != nil {  errors.New( err.Error() ) ; return nil,err }
	return dat, nil
}

func GetPriceKucoin(coin string) (map[string]interface{}, error){
	nameFunction := "GetPriceKucoin"
	resp, err := http.Get("https://api.kucoin.com/v1/open/tick?symbol=" + coin)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }
	if mp, ok := json["data"].(map[string]interface{}); ok{
		return 	mp, nil
	}else{
		return 	nil, errors.New(nameFunction + "() Not foun key json['data']")
	}
}

func GetPriceBinance(coin string) (map[string]interface{}, error){
	nameFunction := "GetPriceBinance"
	resp, err := http.Get("https://www.binance.com/api/v3/ticker/price?symbol=" + coin)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }
	if _, ok := json["price"]; ok{
		json["lastDealPrice"], err = strconv.ParseFloat(json["price"].(string), 64)
		if err != nil{ return nil, err }
		return 	json, nil
	}else{
		return 	nil, errors.New(nameFunction + "() Not foun key json['data']")
	}
}

func GetPriceCryptopia(coin string) (map[string]interface{}, error){
	nameFunction := "GetPriceCryptopia"
	resp, err := http.Get("https://www.cryptopia.co.nz/api/GetMarket/" + coin)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }
	if mp, ok := json["Data"].(map[string]interface{}); ok{
		mp["lastDealPrice"] = mp["LastPrice"]
		return 	mp, nil
	}else{
		return 	nil, errors.New(nameFunction + "() Not foun key json['data']")
	}
}

func GetPriceGate(coin string) (map[string]interface{}, error){
	nameFunction := "GetPriceGate"
	resp, err := http.Get("http://data.gate.io/api2/1/ticker/" + coin)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }
	if _, ok := json["last"].(float64); ok{
		json["lastDealPrice"] = json["last"].(float64)
		return 	json, nil
	}else{
		return 	nil, errors.New(nameFunction + "() Not foun key json['last']")
	}	
}




