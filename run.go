package main

import (
  "github.com/wsxiaoys/terminal"
  // "github.com/wsxiaoys/terminal/color"
	// "fmt"
  "./lText"
// "./lOut"
)

func main() {
  terminal.Stdout.Color("y")
  // i   биржа          монета       цена       верх граница +%               hold
  // 0  cryptopia      HOLD_BTC    0.00000252   ⟙  0.00000277  ⟘  0.00000227  0.143455538592


  sFormat := lText.Line(1, "binance", "btc_usdt", 100000.0, 1000.0, 1000.0, 1000.0, 1000.0, 1000.0, []string{"white","white","white","white","white","white","hidden","white","white"})
  lText.Print(sFormat)
  lText.Print(sFormat)
  lText.Print(sFormat)

  // sFormat = lText.Line(1, "binance", "btc_usdt", 1000.0, 1000.0, 1000.0, 1000.0, 1000.0, 1000.0)
  // lText.Out(sFormat)
  // lText.Out(sFormat)



  //fmt.Printf("%.10f", 0.000123)
  // lText.Cl("import text\n", "@r")
  // lText.Cl("import text\n", "red")
  return
  // color.Print("@rHello world")
  //color.Print("@kHello world")
  //color.Print("@gHello world")
  //color.Print("@yHello world")
  //color.Print("@bHello world")
  //color.Print("@rHello world")
  //color.Print("@mHello world")
  //color.Print("@cHello world")
  //color.Print("@wHello world")
  //color.Print("@dHello world")
  //color.Print("@!Hello world\n")
  //color.Print("@!Hello world\n")
  //color.Print("@.Hello world\n")
  //color.Print("@/Hello world\n")
  //color.Print("@_Hello world\n")

//	a := [...]string{"hello", "world!"}

//	for i := range a{
//		fmt.Println(i)
//	}

}
