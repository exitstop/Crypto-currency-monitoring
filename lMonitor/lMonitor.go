package lMonitor 


import (
	// "net/http"
	// "net/url"
	// "io/ioutil"
	// "encoding/json"
	//"reflect"  //  typ := reflect.TypeOf(resp).Elem(); fmt.Println(typ) // определить тип элемента  
	"fmt"
	"../lText"
	"runtime"
	// "../lCn"
    )


type ListMonitor struct{
	Coin string
	Echange string
	Price float64
	UpPerPercent float64
	DownPerPercent float64
	UpPer float64
	DownPer float64
	UpLine float64
	DownLine float64
	Hodl float64
	CallBack func(string)(map[string]interface{}, error)
}



// ----------------------------------------------------------


type Monitor struct {
	listTask []ListMonitor
}

// ----------------------------------------------------------

func (self *Monitor) AddCoin(l ListMonitor)(err error){
	// ListMonitor{  Coin : "LRCBTC", Echange : "binance", Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBinance }
	if l.UpPerPercent == 0 		{ l.UpPerPercent = 0.05 }
	if l.DownPerPercent == 0 	{ l.DownPerPercent = 0.05 }
	self.listTask = append(self.listTask, l)
	return nil	
}

func (self *Monitor) GetPrice()(err error){
	for _, element := range self.listTask {
		// fmt.Print( index , " ")
		json, err := element.CallBack(element.Coin)
		if err != nil{
		  _, file, line, _ := runtime.Caller(0)
		  lText.ClPrint("error: GetPrice file: " + string(file) + " line: " + fmt.Sprintf("%d", line) + "\n" + err.Error() + "\n", "red")
		 	// fmt.Println( err )
		}else{
			element.Price =  json["lastDealPrice"].(float64)
			fmt.Println( fmt.Sprintf ( "%.8f",  element.Price ) )
		}
		// fmt.Println( element )
	}
	return nil
}