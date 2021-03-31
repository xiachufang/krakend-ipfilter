package gin

import (
	"fmt"
	"net/http"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/gin-gonic/gin"
	ipfilter "github.com/xiachufang/krakend-ipfilter"
)

// Register register a ip filter middleware at gin
func Register(cfg config.ServiceConfig, logger logging.Logger, engine *gin.Engine) {
	filterCfg := ipfilter.ParseConfig(cfg.ExtraConfig, logger)
	fmt.Printf("%+v\n", filterCfg)
	if filterCfg == nil {
		return
	}

	ipFilter := ipfilter.NewIPFilter(filterCfg)
	engine.Use(middleware(ipFilter, logger))
}

func middleware(ipFilter ipfilter.IPFilter, logger logging.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := ctx.ClientIP()
		if ipFilter.Deny(ip) {
			logger.Error(fmt.Sprintf("krakend-ipfilter deny request from: %s", ip))
			ctx.AbortWithStatus(http.StatusForbidden)
			return
		}

		ctx.Next()
	}
}
