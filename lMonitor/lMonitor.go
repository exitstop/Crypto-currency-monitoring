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
	"../lSymbols"
	// "sync"
	// "../lCn"
	"os"
	"os/signal"
	"encoding/json"
    )



type Monitor struct {
	index int
	listTask 				[]lCommon.ListMonitor
	ListTaskSync			map[int]*lCommon.ListMonitor
	ListTaskSyncHistory 	map[int]*lCommon.ListMonitor
	ListTaskSyncStr 		map[string]lCommon.ListMonitor
	listError 				[]string
	Btcusdt float64
	MSignal chan os.Signal
	dbJson map[string]interface{}
	mListEchange map[string]map[string]lSymbols.ListSymbolContent
}


func NewMonitor() *Monitor {
    m := new(Monitor)
    m.ListTaskSync 			= make(map[int]*lCommon.ListMonitor)
    m.ListTaskSyncHistory 	= make(map[int]*lCommon.ListMonitor)
    m.ListTaskSyncStr 	= make(map[string]lCommon.ListMonitor)

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

    m.mListEchange = make(map[string]map[string]lSymbols.ListSymbolContent) 

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

type StructItemReturn struct {
	ExchangeName string
	Data map[string]lSymbols.ListSymbolContent
	CallBack func()(map[string]lSymbols.ListSymbolContent, error)
}

func (self *Monitor) GetPrice()(err error){
	c := make(chan StructItemReturn)



	var exch []StructItemReturn
	exch = append(exch, StructItemReturn{ExchangeName: "binance", CallBack: lSymbols.GetListSumbolsBinance })
	exch = append(exch, StructItemReturn{ExchangeName: "kucoin", CallBack: lSymbols.GetListSumbolsKucoin  })


	for _, iNameExch := range exch {
		go func(iNameExch StructItemReturn, c chan StructItemReturn) {
		        json, err := iNameExch.CallBack()

		        if err != nil{
		          _, file, line, _ := runtime.Caller(0)
		          self.listError = append(self.listError, "error: GetPrice file: " + string(file) + " line: " + fmt.Sprintf("%d", line) + "\n" + err.Error() + "\n")

		        }else{
		        	iNameExch.Data = json
		        	c <- iNameExch
		        }		        
		    } (iNameExch, c);		
	}

	


	for range exch {
		d := <-c
		self.mListEchange[d.ExchangeName] = d.Data		
	}

	for _, item := range self.listTask {
		aa := self.mListEchange[item.Exchange]	
		item.Price = aa[item.Coin].Price
		item.Visible = true
		// item.Index = index
		m := item
		// self.ListTaskSync[item.Index] = &m
		str := fmt.Sprintf("%s-%s",item.Exchange, item.Coin)
		self.ListTaskSyncStr[str] = m
	}
	// for _, item := range self.listTask {
	// 	aa := self.mListEchange[item.Exchange]	
	// 	aa[item.Coin].Visible = true
	// 	aa[item.Coin].Index = item.Index
	// }
	// var gIndex int = 0
	// for index, item := range self.listTask {
	// 	aa := self.mListEchange[item.Exchange]	
	// 	item.Price = aa[item.Coin].Price
	// 	item.Visible = true
	// 	m := item
	// 	self.ListTaskSync[item.Index] = &m
	// 	gIndex = index
	// }
	gIndex := len(self.listTask)



	// if ( len(self.ListTaskSyncHistory) == 0 ){
		
		for index1, itEch := range self.mListEchange {
			for _, coinIt := range itEch {
				//for _, coinIt := range item {
					// if( coinIt.Visible != true ){				
						

						// if ok, mp := self.ListTaskSyncHistory[i], ok{

						
						str := fmt.Sprintf("%s-%s",index1, coinIt.SymbolDual)


						
						// aa := self.ListTaskSyncHistory[index1]

						if mp, ok := self.ListTaskSyncStr[str]; ok{
							m := mp
							// m.Index = mp.Index
							m.Price = coinIt.Price

							if(mp.Visible == true){
								m.Visible = true
							}
							self.ListTaskSyncStr[str] = m
						}else{
							m := lCommon.ListMonitor{  Index: gIndex, Coin : coinIt.SymbolDual, Exchange : index1, Price : coinIt.Price, UpPerPercent : 0.03, 
															DownPerPercent : 0.03, UpPer : 0, DownPer : 0, UpLine : 99999.99, DownLine : 0, Hodl : 0, Visible: false}
							self.ListTaskSyncStr[str] = m	
							gIndex ++

						}
						// fmt.Println(gIndex)
						
						
					// }
				//}
			}
		}

	// }

	// for a, b := range self.ListTaskSyncStr{
	// // for i:=0 ; i < len(self.ListTaskSync); i++ {
	// 	fmt.Println(a, b)
	// }


	count := 0	
	for _, b := range self.ListTaskSyncStr{
		m := b		
		self.ListTaskSync[m.Index] = &m
		// 
		count += m.Index
	}

	// for a, b := range self.ListTaskSyncStr{
	// 	fmt.Println( a, self.ListTaskSync[b.Index].Price ) 

	// }
	
	// for i:=0 ; i < len(self.ListTaskSyncStr); i++ {
	// 	 count -= i
	// }

	// for a, b := range self.ListTaskSyncStr{
	for i:=0 ; i < len(self.ListTaskSyncStr); i++ {
		fmt.Println(self.ListTaskSync[i] , " ---- ")
	}
	fmt.Println( "count " , count )
	fmt.Println( len(self.ListTaskSync), "---" , len(self.ListTaskSyncStr) )


	self.LogSave()

	return nil
}

func (self *Monitor) Print()(err error){

	var soundFlagDown = 0
	var soundFlagUp = 0

	formatLine := fmt.Sprintf("%3s %10s %10s %10s %10s %10s %10s %10s %3c %10s %10s\n", "№", "Exchange", "Coin", "Price", "Up", "Down", "ManUp", "ManDown", '%', "hodl", "USDT")
	lText.ClPrint(formatLine, "yellow")
	// fmt.Println("len: ",len(self.ListTaskSync))

	for i:=0 ; i < len(self.ListTaskSync); i++ {	
		// fmt.Println("i: ",i)
		color := self.priceComparison(&soundFlagUp, &soundFlagDown, i)

		str := fmt.Sprintf("%s-%s",self.ListTaskSync[i].Exchange, self.ListTaskSync[i].Coin)
		m := *self.ListTaskSync[i]
		self.ListTaskSyncStr[str] = m
		// self.ListTaskSyncHistory[i] = &m

		if( self.ListTaskSync[i].Visible == true || self.ListTaskSync[i].SoundOn == true	){
			if( self.ListTaskSync[i].Visible == true ) { 
				self.listTask[i] = *self.ListTaskSync[i]
			}
			lText.Print( lText.Line(*self.ListTaskSync[i], self.Btcusdt, color ) )	
		}	
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
	var timeLimit = 80

	self.ListTaskSync[i].SoundOn = false
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
			// self.ListTaskSync[i].SoundOn = true	
		}else if (self.ListTaskSync[i].Price < self.ListTaskSync[i].PriceLastTick) {
			color = []string{"white","white","white","red","white","white","white","white","white","white","white"}
			// self.ListTaskSync[i].SoundOn = true	
		}

		if( self.ListTaskSync[i].Price > self.ListTaskSync[i].UpPer || self.ListTaskSync[i].Price > self.ListTaskSync[i].UpLine){
			color = []string{"white","white","green","green","green","white","white","white","white","white","white"}	
			self.ListTaskSync[i].SoundOn = true			
			*soundFlagUp = 1
		}else if (self.ListTaskSync[i].Price < self.ListTaskSync[i].DownPer || self.ListTaskSync[i].Price < self.ListTaskSync[i].DownLine) {
			color = []string{"white","white","red","red","red","white","white","white","white","white","white"}
			self.ListTaskSync[i].SoundOn = true	
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

    		for i:=0 ; i < len(self.listTask); i++ {
    			a[self.listTask[i].Coin] = self.ListTaskSync[i].Price
    		}

    		self.dbJson["statistics"] = a	

			b, _ := json.Marshal(self.dbJson)
			lJsonLog.WriteJson(b)
		}
	}else{
		for i:=0 ; i < len(self.listTask); i++ {

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