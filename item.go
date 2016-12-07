//
// @project IPRangeTree 2016
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package iprangetree

import (
	"errors"
	"net"
	"strings"

	"github.com/google/btree"
)

// ErrInvalidItemParse by string
var ErrInvalidItemParse = errors.New("Invalid parse item")

// IPItem IP range
type IPItem struct {
	StartIP net.IP
	EndIP   net.IP
	Data    interface{}
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

	if nil == item || nil == item.StartIP || nil == item.EndIP {
		if nil == err {
			err = ErrInvalidItemParse
		}
	} else if compare(item.StartIP, item.EndIP) > 0 {
		item.StartIP, item.EndIP = item.EndIP, item.StartIP
	}
	return
}

// Less camparing for btree
func (i *IPItem) Less(then btree.Item) bool {
	switch ip := then.(type) {
	case *IPItem:
		return -1 == compare(i.EndIP, ip.StartIP) || -1 == compare(i.StartIP, ip.StartIP)
	case IP:
		return -1 == compare(i.EndIP, net.IP(ip))
	}
	return false
}

// Has IP in range
func (i *IPItem) Has(ip net.IP) bool {
	return compare(i.StartIP, ip) <= 0 && compare(i.EndIP, ip) >= 0
}
