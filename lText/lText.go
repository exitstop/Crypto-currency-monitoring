package lText 


import (
	// https://github.com/wsxiaoys/terminal/blob/master/color/color.go
      "github.com/wsxiaoys/terminal/color"
      "fmt"
      "../lCommon"
    )

var floatBlockSize = 8

var colorMap = map[string]string{
	"red"			: "r",
	"yellow"		: "y",
	"green"			: "g",
	"blue"			: "b",
	"magenta"		: "m",
	"white"			: "w",
	"cyan"			: "c",
	"gray"			: ".",
	"underlining"	: "_",	// подчеркивание
	"incline"		: "/",	// наклон
	"hidden"		: "-",	// наклон
}


func Color(color string, t string) string{
	return "@{" + colorMap[color] + t + "}"
}

func ClPrint(text string, cl string){		// распечать с заданным цветом
   color.Print(Color(cl, ""), text) 
}

func Cl(text string, cl string) string {  	// получить цвет
   return Color(cl, "") + text 
}


func Print(text string){					// печатать в цвете подготовленную строку
  color.Print(text)
}

func NullCount(number float64) int {
	var ret int = 0
	if number > 1{
		var n = int(number)
		for n%10 == 0 || n > 10 {
			n/=10
			ret = ret + 1
		}
		return ret
	}else{
		return ret
	}
}

func GBlock(number float64) string {
	tF := fmt.Sprintf("%c.%df", '%', floatBlockSize - NullCount(number))
	return fmt.Sprintf(tF, number)
}

// func Line(index int, exchange string, coin string, price float64, upPer float64, downPer float64, upLine float64, downLine float64, hodl float64, c []string ) string {
// 	index_ 		:= fmt.Sprintf("%.3d", index)
// 	exchange_	:= fmt.Sprintf("%.10s", exchange)
// 	coin_		:= fmt.Sprintf("%.10s", coin)
// 	price_		:= GBlock(price)
// 	upPer_		:= GBlock(upPer)
// 	downPer_	:= GBlock(downPer) 
// 	upLine_		:= GBlock(upLine) 
// 	downLine_	:= GBlock(downLine)
// 	hodl_		:= GBlock(hodl)
// 	return (Cl(index_, c[0]) + " " +Cl(exchange_,c[1]) + " " +Cl(coin_,c[2]) + " " +Cl(price_,c[3]) + " " +
// 		Cl(upPer_,c[4]) + " " +Cl(downPer_,c[5]) + " " + 	
// 		Cl(upLine_,c[6]) + " " +Cl(downLine_,c[7]) + " " + 
// 		Cl(hodl_,c[8]) + "\n")
// }

func Line(l lCommon.ListMonitor, Btcusdt float64, c []string) string {
	if l.Price == 0 { for i,_ := range c { c[i] = "magenta"} }
	index_ 		:= fmt.Sprintf("%3d", l.Index)
	exchange_	:= fmt.Sprintf("%10s", l.Exchange)
	coin_		:= fmt.Sprintf("%10s", l.Coin)
	l.Hodl 		= l.Hodl * l.Price
	if (l.HodlUsd != -1){	
		l.HodlUsd   = l.Hodl * Btcusdt 
	}
	price_		:= GBlock(l.Price);		if l.Price == 0.0 								{ c[3]="hidden"; }
	upPer_		:= GBlock(l.UpPer);		if l.UpPer == 99999.99	|| l.UpPerPercent == 1	{ c[4]="hidden"; }
	downPer_	:= GBlock(l.DownPer);	if l.DownPer == 0.0 || l.DownPerPercent == 1	{ c[5]="hidden"; } 
	upLine_		:= GBlock(l.UpLine);	if l.UpLine == 99999.99							{ c[6]="hidden"; }  
	downLine_	:= GBlock(l.DownLine);	if l.DownLine == 0								{ c[7]="hidden"; } 
	hodl_		:= GBlock(l.Hodl);		if l.Hodl == 0.0								{ c[8]="hidden"; }
	hodlUsd_	:= GBlock(l.HodlUsd);	if l.HodlUsd <= 0.0								{ c[9]="hidden"; }

	return (Cl(index_, c[0]) + " " +Cl(exchange_,c[1]) + " " +Cl(coin_,c[2]) + " " +Cl(price_,c[3]) + " " +
		Cl(upPer_,c[4]) + " " +Cl(downPer_,c[5]) + " " + 	
		Cl(upLine_,c[6]) + " " +Cl(downLine_,c[7]) + " " + 
		Cl(hodl_,c[8])+ " " + Cl(hodlUsd_,c[9]) + "\n")
}