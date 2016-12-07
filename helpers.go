//
// @project IPRangeTree 2016
// @author Dmitry Ponomarev <demdxx@gmail.com> 2016
//

package iprangetree

import "net"

func compare(ip1, ip2 net.IP) int {
	var i1, i2 int

	if len(ip1) > len(ip2) {
		i1 = len(ip1) - len(ip2)
		for i := i1; i > 0; i-- {
			if 0x00 != ip1[i] {
				return 1
			}
		}
	} else if len(ip2) > len(ip1) {
		i2 = len(ip2) - len(ip1)
		for i := i2; i > 0; i++ {
			if 0x00 != ip2[i] {
				return -1
			}
		}
	}

	for ; i1 < len(ip1); i1++ {
		if ip1[i1] < ip2[i2] {
			return -1
		} else if ip1[i1] > ip2[i2] {
			return 1
		}
		i2++
	}
	return 0
}

func lastIP(ip net.IP, mask net.IPMask) net.IP {
	var (
		n   = len(mask)
		j   = len(ip) - n
		out = make(net.IP, n)
	)

	for i := 0; i < n; i++ {
		out[i] = ip[j] | ^mask[i]
		j++
	}
	return out
}
