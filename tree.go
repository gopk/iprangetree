//
// @project IPRangeTree 2016 - 2018
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 - 2018
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
	return t.Add(&IPItem{StartIP: ip1, EndIP: ip2, Data: val})
}

// Add IP range item
func (t *IPTree) Add(item *IPItem) error {
	if item == nil || (item.StartIP == nil && item.EndIP == nil) {
		return ErrInvalidItem
	}
	if item.StartIP == nil {
		item.StartIP = item.EndIP
	}
	if item.EndIP == nil {
		item.EndIP = item.StartIP
	}
	item.Normalize()
	t.root.ReplaceOrInsert(item)
	return nil
}

// Lookup item by IP
func (t *IPTree) Lookup(ip net.IP) (response *IPItem) {
	ipVal := ip2int(ip)
	t.root.AscendGreaterOrEqual(IP(ip), func(item btree.Item) bool {
		it := item.(*IPItem)
		switch compareExt(ip, ipVal, it) {
		case 1:
			if compare(ip, it.EndIP) <= 0 {
				response = item.(*IPItem)
				return false
			}
		case 0:
			response = it
			return false
		case -1:
			return false
		}
		return true
	})
	return
}
