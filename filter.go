package ipfilter

import (
	"fmt"
	"net"
	"strings"

	"github.com/yl2chen/cidranger"
)

// IPFilter is a interface for allow or deny an ip
type IPFilter interface {
	Allow(ip string) bool
	Deny(ip string) bool
}

// NoopFilter noop, allow always, never deny
type NoopFilter struct{}

// Allow implement IPFilter.Allow
func (noop *NoopFilter) Allow(_ string) bool {
	return true
}

// Deny implement IPFilter.Deny
func (noop *NoopFilter) Deny(_ string) bool {
	return false
}

// CIDRFilter is an ip filter base on cidranger
type CIDRFilter struct {
	allowRanger cidranger.Ranger
	denyRanger  cidranger.Ranger
}

func newRanger(ips []string) cidranger.Ranger {
	ranger := cidranger.NewPCTrieRanger()
	for _, ip := range ips {
		isCIDR := strings.IndexByte(ip, byte('/'))
		if isCIDR < 0 {
			ip = fmt.Sprintf("%s/24", ip)
		}
		_, ipNet, err := net.ParseCIDR(ip)
		if err != nil || ipNet == nil {
			continue
		}
		err = ranger.Insert(cidranger.NewBasicRangerEntry(*ipNet))
		if err != nil {
			continue
		}
	}
	return ranger
}

// NewIPFilter create a cidranger base ip filter
func NewIPFilter(cfg *Config) IPFilter {
	// always allow and never deny
	if cfg == nil || (len(cfg.Deny) == 0) {
		return &NoopFilter{}
	}
	return &CIDRFilter{
		allowRanger: newRanger(cfg.Allow),
		denyRanger:  newRanger(cfg.Deny),
	}
}

// Allow implement IPFilter.Allow
func (f *CIDRFilter) Allow(ip string) bool {
	netIP := net.ParseIP(ip)
	if netIP == nil {
		return false
	}
	if allow, err := f.allowRanger.Contains(netIP); allow && err == nil {
		return true
	}

	deny, err := f.denyRanger.Contains(netIP)
	if deny || err != nil {
		return false
	}
	return true
}

// Deny implement IPFilter.Deny
func (f *CIDRFilter) Deny(ip string) bool {
	return !f.Allow(ip)
}
