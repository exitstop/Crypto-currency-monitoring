// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package lJsonLog

import (
    // "bufio"
    // "fmt"
    "io/ioutil"
    "encoding/json"
    // "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func WriteJson(b []byte){
    err := ioutil.WriteFile("dat1", b, 0644)
    check(err)
}

func ReadJson() map[string]interface{} {
    var ret map[string]interface{} 
    dat, err := ioutil.ReadFile("dat1")
    json.Unmarshal(dat, &ret) 
    check(err)    
    return ret
}

