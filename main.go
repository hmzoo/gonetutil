package main

import (
	"github.com/gocarina/gocsv"
  "gopkg.in/yaml.v2"
  "log"
	"io/ioutil"
	"os"
  "fmt"
  "./lan"
)

//Structure du fichier de blocks
type BlockData struct {
	NumID  int    `csv:"numid"`
	Block  string `csv:"block"`
	Config string `csv:"config"`
}

type BlockDatas []*BlockData


//Structure du fichier de configuration block
type NetworkConf struct {
  Name string
  Cidr int
  Vlanbase int
}

type BlockConf struct {
        Name string
        Gwpos string
        Networks []NetworkConf
}

//Structure du fichier de sortie
type NetData struct {

  NumID int `csv:"numid"`
  VlanID int `csv:"vlanid"`
	Name  string `csv:"name"`
  IPNet string `csv:"ipnet"`
  Mask string `csv:"mask"`
  Gateway string `csv:"gateway"`
  IP string `csv:"ip"`
  Cidr int `csv:"cidr"`
  Size int `csv:"size"`
  IPNum int `csv:"ipnum"`

}

type NetDatas []*NetData

var configs map[string]string
var output NetDatas

func main() {
	configs = make(map[string]string)
  output = NetDatas{}
  blks := LoadBlockDatas("blocks.csv")


  for _,blk := range(blks) {
  config := GetConfig(blk.Config)
  bconf := BlockConf{}
  err := yaml.Unmarshal([]byte(config), &bconf)
  if err != nil {
          log.Fatalf("error: %v", err)
  }

  BuildNetworks(blk.NumID,blk.Block,bconf )
  }

  SaveOutputCSV("networks.csv")

}

//Chargement du fichier blocks
func LoadBlockDatas(csvDatas string) BlockDatas {
	var data BlockDatas
	csvFile, err := os.OpenFile(csvDatas, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer csvFile.Close()
	if err := gocsv.UnmarshalFile(csvFile, &data); err != nil {
		log.Fatalf("error: %v", err)
	}
	return  data
}

//Recuperation du fichier de configuration block
func GetConfig(c string) string {
	if val, ok := configs[c]; ok {
		return val
	}
	b, err := ioutil.ReadFile(c + ".yml")
	if err != nil {
	    log.Fatalf("error: %v", err)
	}

  configs[c]= string(b)
	return string(b)

}

//Sauvegarde du fichier de sortie
func SaveOutputCSV(csvDatas string) error {
	os.Remove(csvDatas)
	resultFile, err := os.OpenFile(csvDatas, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer resultFile.Close()

	err = gocsv.MarshalFile(&output, resultFile)
	if err != nil {
		return err
	}

	return nil

}

//Construction des reseaux

func BuildNetworks(numid int,ipn string, b BlockConf){
  blan:= lan.NewLan()
  blan.SetIPNet(ipn)
  ip:=blan.IPNet.IP
  for _,n := range(b.Networks) {
    nlan := lan.NewLan()
    nlan.SetName(n.Name)
    nlan.SetVlanTag(n.Vlanbase+numid)
    nlan.SetIPCidr( ip.String(),n.Cidr)
    if(b.Gwpos == "first"){
      nlan.SetGateway(nlan.GetFirstIP())
    }else{
      nlan.SetGateway(nlan.GetLastIP())
    }
    ip = nlan.GetNextNetworkIP()
    fmt.Println(nlan.String())
    output=append(output,&NetData{
      NumID :numid,
      VlanID :nlan.VlanTag,
      Name : nlan.Name,
      IPNet :nlan.GetIPNet(),
      Mask : nlan.GetMask(),
      Gateway : nlan.GetGateway(),
      IP  :nlan.GetIP(),
      Cidr  :nlan.GetCidr(),
      Size  :nlan.Size(),
      IPNum :int(lan.Numeric(nlan.IPNet.IP))})
  }

}
