package main

// парсинг произвольных данных на go
// https://blog.golang.org/json-and-go

import (
	"github.com/wsxiaoys/terminal"
	// "github.com/wsxiaoys/terminal/color"
	// "fmt"
	// "./lText"
	"./lCn"
	"./lCommon"
	"./lMonitor"
  "time"
	// "./lKeyboard"
	// "runtime"
	// "net/http"
	// "net/url"
	// "io/ioutil"
	// "encoding/json"
)

func main() {
	terminal.Stdout.Color("y")

	monitor := lMonitor.NewMonitor()

	monitor.AddCoin(lCommon.ListMonitor{Coin: "BTCUSDT", Exchange: "binance", Price: 0.02, LogSavePrice:10246, UpPerPercent: 0.02, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0.03108776, HodlUsd: -1, CallBack: lCn.GetPriceBinance})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "ETHUSDT", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})
	monitor.AddCoin(lCommon.ListMonitor{Coin: "ETCBTC", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})
	monitor.AddCoin(lCommon.ListMonitor{Coin: "LTCUSDT", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})
	monitor.AddCoin(lCommon.ListMonitor{Coin: "NEOUSDT", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})
	monitor.AddCoin(lCommon.ListMonitor{Coin: "LRCBTC", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 583.416, CallBack: lCn.GetPriceBinance})
	monitor.AddCoin(lCommon.ListMonitor{Coin: "AIONBTC", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})
	monitor.AddCoin(lCommon.ListMonitor{Coin: "ADABTC", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "ELFBTC", Exchange: "binance", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "NCASHBTC", Exchange: "binance", Price: 0.13, LogSavePrice:0.00000271,UpPerPercent: 0.13, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "QSPBTC", Exchange: "binance", Price: 0.1, UpPerPercent: 0.1, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "KNCBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "BCDBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})
	monitor.AddCoin(lCommon.ListMonitor{Coin: "BTGBTC", Exchange: "binance", Price: 0.06, UpPerPercent: 0.06, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBinance})

	monitor.AddCoin(lCommon.ListMonitor{Coin: "BTC-XEM", Exchange: "bittrex", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceBittrex})

	// monitor.AddCoin(lCommon.ListMonitor{  Coin : "iotusd",   Exchange : "bitfinex",   Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBitfinex   } )
	// monitor.AddCoin(lCommon.ListMonitor{  Coin : "edousd",   Exchange : "bitfinex",   Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBitfinex   } )
	// monitor.AddCoin(lCommon.ListMonitor{  Coin : "gntusd",   Exchange : "bitfinex",   Price : 0, UpPerPercent : 0, DownPerPercent : 0, UpPer : 0, DownPer : 0, UpLine : 0, DownLine : 0, Hodl : 0, CallBack : lCn.GetPriceBitfinex   } )

	monitor.AddCoin(lCommon.ListMonitor{Coin: "ZPT-ETH", Exchange: "kucoin", Price: 0, UpPerPercent: 0.06, DownPerPercent: 0.06, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceKucoin})
	monitor.AddCoin(lCommon.ListMonitor{Coin: "OCN-BTC", Exchange: "kucoin", Price: 0.05, UpPerPercent: 0.05, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 24895, CallBack: lCn.GetPriceKucoin})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "KEY-BTC", Exchange: "kucoin", Price: 0.07, UpPerPercent: 0.07, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceKucoin})
	monitor.AddCoin(lCommon.ListMonitor{Coin: "TKY-BTC", Exchange: "kucoin", Price: 0, UpPerPercent: 0.06, DownPerPercent: 0.06, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceKucoin})
	// monitor.AddCoin(lCommon.ListMonitor{Coin: "COFI-BTC", Exchange: "kucoin", Price: 0, UpPerPercent: 0, DownPerPercent: 0, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceKucoin})

	monitor.AddCoin(lCommon.ListMonitor{Coin: "jnt_usdt", Exchange: "gate", Price: 0, UpPerPercent: 1, DownPerPercent: 1, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceGate})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "nas_usdt", Exchange: "gate", Price: 0, UpPerPercent: 1, DownPerPercent: 1, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceGate})
	monitor.AddCoin(lCommon.ListMonitor{Coin: "DPY_usdt", Exchange: "gate", Price: 0, UpPerPercent: 1, DownPerPercent: 1, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceGate})

  monitor.AddCoin(lCommon.ListMonitor{Coin: "CEFS_USDT", Exchange: "cryptopia", Price: 0, LogSavePrice: 1516, UpPerPercent: 0.05, DownPerPercent: 0.05, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceCryptopia})
	monitor.AddCoin(lCommon.ListMonitor{Coin: "EVR_BTC", Exchange: "cryptopia", Price: 0, UpPerPercent: 0.2, DownPerPercent: 0.4, UpPer: 0, DownPer: 0, UpLine: 0, DownLine: 0, Hodl: 0, CallBack: lCn.GetPriceCryptopia})
  monitor.AddCoin(lCommon.ListMonitor{Coin: "HOLD_BTC", Exchange: "cryptopia", Price: 0, UpPerPercent: 1, DownPerPercent: 1, UpPer: 0, DownPer: 0, UpLine: 0.00000570, DownLine: 0, Hodl: 56925.80102849, LogSavePrice: 0.00000570, CallBack: lCn.GetPriceCryptopia})

	for {
		monitor.GetPrice()
		lCommon.CallClear()
		monitor.Print()
		time.Sleep(4 * time.Second)
	}

	return

}
