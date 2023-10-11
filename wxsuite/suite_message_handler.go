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

const (
	EventSubscribe   = "subscribe"       //关注
	EventUnSubscribe = "unsubscribe"     //取消关注
	EventKf          = "kf_msg_or_event" //客服消息事件
)

type SuiteEventBase struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Event        string
	EventKey     string
	AgentID      int
}

type SuiteKfEvent struct {
	SuiteEventBase
	Token    string
	OpenKfId string
}

type SuiteMessageHandler interface {
	OnCallbackEvent(*wxcommon.XmlRxEnvelope, *SuiteEventBase)
	OnCallbackKfEvent(*wxcommon.XmlRxEnvelope, *SuiteKfEvent)
}

type DummySuiteMessageHandler struct {
	Logger *log.Logger
}

func (h *DummySuiteMessageHandler) OnCallbackEvent(d *wxcommon.XmlRxEnvelope, ev *SuiteEventBase) {
	if h.Logger != nil {
		h.Logger.Printf("%d To[%s][%d] From[%s]: msg=%s event=%s (%s)\n", ev.CreateTime, ev.ToUserName, ev.AgentID, ev.FromUserName, ev.MsgType, ev.Event, d.Encrypt)
	}
}
func (h *DummySuiteMessageHandler) OnCallbackKfEvent(d *wxcommon.XmlRxEnvelope, ev *SuiteKfEvent) {
}

func decode[T any](msg []byte, t *T) *T {
	xml.NewDecoder(bytes.NewBuffer(msg)).Decode(t)
	return t
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

		ev := &SuiteEventBase{}
		if err = xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(ev); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		switch ev.Event {
		case EventKf:
			h.OnCallbackKfEvent(&req, decode[SuiteKfEvent](payload.Msg, &SuiteKfEvent{}))
		default:
			h.OnCallbackEvent(&req, ev)
		}

		ctx.String(http.StatusOK, "success")
	}
}
