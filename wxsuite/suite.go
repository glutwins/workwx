package wxsuite

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glutwins/workwx/internal/lowlevel/encryptor"
	"github.com/glutwins/workwx/internal/lowlevel/signature"
	"github.com/glutwins/workwx/wxcommon"
)

type SuiteConfig struct {
	SuiteId     string
	SuiteSecret string
}

func RegisterSuiteHandler(g *gin.RouterGroup, cfg *wxcommon.SuiteCallbackConfig, cmdHandler SuiteCallbackHandler, msgHandler SuiteMessageHandler) error {
	enc, err := encryptor.NewWorkwxEncryptor(cfg.EncodingAESKey)
	if err != nil {
		return err
	}

	encWithBody, err := encryptor.NewWorkwxEncryptor(
		cfg.EncodingAESKey,
		encryptor.WithEntropySource(rand.Reader),
	)
	if err != nil {
		return err
	}

	// 回调校验测试
	g.GET(fmt.Sprintf("/suite/%s/*action", cfg.SuiteKey), func(ctx *gin.Context) {
		if !signature.VerifyHTTPRequestSignature(cfg.Token, ctx.Request.URL, "") {
			ctx.Status(http.StatusBadRequest)
			return
		}

		payload, err := enc.Decrypt([]byte(ctx.Query("echostr")))
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		ctx.String(http.StatusOK, string(payload.Msg))
	})

	g.POST(fmt.Sprintf("/suite/%s/contact", cfg.SuiteKey), NewCallbackHandler(cfg, encWithBody, cmdHandler))
	g.POST(fmt.Sprintf("/suite/%s/message", cfg.SuiteKey), NewMessageHandler(cfg, encWithBody, msgHandler))
	return nil
}
