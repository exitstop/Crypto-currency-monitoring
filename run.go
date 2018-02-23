package main

// парсинг произвольных данных на go
// https://blog.golang.org/json-and-go

import (
  "github.com/wsxiaoys/terminal"
  // "github.com/wsxiaoys/terminal/color"
	"fmt"
  "./lText"
  "./lCn"
  // "net/http"
  // "net/url"
  // "io/ioutil"
  // "encoding/json"
)




func main() {
  terminal.Stdout.Color("y")

  sFormat := lText.Line(1, "binance", "btc_usdt", 100000.0, 1000.0, 1000.0, 1000.0, 1000.0, 1000.0, []string{"white","white","white","white","white","white","hidden","white","white"})
  lText.Print(sFormat)
  lText.Print(sFormat)
  lText.Print(sFormat)


  json := lCn.GetPriceKucoin("ZPT-ETH")
  fmt.Println( json["lastDealPrice"] )

  return

}
