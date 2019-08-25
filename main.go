package main

import (
	"flag"
	"fmt"
  "os"
)

var version string
var appname string
var helpPtr *bool
var datafileinPtr *string
var datafileoutPtr *string


func main() {
  version ="0.1"
  appname ="netutil"
  helpPtr = flag.Bool("h", false, "get this help")
  datafileinPtr = flag.String("f", "data.csv", "data file")
  datafileoutPtr = flag.String("o", "result.csv", "data file")

    flag.Usage = usage

    flag.Parse()

    if *helpPtr {
      flag.Usage()
    }
/*
    if len(os.Args) < 2 {
      flag.Usage()
    }
*/
  fmt.Println("data file in :", *datafileinPtr)
  fmt.Println("data file out :", *datafileoutPtr)

}

func usage(){
  fmt.Fprintf(os.Stderr, "netutil v:"+version)
  fmt.Fprintf(os.Stderr, "\nUsage: %s [OPTION]... [ <COMMAND> [OPTION]... [id|rne]]\n", appname)
  flag.PrintDefaults()
  fmt.Fprintf(os.Stderr, "\n")
  os.Exit(0)
}
