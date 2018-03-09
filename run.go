package main

// парсинг произвольных данных на go
// https://blog.golang.org/json-and-go

import (
  "github.com/wsxiaoys/terminal"
  // "github.com/wsxiaoys/terminal/color"
  // "fmt"
  // "./lText"
  // "./lCn"
  "./lCommon"
  "./lMonitor"
  "time"
  // "runtime"
  // "net/http"
  // "net/url"
  // "io/ioutil"
  // "encoding/json"\
  // "./lKeyboard"
 //  "bytes"
  // "bufio"
  "fmt"
  // "os"
  // "strings"
  "net/http"

)

func main() {
  terminal.Stdout.Color("y")


  monitor := lMonitor.NewMonitor()

  // go func(m *lMonitor.Monitor){
  //   //     for {
  //   //     c := lKeyboard.Getch()
  //   //     switch {
  //   //     case bytes.Equal(c, []byte{3}):
  //   //         return
  //   //     case bytes.Equal(c, []byte{27, 91, 68}): // left
  //   //         // fmt.Println("LEFT pressed")
  //   //     default:
  //   //       m.Clear()
  //   //         // fmt.Println("Unknown pressed", c)
  //   //     }
  //   // }
  //  }(monitor)


  monitor.AddCoin(lCommon.ListMonitor{Coin: "BTCUSDT", Exchange: "binance", Price: 0.005, UpPerPercent: 0.001, DownPerPercent: 0.001, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0.15124802, HodlUsd: -1})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "ETHUSDT", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "ETCBTC", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "LTCUSDT", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "NEOUSDT", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "XRPBTC", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "LRCBTC", Exchange: "binance", Price: 0, UpPerPercent: 0.3, DownPerPercent: 0.3, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 583.416})
  // monitor.AddCoin(lCommon.ListMonitor{Coin: "AIONBTC", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "ADABTC", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "TRXBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "ELFBTC", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "NCASHBTC", Exchange: "binance", Price: 0.13,UpPerPercent: 0.13, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "QSPBTC", Exchange: "binance", Price: 0.1, UpPerPercent: 0.1, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "KNCBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  // monitor.AddCoin(lCommon.ListMonitor{Coin: "BCDBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "BTGBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "WAVESBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "CDTBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "XLMBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "BQXBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "ARNBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  // monitor.AddCoin(lCommon.ListMonitor{Coin: "ONTBNB", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "SUBBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "CNDBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  // monitor.AddCoin(lCommon.ListMonitor{Coin: "COBBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})



  // monitor.AddCoin(lCommon.ListMonitor{  Coin : "iotusd",   Exchange : "bitfinex",   Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBitfinex   } )
  // monitor.AddCoin(lCommon.ListMonitor{  Coin : "edousd",   Exchange : "bitfinex",   Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBitfinex   } )
  // monitor.AddCoin(lCommon.ListMonitor{  Coin : "gntusd",   Exchange : "bitfinex",   Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBitfinex   } )

  monitor.AddCoin(lCommon.ListMonitor{Coin: "ZPT-ETH", Exchange: "kucoin", Price: 0, UpPerPercent: 0.06, DownPerPercent: 0.06, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "OCN-BTC", Exchange: "kucoin", Price: 0, LogSavePrice: 0.00000132, UpPerPercent: 0.05, DownPerPercent: 0.04, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 102807.71})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "KEY-BTC", Exchange: "kucoin", Price: 0, UpPerPercent: 0.07, DownPerPercent: 0.07, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "TKY-BTC", Exchange: "kucoin", Price: 0, UpPerPercent: 0.06, DownPerPercent: 0.06, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 68340})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "EXY-BTC", Exchange: "kucoin", Price: 0,  LogSavePrice: 0, UpPerPercent: 0.09, DownPerPercent: 0.09, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  // monitor.AddCoin(lCommon.ListMonitor{Coin: "COFI-BTC", Exchange: "kucoin", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})

  monitor.AddCoin(lCommon.ListMonitor{Coin: "jnt_usdt", Exchange: "gate", Price: 0, UpPerPercent: 1, DownPerPercent: 1, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "nas_usdt", Exchange: "gate", Price: 0, UpPerPercent: 1, DownPerPercent: 1, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "dpy_usdt", Exchange: "gate", Price: 0, UpPerPercent: 1, DownPerPercent: 1, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})

  monitor.AddCoin(lCommon.ListMonitor{Coin: "BTC-XEM", Exchange: "bittrex", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  // monitor.AddCoin(lCommon.ListMonitor{Coin: "ETH-XEM", Exchange: "bittrex", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "BTC-ENRG", Exchange: "bittrex", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})

  monitor.AddCoin(lCommon.ListMonitor{Coin: "CEFS/USDT", Exchange: "cryptopia", Price: 0, UpPerPercent: 1, DownPerPercent: 1, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "EVR/BTC", Exchange: "cryptopia", Price: 0, UpPerPercent: 0.2, DownPerPercent: 0.4, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "HOLD/BTC", Exchange: "cryptopia", Price: 0, UpPerPercent: 1, DownPerPercent: 1, UpPer: 0, DownPer: 0, UpLine: 0.00000570, DownLine: 0, Hodl: 56925.80102849, LogSavePrice: 0.00000570})

  for {

    // lCommon.CallClear()
    monitor.GetPrice()
    lCommon.CallClear()
    monitor.Print()

    resp, err := http.Get("https://www.google.ru" )
    if err != nil{ 
      fmt.Println("error ----")
      if err := lCommon.PlayMusic("./sound/skypeDisconnect.mp3", 3 ) ; err != nil {
      }
    }
    resp.Body.Close()

    time.Sleep(10 * time.Second)
  }

  return

}
