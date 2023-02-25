package ipfilter

import (
	"encoding/json"
	"fmt"

	"github.com/luraproject/lura/v2/config"
	"github.com/luraproject/lura/v2/logging"
)

// Config is config of ipfilter
type Config struct {
	Deny  []string `json:"deny"`
	Allow []string `json:"allow"`
	// header keys to parse real ip, default to []string{X-Forwarded-For, X-Real-Ip}
	IPHeaders []string `json:"ip_headers"`
	// if mode is "deny_all", all ip which not in allow list will deny.
	// if default mod is "allow_all", all ip which not in deny list will allow.
	// default is "deny_all"
	Mode string `json:"mode"`
}

var (
	Namespace = "github_com/xiachufang/krakend-ipfilter"
	defaultIPHeaders = []string{"X-Forwarded-For", "X-Real-Ip"}
	// ModeDeny deny all ip which not in allow ip list
	ModeDeny = "deny_all"
	// ModeAllow allow all ip which not in deny ip list
	ModeAllow = "allow_all"
)

// ParseConfig build ip filter's Config
func ParseConfig(e config.ExtraConfig, logger logging.Logger) *Config {
	v, ok := e[Namespace].(map[string]interface{})
	if !ok {
		return nil
	}

	data, err := json.Marshal(v)
	if err != nil {
		logger.Error(fmt.Sprintf("marshal krakend-ipfilter config error: %s", err.Error()))
		return nil
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		logger.Error(fmt.Sprintf("unmarshal krakend-ipfilter config error: %s", err.Error()))
		return nil
	}
	if len(cfg.IPHeaders) == 0 {
		cfg.IPHeaders = defaultIPHeaders
	}
	if cfg.Mode == "" {
		cfg.Mode = ModeDeny
	}
	return &cfg
}
