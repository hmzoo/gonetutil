package main

import(
  "fmt"
  "./lan"
)


type Binit struct {

    Structure string `csv:"structure" json:"structure"`
    Organization string `csv:"organization" json:"organization"`
    NumID int `csv:"numid" json:"numid"`
  	Block_admin  string `csv:"block_admin" json:"block_admin"`
    Block_peda  string `csv:"block_peda" json:"block_peda"`
}

type Binit []*Binit


func main(){

 fmt.Println("TEST netdata")

 nds:=lan.NewNetDatas()
 n:=nds.Create()
 n.Name="OKETO"
 n.VlanID=12
 lan := lan.NewLan()
 fmt.Println(lan.SetIPNet("192.168.10.0/24"))
 n.SetLan(lan)

 err:=nds.SaveCSV("data_build.csv")

 fmt.Println(err)




}
