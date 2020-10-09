package vendormap

import (
	"crypto/md5"
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	macLength = 12
)

// MACShrink return if MAC is shrunk
func MACShrink(s string) string {
	return strings.ToLower(stripchars(s, ":-."))
}

// MACGlobal return MAC Global assigned
func MACGlobal(s string) bool {
	if len(s) < 2 {
		return false
	}

	fb, _ := strconv.ParseInt(s[0:2], 16, 64)

	return fb&2 == 0
}

// MACVendor return MAC vendor or ""
func MACVendor(s string) string {
	if len(s) < 6 {
		return ""
	}
	return ManufacturerMap[s[0:6]]
}

// MACReal check MAC for real
func MACReal(s string) bool {
	return MACGlobal(s) && MACVendor(s) != ""
}

// MACHash return if MAC is Hash
func MACHash(mac string) string {
	if len(mac) < 12 {
		return ""
	}
	m, _ := net.ParseMAC(fmt.Sprintf("%s.%s.%s", mac[0:4],
		mac[4:8], mac[8:12]))
	return fmt.Sprintf("%X", md5.Sum(m))
}

func stripchars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}
