## krakend-ipfilter

![Test](https://github.com/xiachufang/krakend-ipfilter/workflows/Test/badge.svg) ![Lint](https://github.com/xiachufang/krakend-ipfilter/workflows/Lint/badge.svg)

---

IP filter middleware for [KrakenD(lura)](github.com/luraproject/lura) framework, base on [cidranger](https://github.com/yl2chen/cidranger)


## Example

Deny request from `192.168.0.0/16` but allow `192.168.1.1`:

```json
...
    "extra_config": {
        "github_com/xiachufang/krakend-ipfilter": {
            "allow": [
                "192.168.1.1"
            ],
            "deny": [
                "192.168.0.0/16"
            ]
        }
    }
...
```

## Test

```
make test
```

