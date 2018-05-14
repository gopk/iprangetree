//
// @project IPRangeTree 2016 - 2018
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 - 2018
//

package iprangetree

import (
	"encoding/binary"
	"errors"
	"net"
	"strings"

	"github.com/google/btree"
)

// ErrInvalidItemParse by string
var ErrInvalidItemParse = errors.New("Invalid parse item")

// IPItem IP range
type IPItem struct {
	StartIPUint uint32
	StartIP     net.IP
	EndIP       net.IP
	Data        interface{}
}

// ItemByString parse
func ItemByString(s string) (item *IPItem, err error) {
	s = strings.Trim(s, " \t\n-–+")
	if strings.Contains(s, "-") {
		if arr := strings.Split(s, "-"); 2 == len(arr) {
			item = &IPItem{
				StartIP: net.ParseIP(arr[0]),
				EndIP:   net.ParseIP(arr[1]),
			}
		}
	} else if strings.Contains(s, "/") {
		if ip, inet, e := net.ParseCIDR(s); nil == e {
			item = &IPItem{
				StartIP: ip,
				EndIP:   lastIP(inet.IP, inet.Mask),
			}
		} else {
			err = e
		}
	} else {
		item = &IPItem{StartIP: net.ParseIP(s)}
		item.EndIP = item.StartIP
	}

	if err == nil {
		if item == nil || item.StartIP == nil || item.EndIP == nil {
			err = ErrInvalidItemParse
		} else if compare(item.StartIP, item.EndIP) > 0 {
			item.StartIP, item.EndIP = item.EndIP, item.StartIP
		}
	}
	return
}

// Less camparing for btree
func (i *IPItem) Less(then btree.Item) bool {
	switch ip := then.(type) {
	case *IPItem:
		if i.StartIPUint > 0 {
			if i.StartIPUint < ip.StartIPUint {
				return true
			}
			if ip.StartIPUint > 0 {
				return false
			}
		} else if ip.StartIPUint > 0 {
			return false
		}
		return compare(i.StartIP, ip.StartIP) == -1
	case IP:
		return compare(i.EndIP, net.IP(ip)) == -1
	}
	return false
}

// Compare with the second item
func (i *IPItem) Compare(it interface{}) int {
	switch ip := it.(type) {
	case *IPItem:
		if i.StartIPUint > 0 {
			if i.StartIPUint < ip.StartIPUint {
				return -1
			}
			if i.StartIPUint == ip.StartIPUint {
				return 0
			}
			if ip.StartIPUint > 0 {
				return 1
			}
		} else if ip.StartIPUint > 0 {
			return -1
		}
		return compare(i.StartIP, ip.StartIP)
	case IP:
		return compare(i.EndIP, net.IP(ip))
	case net.IP:
		return compare(i.EndIP, ip)
	}
	return 0
}

// Has IP in range
func (i *IPItem) Has(ip net.IP) bool {
	return compare(i.StartIP, ip) <= 0 && compare(i.EndIP, ip) >= 0
}

// Normalize IP values
func (i *IPItem) Normalize() {
	if i.StartIP != nil {
		if ip := i.StartIP.To4(); ip != nil {
			i.StartIPUint = ip2int(i.StartIP)
			i.StartIP = ip
		}
	}
	if i.EndIP != nil {
		if ip := i.EndIP.To4(); ip != nil {
			i.EndIP = ip
		}
	}
}

func ip2int(ip net.IP) uint32 {
	if len(ip) == 16 {
		return binary.BigEndian.Uint32(ip[12:16])
	}
	return binary.BigEndian.Uint32(ip)
}
