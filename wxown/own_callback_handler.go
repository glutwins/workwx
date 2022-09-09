package wxown

import (
	"bytes"
	"crypto/rand"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glutwins/workwx/internal/lowlevel/encryptor"
	"github.com/glutwins/workwx/internal/lowlevel/signature"
	"github.com/glutwins/workwx/wxcommon"
	"github.com/rs/zerolog"
)

type OwnConfig struct {
	AgentId     int
	AgentSecret string
}

type OwnCallbackBase struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Event        string
	ChangeType   string
}

type OwnCallbackHandler interface {
	OnCallbackChangeBatchJobResult(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.BatchJob)
	OnCallbackChangeContactUnkown(*wxcommon.XmlRxEnvelope, *OwnCallbackBase)
	OnCallbackChangeContactUser(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackUser)
	OnCallbackChangeContactDepart(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackDepart)
	OnCallbackChangeContactTag(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackTag)
	OnCallbackChangeExternalUser(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackExternalUser)
	OnCallbackChangeExternalChat(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackExternalChat)
	OnCallbackChangeExternalTag(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackExternalTag)
}

type DummyOwnCallbackHandler struct {
}

func (h *DummyOwnCallbackHandler) OnCallbackChangeBatchJobResult(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.BatchJob) {
}
func (h *DummyOwnCallbackHandler) OnCallbackChangeContactUnkown(*wxcommon.XmlRxEnvelope, *OwnCallbackBase) {
}
func (h *DummyOwnCallbackHandler) OnCallbackChangeContactUser(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackUser) {
}
func (h *DummyOwnCallbackHandler) OnCallbackChangeContactDepart(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackDepart) {
}
func (h *DummyOwnCallbackHandler) OnCallbackChangeContactTag(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackTag) {
}
func (h *DummyOwnCallbackHandler) OnCallbackChangeExternalUser(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackExternalUser) {
}
func (h *DummyOwnCallbackHandler) OnCallbackChangeExternalChat(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackExternalChat) {
}
func (h *DummyOwnCallbackHandler) OnCallbackChangeExternalTag(*wxcommon.XmlRxEnvelope, *OwnCallbackBase, *wxcommon.XmlCallbackExternalTag) {
}

func NewCallbackHandler(cfg *wxcommon.SuiteCallbackConfig, enc *encryptor.WorkwxEncryptor, h OwnCallbackHandler) gin.HandlerFunc {
	logger := zerolog.New(cfg.LoggerWriter).Level(zerolog.InfoLevel)

	return func(ctx *gin.Context) {
		ev := logger.Info().Str("token", cfg.Token).Str("aeskey", cfg.EncodingAESKey).Str("url", ctx.Request.URL.String())
		defer func() {
			ev.Msg("wxowncallback")
		}()
		var req wxcommon.XmlRxEnvelope
		req.Query = ctx.Request.URL.Query()
		err := ctx.BindXML(&req)
		ev = ev.Err(err)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		ev = ev.Str("encrypt", req.Encrypt)
		if !signature.VerifyHTTPRequestSignature(cfg.Token, ctx.Request.URL, req.Encrypt) {
			ctx.Status(http.StatusBadRequest)
			return
		}

		payload, err := enc.Decrypt([]byte(req.Encrypt))
		ev = ev.Err(err)
		if err != nil {
			ctx.String(http.StatusBadRequest, err.Error())
			return
		}
		ev = ev.Str("decrypt", string(payload.Msg))

		data := &OwnCallbackBase{}
		if err = xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(data); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		ev = ev.Str("event", data.Event)

		switch data.Event {
		case wxcommon.CallbackTypeChangeContact:
			if err := onChangeContact(xml.NewDecoder(bytes.NewBuffer(payload.Msg)), &req, data, h); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
		case wxcommon.CallbackTypeBatchJobResult:
			job := &wxcommon.XmlCallbackJob{}
			if err := xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(job); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			h.OnCallbackChangeBatchJobResult(&req, data, &job.BatchJob)
		case wxcommon.CallbackTypeChangeExternalContact:
			user := &wxcommon.XmlCallbackExternalUser{}
			if err := xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(user); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			h.OnCallbackChangeExternalUser(&req, data, user)
		case wxcommon.CallbackTypeChangeExternalChat:
			chat := &wxcommon.XmlCallbackExternalChat{}
			if err := xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(chat); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			h.OnCallbackChangeExternalChat(&req, data, chat)
		case wxcommon.CallbackTypeChangeExternalTag:
			tag := &wxcommon.XmlCallbackExternalTag{}
			if err := xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(tag); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			h.OnCallbackChangeExternalTag(&req, data, tag)
		default:
			h.OnCallbackChangeContactUnkown(&req, data)
		}

		ctx.String(http.StatusOK, "success")
	}
}

func RegisterOwnHandler(g *gin.RouterGroup, cfg *wxcommon.SuiteCallbackConfig, cmdHandler OwnCallbackHandler) error {
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
	return nil
}
