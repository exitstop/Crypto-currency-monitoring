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

  sFormat := lText.Line(1, "binance", "btc_usdt", 100000.0, 1000.0, 1000.0, 1000.0, 1000.0, 1000.0, []string{"white","white","white","white","white","white","hidden","white","white"})
  lText.Print(sFormat)
  lText.Print(sFormat)
  lText.Print(sFormat)

  return

}
