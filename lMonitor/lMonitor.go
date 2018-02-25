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
	// "sync"
	// "../lCn"
    )


type ListMonitor struct{
	Index int
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
	index int
	listTask 		[]ListMonitor
	ListTaskSync	map[int]ListMonitor
}


func NewMonitor() *Monitor {
    m := new(Monitor)
    m.ListTaskSync = make(map[int]ListMonitor)
    return m
}

// ----------------------------------------------------------

func (self *Monitor) AddCoin(l ListMonitor)(err error){
	// ListMonitor{  Coin : "LRCBTC", Echange : "binance", Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBinance }
	if l.UpPerPercent == 0 		{ l.UpPerPercent = 0.05 }
	if l.DownPerPercent == 0 	{ l.DownPerPercent = 0.05 }
	l.Index = self.index
	self.index = self.index + 1
	self.listTask = append(self.listTask, l)
	return nil	
}

// func (self *Monitor) GetPrice()(err error){
// 	for _, element := range self.listTask {
// 		json, err := element.CallBack(element.Coin)
// 		if err != nil{
// 		  _, file, line, _ := runtime.Caller(0)
// 		  lText.ClPrint("error: GetPrice file: " + string(file) + " line: " + fmt.Sprintf("%d", line) + "\n" + err.Error() + "\n", "red")
// 		}else{
// 			element.Price =  json["lastDealPrice"].(float64)
// 			fmt.Println( fmt.Sprintf ( "%.8f",  element.Price ) )
// 		}
// 	}
// 	return nil
// }


func (self *Monitor) GetPrice()(err error){
	c := make(chan ListMonitor)

	for _, element := range self.listTask {
		go func(element ListMonitor, c chan ListMonitor) {
				// defer waitGroup.Done()
		        json, err := element.CallBack(element.Coin)
		        // json, err := element.CallBack(element.Coin)
		        if err != nil{
		          _, file, line, _ := runtime.Caller(0)
		          lText.ClPrint("error: GetPrice file: " + string(file) + " line: " + fmt.Sprintf("%d", line) + "\n" + err.Error() + "\n", "red")
		          element.Price = 0
		          c <- element
		        }else{
		        	element.Price = json["lastDealPrice"].(float64)
		        	c <- element
		        	// element.Price =  json["lastDealPrice"].(float64)
		        	// fmt.Println( fmt.Sprintf ( "%.8f",  element.Price ) )
		        }
		        
		    } (element, c);		
	}
	
	for range self.listTask {
		m := <-c
		self.ListTaskSync[m.Index] = m
	}

	return nil
}

func (self *Monitor) Print()(err error){
	for i:=0 ; i < len(self.ListTaskSync); i++ {
		fmt.Println( self.ListTaskSync[i] )
	}
	return nil
}