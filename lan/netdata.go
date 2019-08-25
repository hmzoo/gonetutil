package lan

import(
    "os"
    "github.com/gocarina/gocsv"

)

type NetData struct {

  Structure string `csv:"structure" json:"structure"`
  Organization string `csv:"organization" json:"organization"`
  NumID int `csv:"numid" json:"numid"`
	Name  string `csv:"name" json:"name"`
  IPNet string `csv:"ipnet" json:"ipnet"`
  IP string `csv:"ip" json:"ip"`
  Cidr int `csv:"cidr" json:"cidr"`
  Mask string `csv:"mask" json:"mask"`
  Gateway string `csv:"gateway" json:"gateway"`
  VlanID int `csv:"vlanid" json:"vlanid"`
}

type NetDatas []*NetData

func (nd *NetData) SetLan(l *Lan){
  nd.IPNet=l.GetIPNet();
  nd.IP=l.GetIP();
  nd.Cidr=l.GetCidr();
  nd.Mask=l.GetMask();

}

func (nd *NetData) Populate(){
  if(nd.IPNet!=""){
    lan := NewLan()
    if(lan.SetIPNet(nd.IPNet)!=nil){return}
    nd.SetLan(lan)
    return
  }
}

func (nds *NetDatas) Populate(){
for _,n :=range(*nds){
  n.Populate();
}
}

func NewNetDatas() NetDatas {
  return NetDatas{}
}
func (nds *NetDatas) Create() *NetData {
  nd := &NetData{}
  (*nds)=append((*nds),nd)
  return nd
}

func LoadCSVNetDatas(csvnetDatas string) (error, *NetDatas) {
	var data NetDatas
	csvFile, err := os.OpenFile(csvnetDatas, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err, nil
	}
	defer csvFile.Close()
	if err := gocsv.UnmarshalFile(csvFile, &data); err != nil {
		return err, nil
	}
	return nil, &data
}

func (data NetDatas) SaveCSV(csvnetDatas string) error {
	os.Remove(csvnetDatas)
	resultFile, err := os.OpenFile(csvnetDatas, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer resultFile.Close()

	err = gocsv.MarshalFile(&data, resultFile)
	if err != nil {
		return err
	}

	return nil

}
