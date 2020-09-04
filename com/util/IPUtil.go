package util

import (
	"fmt"
	"github.com/go-basic/ipv4"
	"github.com/go-errors/errors"
	"net"
	"strings"
)

type NetworkAddress struct {
	Host string
	Port string
}

/**
 * 获取本机ip
 */
func GetLocalIp() (ip string, ips interface{}) {
	ip = ipv4.LocalIP()
	ips, _ = ipv4.LocalIPv4s()
	return ip, ips
}

// ParseIPAddress parses an IP address and removes port and/or IPV6 format
func ParseIPAddress(input string) (string, error) {
	addr, err := SplitHostPort(input)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed to split network address '%s' by host and port", input))
	}

	ip := net.ParseIP(addr.Host)
	if ip == nil {
		return addr.Host, nil
	}

	if ip.IsLoopback() {
		if strings.Contains(addr.Host, ":") {
			// IPv6
			return "::1", nil
		}
		return "127.0.0.1", nil
	}

	return ip.String(), nil
}

/**
 * 拆分主机和端口号，没有则指定默认主机和端口
 */
func SplitHostPortDefault(input, defaultHost, defaultPort string) (NetworkAddress, error) {
	addr := NetworkAddress{
		Host: defaultHost,
		Port: defaultPort,
	}
	if len(input) == 0 {
		return addr, nil
	}

	start := 0
	// Determine if IPv6 address, in which case IP address will be enclosed in square brackets
	if strings.Index(input, "[") == 0 {
		addrEnd := strings.LastIndex(input, "]")
		if addrEnd < 0 {
			// Malformed address
			return addr, fmt.Errorf("Malformed IPv6 address: '%s'", input)
		}

		start = addrEnd
	}
	if strings.LastIndex(input[start:], ":") < 0 {
		// There's no port section of the input
		// It's still useful to call net.SplitHostPort though, since it removes IPv6
		// square brackets from the address
		input = fmt.Sprintf("%s:%s", input, defaultPort)
	}

	host, port, err := net.SplitHostPort(input)
	if err != nil {
		return addr, errors.New(fmt.Sprintf("net.SplitHostPort failed for '%s'", input))
	}

	if len(host) > 0 {
		addr.Host = host
	}
	if len(port) > 0 {
		addr.Port = port
	}

	return addr, nil
}

/**
 * 拆分主机和端口号
 */
func SplitHostPort(input string) (NetworkAddress, error) {
	if len(input) == 0 {
		return NetworkAddress{}, fmt.Errorf("Input is empty")
	}
	return SplitHostPortDefault(input, "", "")
}
