package wxsuite

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glutwins/workwx/internal/lowlevel/encryptor"
	"github.com/glutwins/workwx/internal/lowlevel/signature"
)

type SuiteConfig struct {
	SuiteId        string
	SuiteSecret    string
	Token          string
	EncodingAESKey string
}

func RegisterSuiteHandler(g *gin.RouterGroup, cfg *SuiteConfig, cmdHandler SuiteCallbackHandler, msgHandler SuiteMessageHandler) error {
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
	g.GET(fmt.Sprintf("/suite/%s/*action", cfg.SuiteId), func(ctx *gin.Context) {
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

	g.POST(fmt.Sprintf("/suite/%s/contact", cfg.SuiteId), NewCallbackHandler(cfg, encWithBody, cmdHandler))
	g.POST(fmt.Sprintf("/suite/%s/message", cfg.SuiteId), NewMessageHandler(cfg, encWithBody, msgHandler))
	return nil
}
