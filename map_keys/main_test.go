package map_keys

import (
	"net"
	"testing"
)

var ips = []net.IP{
	net.IPv4(192, 168, 0, 0),
	net.IPv4(192, 168, 0, 1),
	net.IPv4(192, 168, 0, 2),
	net.IPv4(192, 168, 0, 2),
	net.IPv4(192, 168, 0, 3),
	net.IPv4(192, 168, 0, 3),
	net.IPv4(192, 168, 0, 3),
	net.IPv4(192, 168, 0, 3),
	net.IPv4(192, 168, 0, 3),
	net.IPv4(192, 168, 0, 3),
	net.IPv4(192, 168, 0, 3),
	net.IPv4(192, 168, 0, 3),
	net.IPv4(192, 168, 0, 3),
	net.IPv4(192, 168, 0, 3),
}

func BenchmarkStrAsKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ipStrAsKey(ips)
	}
}

func BenchmarkIntAsKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ipIntAsKey(ips)
	}
}

func BenchmarkGetStrAsKey(b *testing.B) {
	m := ipStrAsKey(ips)
	for i := 0; i < b.N; i++ {
		m.Load("192.168.0.1")
	}
}

func BenchmarkGetIntAsKey(b *testing.B) {
	m := ipIntAsKey(ips)
	ipInt := 3232235521 // 192.168.0.1
	for i := 0; i < b.N; i++ {
		m.Load(ipInt)
	}
}
