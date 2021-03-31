package gin

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/gin-gonic/gin"
	ipfilter "github.com/xiachufang/krakend-ipfilter"
)

func TestRegister(t *testing.T) {
	cfg := config.ServiceConfig{ExtraConfig: map[string]interface{}{
		ipfilter.Namespace: map[string]interface{}{
			"deny": []string{
				"127.0.0.1",
				"192.168.0.0/16",
			},
			"allow": []string{
				"192.168.1.1",
			},
		},
	}}

	gin.SetMode(gin.TestMode)
	eng := gin.New()
	Register(cfg, logging.NoOp, eng)

	eng.GET("/", func(ctx *gin.Context) {
		_, err := ctx.Writer.WriteString("ip: " + ctx.ClientIP())
		if err != nil {
			t.Log("write response error: ", err)
		}
	})
	testcases := map[string]int{
		"127.0.0.1":    http.StatusForbidden,
		"192.168.22.1": http.StatusForbidden,
		"192.168.1.1":  http.StatusOK,
		"123.11.12.3":  http.StatusOK,
		"4ff1:4027:9788:741c:7c56:1970:227a:033e": http.StatusOK,
	}
	for ip, excepted := range testcases {
		testSpecifiedIP(t, eng, ip, excepted)
	}
}

func testSpecifiedIP(t *testing.T, eng *gin.Engine, ip string, status int) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Log(err)
	}
	req.Header.Add("X-Forwarded-For", ip)

	w := httptest.NewRecorder()
	eng.ServeHTTP(w, req)
	body, err := ioutil.ReadAll(w.Result().Body)
	if err != nil {
		t.Log(err)
	}
	if w.Result().StatusCode != status {
		t.Fatal("ip filter test fail, with status code: ", w.Result().StatusCode, " body: ", string(body))
	}
}
