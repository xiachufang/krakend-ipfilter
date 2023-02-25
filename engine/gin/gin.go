package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luraproject/lura/v2/config"
	"github.com/luraproject/lura/v2/logging"
	ipfilter "github.com/xiachufang/krakend-ipfilter/v2"
)

// Register register a ip filter middleware at gin
func Register(cfg *config.ServiceConfig, logger logging.Logger, engine *gin.Engine) {
	filterCfg := ipfilter.ParseConfig(cfg.ExtraConfig, logger)
	if filterCfg == nil {
		return
	}

	ipFilter := ipfilter.NewIPFilter(filterCfg)
	engine.Use(middleware(ipFilter, filterCfg, logger))
}

func middleware(ipFilter ipfilter.IPFilter, cfg *ipfilter.Config, logger logging.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ""
		for _, header := range cfg.IPHeaders {
			ip = ctx.Request.Header.Get(header)
			if ip != "" {
				break
			}
		}
		if ipFilter.Deny(ip) {
			logger.Error(fmt.Sprintf("krakend-ipfilter deny request from: %s", ip))
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}
