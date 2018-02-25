package main

// парсинг произвольных данных на go
// https://blog.golang.org/json-and-go

import (
  "github.com/wsxiaoys/terminal"
  // "github.com/wsxiaoys/terminal/color"
	// "fmt"
  // "./lText"
  "./lCn"
  "./lMonitor"
  "./lCommon"
  "time"
  // "runtime"
  // "net/http"
  // "net/url"
  // "io/ioutil"
  // "encoding/json"
)




func main() {
  terminal.Stdout.Color("y")

  monitor := lMonitor.NewMonitor()

  monitor.AddCoin(lCommon.ListMonitor{  Coin : "BTCUSDT",   Exchange : "binance",   Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBinance   } )
  monitor.AddCoin(lCommon.ListMonitor{  Coin : "LRCBTC",    Exchange : "binance",   Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBinance   } )
  monitor.AddCoin(lCommon.ListMonitor{  Coin : "AIONBTC",   Exchange : "binance",   Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBinance   } )
  monitor.AddCoin(lCommon.ListMonitor{  Coin : "ZPT-ETH",   Exchange : "kucoin",    Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceKucoin    } )
  monitor.AddCoin(lCommon.ListMonitor{  Coin : "OCN-BTC",   Exchange : "kucoin",    Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceKucoin    } )
  monitor.AddCoin(lCommon.ListMonitor{  Coin : "KEY-BTC",   Exchange : "kucoin",    Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceKucoin    } )
  monitor.AddCoin(lCommon.ListMonitor{  Coin : "HOLD_BTC",  Exchange : "cryptopia", Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceCryptopia } )
  monitor.AddCoin(lCommon.ListMonitor{  Coin : "jnt_usdt",  Exchange : "gate",      Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceGate      } )


  for{
    monitor.GetPrice()
    lCommon.CallClear()
    monitor.Print()
    time.Sleep(2 * time.Second)
  }

  return

}
