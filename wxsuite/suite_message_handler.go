package wxsuite

import (
	"bytes"
	"encoding/xml"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glutwins/workwx/internal/lowlevel/encryptor"
	"github.com/glutwins/workwx/internal/lowlevel/signature"
	"github.com/glutwins/workwx/wxcommon"
)

type SuiteEvent struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Event        string
	EventKey     string
	AgentID      int
}

type SuiteMessageHandler interface {
	OnCallbackEvent(*wxcommon.XmlRxEnvelope, *SuiteEvent)
}

type DummySuiteMessageHandler struct {
	Logger *log.Logger
}

func (h *DummySuiteMessageHandler) OnCallbackEvent(d *wxcommon.XmlRxEnvelope, ev *SuiteEvent) {
	if h.Logger != nil {
		h.Logger.Printf("%d To[%s][%d] From[%s]: msg=%s event=%s (%s)\n", ev.CreateTime, ev.ToUserName, ev.AgentID, ev.FromUserName, ev.MsgType, ev.Event, d.Encrypt)
	}
}

func NewMessageHandler(cfg *wxcommon.SuiteCallbackConfig, enc *encryptor.WorkwxEncryptor, h SuiteMessageHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req wxcommon.XmlRxEnvelope
		req.Query = ctx.Request.URL.Query()
		if err := ctx.BindXML(&req); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		if !signature.VerifyHTTPRequestSignature(cfg.Token, ctx.Request.URL, req.Encrypt) {
			ctx.Status(http.StatusBadRequest)
			return
		}

		payload, err := enc.Decrypt([]byte(req.Encrypt))
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}

		ev := &SuiteEvent{}
		if err = xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(ev); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		h.OnCallbackEvent(&req, ev)

		ctx.String(http.StatusOK, "success")
	}
}
