package ipfilter

import (
	"testing"

	"github.com/yl2chen/cidranger"
)

func TestCIDRFilter_Allow(t *testing.T) {
	type fields struct {
		allowRanger cidranger.Ranger
		denyRanger  cidranger.Ranger
	}
	type args struct {
		ip string
	}
	tests := []struct {
		name string
		cfg  *Config
		args map[string]bool
		want bool
	}{
		{
			name: "allow only",
			cfg: &Config{
				Allow: []string{
					"127.0.0.1",
					"127.0.0.2",
					"127.0.0.3",
					"192.168.1.0/24",
					"9ff1:4027:9788:741c:7c56:1970:227a:033e",
					"4ff1:4027:9788:741c:0000:0000:0000:0000/64",
				},
			},
			args: map[string]bool{
				"127.0.0.1":   true,
				"127.0.0.2":   true,
				"127.0.0.3":   true,
				"127.0.0.4":   false,
				"192.168.1.1": true,
				"191.168.1.1": false,
				"9ff1:4027:9788:741c:7c56:1970:227a:033e": true,
				"4ff1:4027:9788:741c:7c56:1970:227a:033e": true,
				"5ff1:4027:9788:741c:7c56:1970:227a:033e": false,
			},
		},
		{
			name: "deny only",
			cfg: &Config{
				Deny: []string{
					"127.0.0.1",
					"127.0.0.2",
					"127.0.0.3",
					"192.168.1.0/24",
					"9ff1:4027:9788:741c:7c56:1970:227a:033e",
					"4ff1:4027:9788:741c:0000:0000:0000:0000/64",
				},
			},
			args: map[string]bool{
				"127.0.0.1":   false,
				"127.0.0.2":   false,
				"127.0.0.3":   false,
				"127.0.0.4":   true,
				"192.168.1.1": false,
				"191.168.1.1": true,
				"9ff1:4027:9788:741c:7c56:1970:227a:033e": false,
				"4ff1:4027:9788:741c:7c56:1970:227a:033e": false,
				"5ff1:4027:9788:741c:7c56:1970:227a:033e": true,
			},
		},
		{
			name: "allow and deny",
			cfg: &Config{
				Allow: []string{
					"127.0.0.0/24",
					"4ff1:4027:9788:741c:0000:0000:0000:0000/64",
				},
				Deny: []string{
					"127.0.0.1",
					"4ff1:4027:9788:741c:7c56:1970:227a:033e",
				},
			},
			args: map[string]bool{
				"127.0.0.1":   false,
				"127.0.0.2":   true,
				"192.168.1.1": false,
				"4ff1:4027:9788:741c:7c56:1970:227a:033e": false,
				"4ff1:4027:9788:741c:7c56:1970:227a:133e": true,
				"9ff1:4027:9788:741c:7c56:1970:227a:133e": false,
			},
		},
		{
			name: "invalid ip",
			cfg: &Config{
				Allow: []string{
					"127.0.0.0/24",
				},
				Deny: []string{
					"127.0.0.1",
				},
			},
			args: map[string]bool{
				"":        false,
				"a.b.c.d": false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ranger := NewIPFilter(tt.cfg)
			for ip, want := range tt.args {
				got := ranger.Allow(ip)
				if got != want {
					t.Errorf("Allow(%s) = %v, want %v", ip, got, want)
				}
			}
		})
	}
}
