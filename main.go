// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// sudo dlv debug --headless --listen=:2345 --log --api-version=2 exec /mypath/binary -- eth0 4 192.168.0.2 443 1

package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	"github.com/mytechnotalent/turbo-attack/convert"
	"github.com/mytechnotalent/turbo-attack/routine"
	"github.com/mytechnotalent/turbo-attack/sudo"
)

func main() {
	if runtime.GOOS != "linux" {
		fmt.Println("application will only run on linux")
		return
	}

	sudo.Check()

	if len(os.Args) != 6 {
		fmt.Println("usage: turbo-attack_010_linux_arm64 <ethInterface> <ipVersion> <ip> <port> <count>")
		return
	}

	ethInterface := os.Args[1]
	ipVersion := os.Args[2]
	ip := os.Args[3]
	port := os.Args[4]
	count := os.Args[5]

	var wg sync.WaitGroup
	if ipVersion == "4" {
		ipv4Byte, portByte, countInt := convert.IPV4(&ethInterface, &ip, &port, &count)
		for i := 0; i < *countInt; i++ {
			wg.Add(1)
			routine.IPv4(&ethInterface, ipv4Byte, portByte)
			wg.Done()
		}
		wg.Wait()
	} else if ipVersion == "6" {
		ipv6Byte, portByte, countInt := convert.IPV6(&ethInterface, &ip, &port, &count)
		for i := 0; i < *countInt; i++ {
			routine.IPv6(&ethInterface, ipv6Byte, portByte)
		}
	} else {
		fmt.Println("valid: 4 or 6")
		return
	}
	wg.Wait()
}
