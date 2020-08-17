package tracesite

import (
	"errors"
	"fmt"
	"net"
	"syscall"
	"time"

	"github.com/urfave/cli"
)

func socketAddr() ([4]byte, error) {
	socketAddr := [4]byte{0, 0, 0, 0}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return socketAddr, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if len(ipnet.IP.To4()) == net.IPv4len {
				copy(socketAddr[:], ipnet.IP.To4())
				return socketAddr, nil
			}
		}
	}
	err = errors.New("Not connected to the Internet")
	return socketAddr, err
}

func destAddr(dest string) ([4]byte, error) {
	destAddr := [4]byte{0, 0, 0, 0}
	addrs, err := net.LookupHost(dest)
	if err != nil {
		return destAddr, err
	}
	addr := addrs[0]

	ipAddr, err := net.ResolveIPAddr("ip", addr)
	if err != nil {
		return destAddr, err
	}
	copy(destAddr[:], ipAddr.IP.To4())
	return destAddr, nil
}

func Tracesite(options *cli.Context) error {

	ttl := options.Int("hop")
	tv := syscall.NsecToTimeval(1000 * 1000 * (int64)(options.Int("timeout")))
	retries := 0

	socketAddr, err := socketAddr()

	if err != nil {
		return err
	}

	destAddr, err := destAddr(options.Args().Get(0))
	if err != nil {
		return err
	}

	sendSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		return err
	}

	recvSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		return err
	}

	defer syscall.Close(recvSocket)
	defer syscall.Close(sendSocket)

	fmt.Printf("tracing [%v] - %v with packetSize=%v, maxHops=%v, startHop=%v, timeout=%v\n\n", options.Args().Get(0), destAddr, options.Int("packetsize"), options.Int("maxhops"), options.Int("hop"), options.Int("timeout"))
	for {
		start := time.Now()

		syscall.SetsockoptInt(sendSocket, 0x0, syscall.IP_TTL, ttl)
		syscall.SetsockoptTimeval(recvSocket, syscall.SOL_SOCKET, syscall.SO_RCVTIMEO, &tv)

		syscall.Bind(recvSocket, &syscall.SockaddrInet4{Port: options.Int("port"), Addr: socketAddr})
		syscall.Sendto(sendSocket, []byte{0x0}, 0, &syscall.SockaddrInet4{Port: options.Int("port"), Addr: destAddr})

		p := make([]byte, options.Int("packetsize"))
		destAddrString := fmt.Sprintf("%v.%v.%v.%v", destAddr[0], destAddr[1], destAddr[2], destAddr[3])

		n, from, err := syscall.Recvfrom(recvSocket, p, 0)

		elapsed := time.Since(start)
		if err == nil {
			retries = 0
			hop := Hop{Status: true, Addr: from, N: n, ElapsedTime: elapsed, TTL: ttl}
			if hop.IP() == destAddrString || ttl >= options.Int("maxhops") {
				break
			}
			ttl += 1
			fmt.Printf("\n%v. %v // [%v] // %v", hop.TTL, hop.Domain(), hop.IP(), hop.ElapsedTime)

		} else {
			hop := Hop{Status: false, N: n, ElapsedTime: elapsed, TTL: ttl}
			if retries < options.Int("retries") {
				if retries == 0 {
					fmt.Printf("\n%v. ", hop.TTL)
				}
				fmt.Printf("* ")
				retries += 1
			} else {
				retries = 0
				ttl += 1
				fmt.Println()
			}
		}

	}

	return nil

}
