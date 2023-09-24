## krakend-ipfilter

![Test](https://github.com/xiachufang/krakend-ipfilter/workflows/Test/badge.svg) ![Lint](https://github.com/xiachufang/krakend-ipfilter/workflows/Lint/badge.svg)

---

IP filter middleware for [KrakenD(lura)](https://github.com/luraproject/lura) framework, base on [cidranger](https://github.com/yl2chen/cidranger)

## Install

```shell
go get github.com/xiachufang/krakend-ipfilter/v2
```


## Usage

[Example](https://github.com/xiachufang/krakend-ipfilter/tree/master/example)

```go
package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/luraproject/lura/v2/config"
	"github.com/luraproject/lura/v2/logging"
	"github.com/luraproject/lura/v2/proxy"
	krakendgin "github.com/luraproject/lura/v2/router/gin"
	"github.com/luraproject/lura/v2/transport/http/client"
	http "github.com/luraproject/lura/v2/transport/http/server"
	ipfilter "github.com/xiachufang/krakend-ipfilter/v2/engine/gin"
)

func main() {
	port := flag.Int("p", 0, "Port of the service")
	logLevel := flag.String("l", "ERROR", "Logging level")
	debug := flag.Bool("d", false, "Enable the debug")
	configFile := flag.String("c", "/etc/krakend/configuration.json", "Path to the configuration filename")
	flag.Parse()

	parser := config.NewParser()
	serviceConfig, err := parser.Parse(*configFile)
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	serviceConfig.Debug = serviceConfig.Debug || *debug
	if *port != 0 {
		serviceConfig.Port = *port
	}

	logger, err := logging.NewLogger(*logLevel, os.Stdout, "[KRAKEND]")
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	engine := gin.Default()

	// register krakend-ipfilter
	ipfilter.Register(&serviceConfig, logger, engine)

	routerFactory := krakendgin.NewFactory(krakendgin.Config{
		Engine:         engine,
		ProxyFactory:   proxy.NewDefaultFactory(proxy.CustomHTTPProxyFactory(client.NewHTTPClient), logger),
		Middlewares:    []gin.HandlerFunc{},
		Logger:         logger,
		HandlerFactory: krakendgin.EndpointHandler,
		RunServer:      http.RunServer,
	})

	routerFactory.New().Run(serviceConfig)
}

```

## Config Example

### Only allow specified IP
> Complete example configuration file: [allow.json](https://github.com/xiachufang/krakend-ipfilter/blob/master/example/allow.json)

```json
    "extra_config": {
        "github_com/xiachufang/krakend-ipfilter": {
            "allow": [
                "192.168.1.1",
                "8.8.8.8",
                "127.0.0.1/8"
            ]
        }
    }
```


### Only deny specified IP
> Complete example configuration file: [deny.json](https://github.com/xiachufang/krakend-ipfilter/blob/master/example/deny.json)

```json
    "extra_config": {
        "github_com/xiachufang/krakend-ipfilter": {
            "deny": [
                "192.168.1.1",
                "8.8.8.8",
                "127.0.0.1/8"
            ]
        }
    }
```

### Allow IP within a range but deny a specific one
> Complete example configuration file: [deny_allow.json](https://github.com/xiachufang/krakend-ipfilter/blob/master/example/deny_allow.json)

```json
    "extra_config": {
        "github_com/xiachufang/krakend-ipfilter": {
            "allow": [
                "127.0.0.0/24"
            ],
            "deny": [
                "127.0.0.1"
            ]
        }
    }
```


## Test

```
make test
```

