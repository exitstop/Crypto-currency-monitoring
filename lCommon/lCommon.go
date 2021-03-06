package lCommon

import (
	"io"
	"os"
	"os/exec"
	"runtime"
	"github.com/hajimehoshi/oto"
	"github.com/hajimehoshi/go-mp3"
	"time"
	// "strconv"
	"sync/atomic"
    "fmt"
)


var ChanLog *chan string
var numLog uint64

type ListMonitor struct{
	Index int
	Coin string
	Exchange string
	Price float64
	PriceLast float64	
	LogSavePrice float64	
	// EntryPrice float64
	PriceLastTick float64
	Time int
	UpPerPercent float64
	DownPerPercent float64
	UpPer float64
	DownPer float64
	UpLine float64
	DownLine float64
	Hodl float64
	HodlUsd float64
	Visible bool
	SoundOn bool
	CallBack func(string)(map[string]interface{}, error)
}


var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}

func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}


func PlayMusic(path string, div int64) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}
	defer d.Close()

	p, err := oto.NewPlayer(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}	
	defer p.Close()

	if _, err := io.CopyN(p, d, div * 8192 * 22  ); err != nil {
		return err
	}	

	return nil
}

func Log(str string){
	atomic.AddUint64(&numLog, 1)
	*ChanLog <- fmt.Sprintf("[%d] %s", numLog, str)
}

func LogStop(){
	*ChanLog <- "stop"
	numLog = 0
}


func LogPrint(){
	for {
	  lg := <- *ChanLog
	  if( lg == "stop" ){
	    break
	  }
	  fmt.Println("[log] ", lg)
	}
}

func CheckTimeoit(c1 chan time.Time){


    for {
      select {
        case msg1 := <- c1:
           fmt.Println("\nmain            time: ", msg1.Format("2006/01/02 15:04:05"))
        case <- time.After(time.Second*25):
          fmt.Println("\ntimeout")

          go LogPrint()
          LogStop()

          if err := PlayMusic("./sound/soundBelt.mp3", 3 ) ; err != nil {
          }
      }
    }
}

