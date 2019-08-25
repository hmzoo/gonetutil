package main

import (
	"./lan"
	"fmt"
)

func main() {

	fmt.Println("TEST Lan")

	l := lan.NewLan()
	l.SetIPNet("192.168.8.25/24")
  fmt.Println(l.GetIP())
	fmt.Println(l.GetMask())
  fmt.Println(l.GetFirstIP())
  fmt.Println(l.GetLastIP())

}
