//
// @project IPRangeTree 2016
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package iprangetree

import (
	"log"
	"net"
	"testing"
)

func TestTree(t *testing.T) {
	var (
		tree  = testTree()
		tests = [][]string{
			{"86.100.32.10", "86.100.32.0"},
			{"86.200.10.10", "86.200.10.0"},
			{"86.200.31.37", "86.200.10.0"},
			{"86.200.34.37", "86.200.30.0"},
			{"82.100.34.37", "82.100.20.0"},
		}
	)

	for _, ts := range tests {
		if r := tree.Lookup(net.ParseIP(ts[0])); nil != r {
			if ts[1] != r.Data.(string) {
				t.Errorf("Invalid IP [%s != %v] result", ts[1], r.Data)
			} else {
				t.Logf("IP %s correct", ts[1])
			}
		} else {
			t.Errorf("Invalid IP [%s] lookup", ts[0])
		}
	} // end for
}

func testTree() *IPTree {
	var (
		tree = New(2)
		ips  = []string{
			"86.100.32.0-86.100.32.255",
			"86.200.10.0-86.200.32.255",
			"86.200.30.0-86.200.36.255",
			"82.100.20.0-83.100.36.255",
			"99.100.10.0/24",
		}
	)

	for _, v := range ips {
		if item, err := ItemByString(v); nil != err {
			log.Println(err)
		} else {
			item.Data = item.StartIP.String()
			tree.Add(item)
		}
	}

	return tree
}
