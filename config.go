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
}

var (
	Namespace = "github_com/xiachufang/krakend-ipfilter"
	// default parse real ip from `X-Forwarded-For` or `X-Real-Ip`
	defaultIPHeaders = []string{"X-Forwarded-For", "X-Real-Ip"}
)

// ParseConfig build ip filter's Config
func ParseConfig(e config.ExtraConfig, logger logging.Logger) *Config {
	v, ok := e[Namespace].(map[string]interface{})
	if !ok {
		return nil
	}

	data, err := json.Marshal(v)
	if err != nil {
		logger.Error(fmt.Sprintf("marshal krakend-ipfilter config error: %v", err))
		return nil
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		logger.Error(fmt.Sprintf("unmarshal krakend-ipfilter config error: %v", err))
		return nil
	}
	if len(cfg.IPHeaders) == 0 {
		cfg.IPHeaders = defaultIPHeaders
	}
	return &cfg
}
