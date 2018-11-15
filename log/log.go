package main

import (
    "log"
    "os"
    logrous "github.com/sirupsen/logrus"

)
var (
    outfile, _ = os.Create("my.log") 
    l          = log.New(outfile, "", 0)
)

func main() {
    l.Println("hello, log!!!")
    logrous.WithFields(logrous.Fields{
      "animal": "walrus",
      "number": 1,
      "size":   10,
    }).Info("A walrus appears")   
}
