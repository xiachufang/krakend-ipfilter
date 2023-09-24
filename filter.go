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
}

// NoopFilter noop, allow always, never deny
type NoopFilter struct{}

// Allow implement IPFilter.Allow
func (noop *NoopFilter) Allow(_ string) bool {
	return true
}

// CIDRFilter is an ip filter base on cidranger
type CIDRFilter struct {
	allowRanger cidranger.Ranger
	denyRanger  cidranger.Ranger
}

var emptyRanger cidranger.Ranger

func isIPV6(ip string) bool {
	return strings.IndexByte(ip, byte(':')) >= 0
}

func isCIDRIP(ip string) bool {
	return strings.IndexByte(ip, byte('/')) >= 0
}

func newRanger(ips []string) cidranger.Ranger {
	if len(ips) == 0 {
		return emptyRanger
	}
	ranger := cidranger.NewPCTrieRanger()
	for _, ip := range ips {
		if !isCIDRIP(ip) {
			if isIPV6(ip) {
				ip = fmt.Sprintf("%s/128", ip)
			} else {
				ip = fmt.Sprintf("%s/32", ip)
			}
		}
		_, ipNet, err := net.ParseCIDR(ip)
		if err != nil || ipNet == nil {
			continue
		}
		if err := ranger.Insert(cidranger.NewBasicRangerEntry(*ipNet)); err != nil {
			continue
		}
	}
	return ranger
}

// NewIPFilter create a cidranger base ip filter
func NewIPFilter(cfg *Config) IPFilter {
	if cfg == nil || (len(cfg.Deny) == 0 && len(cfg.Allow) == 0) {
		return &NoopFilter{}
	}
	return &CIDRFilter{
		allowRanger: newRanger(cfg.Allow),
		denyRanger:  newRanger(cfg.Deny),
	}
}

// Allow implement IPFilter.Allow
// 1. If the "allow" list is configured, only the IPs in the "allow" list are allowed access. All other IPs are denied access.
// 2. If the "deny" list is configured, the IPs in the "deny" list are denied access.
// 3. If both the "allow" and "deny" lists are configured, both rules are applied simultaneously.
func (f *CIDRFilter) Allow(ip string) bool {
	realIP := net.ParseIP(ip)
	if realIP == nil {
		return false
	}

	if f.denyRanger != emptyRanger {
		deny, err := f.denyRanger.Contains(realIP)
		if err != nil || deny {
			return false
		}
	}

	if f.allowRanger != emptyRanger {
		allow, err := f.allowRanger.Contains(realIP)
		if err != nil || !allow {
			return false
		}
	}

	return true
}
