package main

import(
  "fmt"
  "./lan"
)



func main(){

 fmt.Println("TEST netdata")

 _,lans := lan.LoadCSVNetDatas("data.csv")


 err:=lans.SaveCSV("data_r.csv")

 fmt.Println(err)




}
