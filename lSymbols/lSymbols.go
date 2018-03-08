package lSymbols

import (  
	// "fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"errors"
 )




func Connect(resp *http.Response) (map[string]interface{}, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{ return nil, err }
	var dat map[string]interface{}  // 
	if err := json.Unmarshal(body, &dat); err != nil {  errors.New( err.Error() ) ; return nil,err }
	return dat, nil
}

func Connect2(resp *http.Response) ([]interface{}, error) {
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{ return nil, err }
	var dat []interface{}  // 
	if err := json.Unmarshal(body, &dat); err != nil {  errors.New( err.Error() ) ; return nil,err }
	return dat, nil
}


type ListSymbolContent struct{
	CoinTypePair string     // ETH BTC USDT
	CoinPair string			// EXY
	SymbolDual string		// EXY-ETH
	Price float64
	Volume float64
	Trading bool			// true false
	Buy float64			// true false
	Sell float64			// true false
	Visible bool			// true false
	Index int			// true false
}

func TakeFloat(item map[string]interface{},key string) (float64) {
	if  mp, ok := item[key].(float64); ok{ 		
		return mp
	}
	return 0
}

func GetListSumbolsKucoin()  ( map[string]ListSymbolContent, error) {
	nameFunction := "GetListSubolsKucoin"
	retMapt := make(map[string]ListSymbolContent)
	resp, err := http.Get("https://api.kucoin.com/v1/market/open/symbols" )
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }
	if mp, ok := json["data"].( []interface{} ); ok{
		for _,item := range mp {
			item := item.(map[string]interface{})

			subIt := ListSymbolContent{	CoinTypePair : 		item["coinTypePair"].(string), 
										CoinPair 	 : 		item["coinType"].(string),
										SymbolDual 	 : 		item["symbol"].(string), 
										Price 		 : 		item["lastDealPrice"].(float64), 
										Trading 	 : 		item["trading"].(bool), 
										Buy		 	 : 		TakeFloat(item, "buy"), 
										Sell	 	 : 		item["sell"].(float64), 
										Visible	 	 : 		false, 
										Volume 		 : 		item["volValue"].(float64) }			

			retMapt[item["symbol"].(string)] = subIt
		}
		return 	retMapt, nil
	}else{
		return 	nil, errors.New(nameFunction + "() Not found key json['data']")
	}

}

func GetListSumbolsBinance()  ( map[string]ListSymbolContent, error) {
	nameFunction := "GetListSumbolsBinance"
	retMapt := make(map[string]ListSymbolContent)
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price" )
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect2(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }

	for _,item := range json {
		item := item.(map[string]interface{})

		price_, _ := strconv.ParseFloat(item["price"].(string), 64)
		// if(err != nil) { }
		subIt := ListSymbolContent{  SymbolDual 	 : 		item["symbol"].(string), 
									Price 		 : 		price_,
									Visible	 	 : 		false}	
		retMapt[item["symbol"].(string)] = subIt		

	}
	return 	retMapt, nil
}


func GetListSumbolsGateIo()  ( map[string]ListSymbolContent, error) {
	nameFunction := "GetListSumbolsGateIo"
	retMapt := make(map[string]ListSymbolContent)
	resp, err := http.Get("http://data.gate.io/api2/1/tickers" )
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }

	for name,item := range json {
		item := item.(map[string]interface{})
		subIt := ListSymbolContent{  SymbolDual 	 : 	name, 
									Price 		 : 		TakeFloat(item, "last"),
									Visible	 	 : 		false}	
		retMapt[name] = subIt
		
	}
	return 	retMapt, nil
}


func GetListSumbolsBittrex()  ( map[string]ListSymbolContent, error) {
	nameFunction := "GetListSumbolsBittrex"
	retMapt := make(map[string]ListSymbolContent)
	resp, err := http.Get("https://bittrex.com/api/v1.1/public/getmarketsummaries" )
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }

	if mp, ok := json["result"].([]interface{}); ok{
		for _,item := range mp {
			item := item.(map[string]interface{})

			name := item["MarketName"].(string)
			subIt := ListSymbolContent{	SymbolDual 	 : 		name, 
										Price 		 : 		TakeFloat(item, "Last"), 
										Visible	 	 : 		false}			

			retMapt[name] = subIt
		}
		return 	retMapt, nil
	}else{
		return 	nil, errors.New(nameFunction + "() Not found key json['result']")
	}

}


func GetListSumbolsBitfinex()  ( map[string]ListSymbolContent, error) {
	nameFunction := "GetListSumbolsBitfinex"
	retMapt := make(map[string]ListSymbolContent)
	resp, err := http.Get("https://api.bitfinex.com/v1/symbols_details" )
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }

	if mp, ok := json["result"].([]interface{}); ok{
		for _,item := range mp {
			item := item.(map[string]interface{})

			name := item["MarketName"].(string)
			subIt := ListSymbolContent{	SymbolDual 	 : 		name, 
										Price 		 : 		TakeFloat(item, "Last"), 
										Visible	 	 : 		false}			

			retMapt[name] = subIt
		}
		return 	retMapt, nil
	}else{
		return 	nil, errors.New(nameFunction + "() Not found key json['result']")
	}

}




func GetListSumbolsCryptopia()  ( map[string]ListSymbolContent, error) {
	nameFunction := "GetListSumbolsCryptopia"
	retMapt := make(map[string]ListSymbolContent)
	resp, err := http.Get("https://www.cryptopia.co.nz/api/GetMarkets" )
	if err != nil{ return nil, errors.New(nameFunction + "() -> Get() " + err.Error()) }
	json, err := Connect(resp)
	if err != nil{ return nil, errors.New(nameFunction + "() -> Connect() " + err.Error() ) }

	if mp, ok := json["Data"].([]interface{}); ok{
		for _,item := range mp {
			item := item.(map[string]interface{})

			name := item["Label"].(string)
			subIt := ListSymbolContent{	SymbolDual 	 : 		name, 
										Price 		 : 		TakeFloat(item, "LastPrice"), 
										Visible	 	 : 		false}			

			retMapt[name] = subIt
		}
		return 	retMapt, nil
	}else{
		return 	nil, errors.New(nameFunction + "() Not found key json['result']")
	}

}