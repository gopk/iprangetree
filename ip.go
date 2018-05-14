//
// @project IPRangeTree 2016 - 2018
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016 - 2018
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
		return compare(net.IP(ip), v.StartIP) == -1
	case IP:
		return compare(net.IP(ip), net.IP(v)) == -1
	}
	return false
}
