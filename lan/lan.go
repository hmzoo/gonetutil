package lan

import(
  	"net"
    "strconv"
    "errors"
    "strings"
    "math"
)

type Lan struct {
  IPNet *net.IPNet
  Gateway net.IP
}

func NewLan() *Lan{
  return &Lan{}
}


func (lan *Lan) SetGateway( t string){
  lan.Gateway = net.ParseIP(t)
}

func (lan *Lan) SetIPNet( t string) error {
  _, ipv4Net, err := net.ParseCIDR(t)
  if(err != nil ){
    return err
  }
  lan.IPNet= ipv4Net
  return nil
}

func (lan *Lan) SetIPMasq( t string,m string) error {
 ipv4Net := net.ParseIP(t)
 ipm := net.ParseIP(m)
  if(ipv4Net == nil || ipm == nil){
    return errors.New("invalid mask or invalid ip")
  }
  lan.IPNet= &net.IPNet{IP:ipv4Net,Mask:net.IPv4Mask(ipm[0],ipm[1],ipm[2],ipm[3])}
  return nil
}

func (lan *Lan) SetIPCidr( t string,c int) error {
 ipv4Net := net.ParseIP(t)
 ipv4Mask := net.CIDRMask(c, 32)
  if(ipv4Net == nil || ipv4Mask == nil){
    return errors.New("invalid cidr or invalid ip")
  }
  lan.IPNet= &net.IPNet{IP:ipv4Net,Mask:ipv4Mask}
  return nil
}

func (l Lan) GetGateway() string {
	if l.Gateway == nil {
		return ""
	}
	return l.Gateway.String()
}

func (l Lan) GetIPNet() string {
	if l.IPNet == nil {
		return ""
	}
	return l.IPNet.String()
}

func (l Lan) GetIP() string {
	if l.IPNet == nil {
		return ""
	}
	return l.IPNet.IP.String()

}

func (l Lan) GetMask() string {
	if l.IPNet == nil {
		return ""
	}
	m := l.IPNet.Mask
	var s []string
	for _, v := range m {
		s = append(s, strconv.Itoa(int(v)))
	}
	return strings.Join(s, ".")
}

func (l Lan) GetCidr() int {
	if l.IPNet == nil {
		return 0
	}
	m := l.IPNet.Mask
	ones, _ := m.Size()
	return ones
}

func (l Lan) GetCountIP() int {
	if l.IPNet == nil {
		return 0
	}
	r := l.Size()
	if r > 2 {
		r = r - 2
	}
	return r
}

func (l Lan) FirstIP() net.IP {
	if l.IPNet == nil {
		return nil
	}
	r := l.Size()
	if r < 3 {
		return l.IPNet.IP
	}
	return NextIP(l.IPNet.IP, 1)
}

func (l Lan) GetFirstIP() string {
	ip := l.FirstIP()
  if(ip!=nil){
    return ip.String()
  }
  return ""
}

func (l Lan) LastIP() net.IP {
	if l.IPNet == nil {
		return nil
	}
	r := l.Size()
	if r < 3 {
		return NextIP(l.IPNet.IP, uint(r-1))
	}
	return NextIP(l.IPNet.IP, uint(r-2))
}

func (l Lan) GetLastIP() string {
	ip := l.LastIP()
  if(ip!=nil){
    return ip.String()
  }
  return ""
}

func (l Lan) Size() int {
	if l.IPNet == nil {
		return 0
	}
	m := l.IPNet.Mask
	ones, _ := m.Size()
	r := int(math.Pow(2, 32-float64(ones)))

	return r
}

//USEFULL
func NextIP(ip net.IP, inc uint) net.IP {
	i := ip.To4()
	v := uint(i[0])<<24 + uint(i[1])<<16 + uint(i[2])<<8 + uint(i[3])
	v += inc
	v3 := byte(v & 0xFF)
	v2 := byte((v >> 8) & 0xFF)
	v1 := byte((v >> 16) & 0xFF)
	v0 := byte((v >> 24) & 0xFF)
	return net.IPv4(v0, v1, v2, v3)
}
