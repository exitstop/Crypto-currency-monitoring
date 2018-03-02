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
	"../lJsonLog"
	// "sync"
	// "../lCn"
	"os"
	"os/signal"
	"encoding/json"
    )



type Monitor struct {
	index int
	listTask 		[]lCommon.ListMonitor
	ListTaskSync	map[int]*lCommon.ListMonitor
	listError 		[]string
	Btcusdt float64
	MSignal chan os.Signal
	dbJson map[string]interface{}
}


func NewMonitor() *Monitor {
    m := new(Monitor)
    m.ListTaskSync = make(map[int]*lCommon.ListMonitor)

    m.MSignal = make(chan os.Signal, 1)
    signal.Notify(m.MSignal, os.Interrupt)

    go func(c chan os.Signal, m *Monitor){
    	for i := 0; i < 1; i++{
    		s := <-c
    		fmt.Println("Got signal:", s)

    		// b, _ := json.Marshal(m.dbJson)
    		// lJsonLog.WriteJson(b)
    	}
    	os.Exit(0)
    }(m.MSignal, m)    

    // var err error 
    // m.dbJson, err = lJsonLog.ReadJson();
    // if err != nil {
    // 	fmt.Println("Not found dat1")
    // 	m.dbJson = make(map[string]interface{})
    // 	m.dbJson["statistics"] = 555

    // 	b, _ := json.Marshal(m.dbJson)
    // 	lJsonLog.WriteJson(b)
    // }
    return m
}

func (self *Monitor) AddCoin(l lCommon.ListMonitor)(err error){
	// lCommon.ListMonitor{  Coin : "LRCBTC", Echange : "binance", Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBinance }
	if l.UpPerPercent == 0 		{ l.UpPerPercent = 0.03 }
	if l.DownPerPercent == 0 	{ l.DownPerPercent = 0.03 }	

	if (l.Coin == "BTCUSDT")	{ l.HodlUsd = -1 }

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
		        	if mp, ok := json["lastDealPrice"].(float64); ok{
		        		element.Price = mp
		        	}else{
		        		element.Price = 0
		        	}
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

	self.LogSave()

	return nil
}

func (self *Monitor) Print()(err error){

	var soundFlagDown = 0
	var soundFlagUp = 0

	formatLine := fmt.Sprintf("%3s %10s %10s %10s %10s %10s %10s %10s %3c %10s %10s\n", "№", "Exchange", "Coin", "Price", "Up", "Down", "ManUp", "ManDown", '%', "hodl", "USDT")
	lText.ClPrint(formatLine, "yellow")
	// fmt.Println(formatLine)

	for i:=0 ; i < len(self.ListTaskSync); i++ {		

		color := self.priceComparison(&soundFlagUp, &soundFlagDown, i)

		self.listTask[i] = *self.ListTaskSync[i]

		lText.Print( lText.Line(*self.ListTaskSync[i], self.Btcusdt, color ) )		
	}
	self.soundAllert(soundFlagUp, soundFlagDown)

	lText.ClPrint("\n", "white")
	lText.ClPrint("\n", "white")
	for _,i := range self.listError{
		 lText.ClPrint(i, "red")
	}
	self.listError = self.listError[:0]
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
	var timeLimit = 130
	if self.ListTaskSync[i].PriceLast == 0{
		self.ListTaskSync[i].Time = timeLimit
		self.ListTaskSync[i].PriceLast = self.ListTaskSync[i].Price
		self.ListTaskSync[i].PriceLastTick = self.ListTaskSync[i].Price
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
		if self.ListTaskSync[i].Time%3 == 1{
			self.ListTaskSync[i].PriceLastTick = self.ListTaskSync[i].Price
		}	
	}

	color := []string{"white","white","white","white","white","white","white","white","white","white","white"}

	if( self.ListTaskSync[i].Coin == "BTCUSDT") { self.Btcusdt = self.ListTaskSync[i].Price }

	if( self.ListTaskSync[i].Price != 0){
		if( self.ListTaskSync[i].Price > self.ListTaskSync[i].PriceLastTick){
			color = []string{"white","white","white","green","white","white","white","white","white","white","white"}			
		}else if (self.ListTaskSync[i].Price < self.ListTaskSync[i].PriceLastTick) {
			color = []string{"white","white","white","red","white","white","white","white","white","white","white"}
		}

		if( self.ListTaskSync[i].Price > self.ListTaskSync[i].UpPer || self.ListTaskSync[i].Price > self.ListTaskSync[i].UpLine){
			color = []string{"white","white","green","green","green","white","white","white","white","white","white"}			
			*soundFlagUp = 1
		}else if (self.ListTaskSync[i].Price < self.ListTaskSync[i].DownPer || self.ListTaskSync[i].Price < self.ListTaskSync[i].DownLine) {
			color = []string{"white","white","red","red","red","white","white","white","white","white","white"}
			*soundFlagDown = -1
		}
	}
	return color
}


func (self *Monitor) LogSave(){
	if _, ok := self.dbJson["statistics"]; !ok{
		fmt.Println("LogSave()")
		var err error 
		self.dbJson, err = lJsonLog.ReadJson();
		if err != nil {
			fmt.Println("Not found dat1")
			self.dbJson = make(map[string]interface{})
			self.dbJson["statistics"] = make(map[string]*interface{})
    		a := make(map[string]interface{})

    		for i:=0 ; i < len(self.ListTaskSync); i++ {
    			a[self.listTask[i].Coin] = self.ListTaskSync[i].Price
    		}

    		self.dbJson["statistics"] = a	

			b, _ := json.Marshal(self.dbJson)
			lJsonLog.WriteJson(b)
		}
	}else{
		for i:=0 ; i < len(self.ListTaskSync); i++ {

    		if mp, ok := self.dbJson["statistics"].(map[string]interface{}); ok{
    			if ( self.listTask[self.ListTaskSync[i].Index].LogSavePrice == 0) {
    				self.ListTaskSync[i].LogSavePrice = mp[self.listTask[i].Coin].(float64)
    			}else{
    				self.ListTaskSync[i].LogSavePrice = self.listTask[self.ListTaskSync[i].Index].LogSavePrice
    			}
    		}else{
    			// return 	nil, errors.New(nameFunction + "() Not found key dbJson['statistics']")
    		}
    	}
	}
}

func (self *Monitor) Clear(){
	if err := lCommon.PlayMusic("./sound/obj_belltower.mp3", 2 ) ; err != nil {
		// self.listError = append(self.listError, err.Error() )
	}
	for i:=0 ; i < len(self.ListTaskSync); i++ {
		self.ListTaskSync[i].Time = 0
	}
}