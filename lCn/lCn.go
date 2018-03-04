package lCn 


import (
	"net/http"
	"errors"
	// "net/url"
	"io/ioutil"
	"encoding/json"
	//"reflect"  //  typ := reflect.TypeOf(resp).Elem(); fmt.Println(typ) // определить тип элемента  
	"strconv"
	// "time"
	// "fmt"
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
		return 	nil, errors.New(nameFunction + "() Not found key json['data']")
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
		return 	nil, errors.New(nameFunction + "() Not found key json['data']")
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
		return 	nil, errors.New(nameFunction + "() Not found key json['data']")
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
		return 	nil, errors.New(nameFunction + "() Not found key json['last']")
	}	
}

func GetPriceBittrex(coin string) (map[string]interface{}, error){
	nameFunction := "Bittrex"
	resp, err := http.Get("https://bittrex.com/api/v1.1/public/getticker?market=" + coin)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }
	if mp, ok := json["result"].(map[string]interface{}); ok{
		mp["lastDealPrice"] = mp["Last"]
		return 	mp, nil
	}else{
		return 	nil, errors.New(nameFunction + "() Not found key json['data']")
	}	
}

//  https://docs.bitfinex.com/v1/reference#rest-public-symbol-details
func GetPriceBitfinex(coin string) (map[string]interface{}, error){
	nameFunction := "Bitfinex"
	resp, err := http.Get("https://api.bitfinex.com/v1/pubticker/" + coin)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }
	if _, ok := json["last_price"]; ok{
		json["lastDealPrice"], err = strconv.ParseFloat(json["last_price"].(string), 64)
		if err != nil{ return nil, err }
		return 	json, nil
	}else{
		return 	nil, errors.New(nameFunction + "() Not found key json['last_price']")
	}	
}



// func GetListSubolsKucoin() {
// 	client := &http.Client{}

// 	req, _ := http.NewRequest("GET", "https://api.kucoin.com/v1/market/open/symbols", nil)

// 	resp, err := client.Do(req)

// 	if err != nil {
// 		fmt.Println("Errored when sending request to the server")
// 		return
// 	}

// 	defer resp.Body.Close()
// 	resp_body, _ := ioutil.ReadAll(resp.Body)

// 	fmt.Println(resp.Status)
// 	fmt.Println(string(resp_body))
// }