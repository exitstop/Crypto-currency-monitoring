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
	"../lCommon"
	// "sync"
	// "../lCn"
    )



type Monitor struct {
	index int
	listTask 		[]lCommon.ListMonitor
	ListTaskSync	map[int]lCommon.ListMonitor
	listError 		[]string
}


func NewMonitor() *Monitor {
    m := new(Monitor)
    m.ListTaskSync = make(map[int]lCommon.ListMonitor)
    return m
}

func (self *Monitor) AddCoin(l lCommon.ListMonitor)(err error){
	// lCommon.ListMonitor{  Coin : "LRCBTC", Echange : "binance", Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBinance }
	if l.UpPerPercent == 0 		{ l.UpPerPercent = 0.05 }
	if l.DownPerPercent == 0 	{ l.DownPerPercent = 0.05 }
	l.Index = self.index
	self.index = self.index + 1
	self.listTask = append(self.listTask, l)
	return nil	
}

func (self *Monitor) GetPrice()(err error){
	c := make(chan lCommon.ListMonitor)
	for _, element := range self.listTask {
		go func(element lCommon.ListMonitor, c chan lCommon.ListMonitor) {
				// defer waitGroup.Done()
		        json, err := element.CallBack(element.Coin)
		        // json, err := element.CallBack(element.Coin)
		        if err != nil{
		          _, file, line, _ := runtime.Caller(0)
		          self.listError = append(self.listError, "error: GetPrice file: " + string(file) + " line: " + fmt.Sprintf("%d", line) + "\n" + err.Error() + "\n")
		          // lText.ClPrint("error: GetPrice file: " + string(file) + " line: " + fmt.Sprintf("%d", line) + "\n" + err.Error() + "\n", "red")
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
		color := []string{"white","white","white","white","white","white","white","white","white"}
		lText.Print( lText.Line(self.ListTaskSync[i], color ) )		
	}
	lText.ClPrint("\n", "white")
	lText.ClPrint("\n", "white")
	for _,i := range self.listError{
		 lText.ClPrint(i, "red")
	}
	return nil
}