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
	ListTaskSync	map[int]*lCommon.ListMonitor
	listError 		[]string
}


func NewMonitor() *Monitor {
    m := new(Monitor)
    m.ListTaskSync = make(map[int]*lCommon.ListMonitor)
    return m
}

func (self *Monitor) AddCoin(l lCommon.ListMonitor)(err error){
	// lCommon.ListMonitor{  Coin : "LRCBTC", Echange : "binance", Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBinance }
	if l.UpPerPercent == 0 		{ l.UpPerPercent = 0.02 }
	if l.DownPerPercent == 0 	{ l.DownPerPercent = 0.02 }	
	if l.UpLine == 0 			{ l.UpLine = 99999.99 }	
	if l.DownLine == 0 			{ l.DownLine = 0.0 }	
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
		self.ListTaskSync[m.Index] = &m
	}
	return nil
}

func (self *Monitor) Print()(err error){

	var soundFlagDown = 0
	var soundFlagUp = 0
	for i:=0 ; i < len(self.ListTaskSync); i++ {		

		color := self.priceComparison(&soundFlagUp, &soundFlagDown, i)

		self.listTask[i] = *self.ListTaskSync[i]

		lText.Print( lText.Line(*self.ListTaskSync[i], color ) )		
	}
	self.soundAllert(soundFlagUp, soundFlagDown)

	lText.ClPrint("\n", "white")
	lText.ClPrint("\n", "white")
	for _,i := range self.listError{
		 lText.ClPrint(i, "red")
	}
	self.listError = self.listError[:]
	return nil
}

func (self *Monitor) soundAllert(soundFlagUp int, soundFlagDown int){
	if(soundFlagUp == 1){
		go func(){
			if err := lCommon.PlayMusic("./sound/2-sirena-temnoe-vremya.mp3", 2 ) ; err != nil {
				self.listError = append(self.listError, err.Error() )
			}
		}()
	}
	if(soundFlagDown == -1){
		go func(){
			if err := lCommon.PlayMusic("./sound/obj_belltower.mp3", 2 ) ; err != nil {
				self.listError = append(self.listError, err.Error() )
			}
		}()
	}
}

func (self *Monitor) priceComparison(soundFlagUp* int, soundFlagDown* int, i int)([]string){
	var timeLimit = 90
	if self.ListTaskSync[i].PriceLast == 0{
		self.ListTaskSync[i].Time = timeLimit
		self.ListTaskSync[i].PriceLast = self.ListTaskSync[i].Price
		self.ListTaskSync[i].UpPer = 	self.ListTaskSync[i].PriceLast + self.ListTaskSync[i].PriceLast*self.ListTaskSync[i].UpPerPercent
		self.ListTaskSync[i].DownPer = 	self.ListTaskSync[i].PriceLast - self.ListTaskSync[i].PriceLast*self.ListTaskSync[i].DownPerPercent
	}else{
		if self.ListTaskSync[i].Time == 0{
			self.ListTaskSync[i].PriceLast = self.ListTaskSync[i].Price
			self.ListTaskSync[i].Time = timeLimit
			self.ListTaskSync[i].UpPer = 	self.ListTaskSync[i].PriceLast + self.ListTaskSync[i].PriceLast*self.ListTaskSync[i].UpPerPercent
			self.ListTaskSync[i].DownPer = 	self.ListTaskSync[i].PriceLast - self.ListTaskSync[i].PriceLast*self.ListTaskSync[i].DownPerPercent
		}else{
			self.ListTaskSync[i].Time = self.ListTaskSync[i].Time - 1
		}			
	}

	color := []string{"white","white","white","white","white","white","white","white","white"}
	if( self.ListTaskSync[i].Price > self.ListTaskSync[i].UpPer || self.ListTaskSync[i].Price > self.ListTaskSync[i].UpLine){
		color = []string{"white","white","white","green","white","white","white","white","white"}			
		*soundFlagUp = 1
	}else if (self.ListTaskSync[i].Price < self.ListTaskSync[i].DownPer || self.ListTaskSync[i].Price < self.ListTaskSync[i].DownLine) {
		color = []string{"white","white","white","red","white","white","white","white","white"}
		*soundFlagDown = -1
	}
	return color
}