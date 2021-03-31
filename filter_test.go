package ipfilter

import (
	"fmt"
	"testing"
)

func TestCIDRFilter(t *testing.T) {
	cfg := &Config{
		Deny: []string{
			"160.136.0.0/16",
			"219.181.187.0/24",
			"192.168.1.0/24",
			"128.168.1.0/24",
			"8.8.8.8",
			"114.114.114.114",
			"9ff1:4027:9788:741c:7c56:1970:227a:033e",
			"4ff1:4027:9788:741c:0000:0000:0000:0000/64",
		},
	}
	f := NewIPFilter(cfg)
	testCases := map[string]bool{
		"160.137.1.2":  false,
		"1.1":          true,
		"8.8.8.8":      true,
		"128.168.2.1":  false,
		"128.168.1.11": true,
		"":             true,
		"d905:a82f:8d20:6005:c0d3:6e6c:7d6a:a11e": false,
		"9ff1:4027:9788:741c:7c56:1970:227a:033e": true,
		"4ff1:4027:9788:741c:7c56:1970:227a:033e": true,
	}
	for ip, v := range testCases {
		if r := f.Deny(ip); r != v {
			t.Fatal(fmt.Sprintf("%s should %v, but %v", ip, v, r))
		}
	}

	var cfg2 Config
	f = NewIPFilter(&cfg2)
	for ip := range testCases {
		if f.Deny(ip) {
			t.Fatal("noop filter should never deny: ", ip)
		}
	}
}
