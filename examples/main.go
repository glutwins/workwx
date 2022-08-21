package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/glutwins/workwx/store"
	"github.com/glutwins/workwx/wxcommon"
	"github.com/glutwins/workwx/wxsuite"
	"github.com/go-redis/redis/v8"
)

type handler struct {
	wxsuite.DummySuiteCallbackHandler
	tokenCache store.TokenCache
}

func (h handler) OnCallbackSuiteTicket(raw *wxcommon.XmlRxEnvelope, base *wxsuite.SuiteCallbackBase, ticket string) {
	h.tokenCache.SetSuiteTicket(raw.ToUserName, ticket)
}

func main() {
	tokenCache := store.NewRedisTokenStore("suite", &redis.Options{Addr: "127.0.0.1:6379"})
	sc := wxsuite.NewSuiteClient("suiteId", "suiteSecret", tokenCache)
	fmt.Println(sc.GetSuiteToken())
	r := gin.Default()
	g := r.Group("/v1")
	wxsuite.RegisterSuiteHandler(g, &wxsuite.SuiteConfig{
		SuiteId:        "suiteId",
		SuiteSecret:    "suiteSecret",
		Token:          "token",
		EncodingAESKey: "encodingAESKey",
	}, &handler{tokenCache: tokenCache}, &wxsuite.DummySuiteMessageHandler{})
	r.Run()
}
