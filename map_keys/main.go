package map_keys

import (
	"net"
	"sync"
)

func ipStrAsKey(ips []net.IP) *sync.Map {
	m := sync.Map{}
	for _, ip := range ips {
		m.Store(ip.String(), "foo")
	}
	return &m
}

func IPtoInt(i net.IP) uint32 {
	return uint32(i[0]<<24) + uint32(i[1]<<16) + uint32(i[2]<<8) + uint32(i[3])
}

func ipIntAsKey(ips []net.IP) *sync.Map {
	m := sync.Map{}
	for _, ip := range ips {
		m.Store(IPtoInt(ip), "foo")
	}
	return &m
}
