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
	cfg *Config
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
	if cfg == nil || (len(cfg.Deny) == 0 && len(cfg.Allow) == 0) {
		return &NoopFilter{}
	}
	return &CIDRFilter{
		allowRanger: newRanger(cfg.Allow),
		denyRanger:  newRanger(cfg.Deny),
		cfg: cfg,
	}
}

// Allow implement IPFilter.Allow
func (f *CIDRFilter) Allow(ip string) bool {
	netIP := net.ParseIP(ip)
	if netIP == nil {
		return f.cfg.Mode == ModeAllow
	}
	if f.cfg.Mode == ModeAllow {
		if deny, err := f.denyRanger.Contains(netIP); deny || err != nil {
			return false
		}
		return true
	}
	if allow, err := f.allowRanger.Contains(netIP); allow && err == nil {
		return true
	}
	return false
}