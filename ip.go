//
// @project IPRangeTree 2016
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package iprangetree

import (
	"net"

	"github.com/google/btree"
)

// IP implementation btree campare
type IP net.IP

// Less camparing for btree
func (ip IP) Less(then btree.Item) bool {
	switch v := then.(type) {
	case *IPItem:
		return -1 == compare(net.IP(ip), v.StartIP)
	case IP:
		return -1 == compare(net.IP(ip), net.IP(v))
	}
	return false
}
