package main

// парсинг произвольных данных на go
// https://blog.golang.org/json-and-go

import (
  "github.com/wsxiaoys/terminal"
  // "github.com/wsxiaoys/terminal/color"
	"fmt"
  "./lText"
  "./lCn"
  "./lMonitor"
  "runtime"
  // "net/http"
  // "net/url"
  // "io/ioutil"
  // "encoding/json"
)




func main() {
  terminal.Stdout.Color("y")

  sFormat := lText.Line(1, "binance", "LRCBTC", 100000.0, 1000.0, 1000.0, 1000.0, 1000.0, 1000.0, []string{"white","white","white","white","white","white","hidden","white","white"})
  lText.Print(sFormat)
  lText.Print(sFormat)
  lText.Print(sFormat)

  lMonitor.ListMonitorInit(lMonitor.ListMonitor{  Coin : "LRCBTC", Echange : "binance", Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBinance })
  // fmt.Println( l ) 

  lMonitor.Go()

  json, err := lCn.GetPriceKucoin("ZPT-ETH")
  if err != nil{
    _, file, line, _ := runtime.Caller(0)
    lText.ClPrint("error: GetPriceKucoin file: " + string(file) + " line: " + fmt.Sprintf("%d", line) + "\n", "red")
  }else{
    fmt.Println( json["lastDealPrice"] )
  }

  return

}
