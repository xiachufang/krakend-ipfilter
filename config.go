package ipfilter

import (
	"encoding/json"
	"fmt"

	"github.com/luraproject/lura/config"
	"github.com/luraproject/lura/logging"
)

// Config is config of ipfilter
type Config struct {
	Deny  []string
	Allow []string
}

// Namespace is ipfilter's config key in extra config
const Namespace = "github_com/xiachufang/krakend-ipfilter"

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

	return &cfg
}
