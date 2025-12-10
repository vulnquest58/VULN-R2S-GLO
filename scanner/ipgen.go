package scanner

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func IsPrivateIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip != nil && (ip.IsPrivate() || ip.IsLoopback())
}

func GenerateRandomPublicIP() string {
	for {
		ip := fmt.Sprintf("%d.%d.%d.%d",
			rand.Intn(223)+1,
			rand.Intn(256),
			rand.Intn(256),
			rand.Intn(254)+1,
		)
		if !IsPrivateIP(ip) {
			return ip
		}
	}
}
