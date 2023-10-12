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
	SuiteCallbackTypeEventSubscribe          = "subscribe"       //关注
	SuiteCallbackTypeEventUnSubscribe        = "unsubscribe"     //取消关注
	SuiteCallbackTypeKfMsgOrEvent     string = "kf_msg_or_event" //客服消息事件
)

type SuiteMessageHandler interface {
	OnCallbackEvent(*wxcommon.XmlRxEnvelope, *wxcommon.SuiteEventBase) error
	OnKfMsgOrEvent(*wxcommon.XmlRxEnvelope, *wxcommon.SuiteKfEvent) error
}

type DummySuiteMessageHandler struct {
	Logger *log.Logger
}

func (h *DummySuiteMessageHandler) OnCallbackEvent(d *wxcommon.XmlRxEnvelope, ev *wxcommon.SuiteEventBase) {
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

		ev := &wxcommon.SuiteEventBase{}
		if err = xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(ev); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}
		if err = onEventsOrMessageCallback(xml.NewDecoder(bytes.NewBuffer(payload.Msg)), &req, ev, h); err != nil {
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.String(http.StatusOK, "success")
	}
}

func onEventsOrMessageCallback(dec *xml.Decoder, req *wxcommon.XmlRxEnvelope, event *wxcommon.SuiteEventBase, h SuiteMessageHandler) (err error) {
	switch event.Event {
	case SuiteCallbackTypeKfMsgOrEvent:
		kfEvent := &wxcommon.SuiteKfEvent{}
		if err := dec.Decode(kfEvent); err != nil {
			return err
		}
		err = h.OnKfMsgOrEvent(req, kfEvent)
	default:
		err = h.OnCallbackEvent(req, event)
	}
	return
}
