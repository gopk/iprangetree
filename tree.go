//
// @project IPRangeTree 2016
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package iprangetree

import (
	"errors"
	"net"

	"github.com/google/btree"
)

// ErrInvalidItem message
var ErrInvalidItem = errors.New("Invalid IP range item")

// IPTree base
type IPTree struct {
	root *btree.BTree
}

// New tree object
func New(degree int) *IPTree {
	return &IPTree{root: btree.New(degree)}
}

// AddRange IPs vith value
func (t *IPTree) AddRange(ip1, ip2 net.IP, val interface{}) error {
	return t.Add(&IPItem{
		StartIP: ip1,
		EndIP:   ip2,
		Data:    val,
	})
}

// Add IP range item
func (t *IPTree) Add(item *IPItem) error {
	if nil == item || (nil == item.StartIP && nil == item.EndIP) {
		return ErrInvalidItem
	}
	if nil == item.StartIP {
		item.StartIP = item.EndIP
	}
	if nil == item.EndIP {
		item.EndIP = item.StartIP
	}
	t.root.ReplaceOrInsert(item)
	return nil
}

// Lookup item by IP
func (t *IPTree) Lookup(ip net.IP) (response *IPItem) {
	t.root.AscendGreaterOrEqual(IP(ip), func(item btree.Item) bool {
		switch compare(ip, item.(*IPItem).StartIP) {
		case 1:
			if compare(ip, item.(*IPItem).EndIP) <= 0 {
				response = item.(*IPItem)
				return false
			}
		case 0:
			response = item.(*IPItem)
			return false
		case -1:
			return false
		}
		return true
	})
	return
}
