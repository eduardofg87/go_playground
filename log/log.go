package main

import (
    "log"
    "os"
    logrous "github.com/sirupsen/logrus"

)
var (
    outfile, _ = os.Create("some.log") 
    l          = log.New(outfile, "", 0)
)

func main() {
    //first type of log
    l.Println("Some log")
    //second type of log
    logrous.WithFields(logrous.Fields{
      "animal": "walrus",
      "number": 1,
      "size":   10,
    }).Info("A walrus appears")   
    
}
