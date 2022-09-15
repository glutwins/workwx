package wxsuite

import (
	"bytes"
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glutwins/workwx/internal/lowlevel/encryptor"
	"github.com/glutwins/workwx/internal/lowlevel/signature"
	"github.com/glutwins/workwx/store"
	"github.com/glutwins/workwx/wxcommon"
	"github.com/rs/zerolog"
)

type SuiteCallbackHandler interface {
	OnCallbackSuiteTicket(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, string)
	OnCallbackCreateAuth(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth)
	OnCallbackChangeAuth(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth)
	OnCallbackCancelAuth(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth)
	OnCallbackResetPermanentCode(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth)
	OnCallbackShareAgentChange(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackShare)
	OnCallbackShareChainChange(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackShare)
	OnCallbackChangeContactUnkown(*wxcommon.XmlRxEnvelope, *SuiteCallbackData)
	OnCallbackChangeContactUser(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackUser)
	OnCallbackChangeContactDepart(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackDepart)
	OnCallbackChangeContactTag(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackTag)
	OnCallbackChangeExternalUser(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackExternalUser)
	OnCallbackChangeExternalChat(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackExternalChat)
	OnCallbackChangeExternalTag(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackExternalTag)
}

type DummySuiteCallbackHandler struct {
	TokenCache store.TokenCache
}

func (h *DummySuiteCallbackHandler) OnCallbackSuiteTicket(raw *wxcommon.XmlRxEnvelope, base *SuiteCallbackBase, ticket string) {
	h.TokenCache.SetSuiteTicket(base.SuiteId, ticket)
}
func (h *DummySuiteCallbackHandler) OnCallbackCreateAuth(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeAuth(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth) {
}
func (h *DummySuiteCallbackHandler) OnCallbackCancelAuth(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth) {
}
func (h *DummySuiteCallbackHandler) OnCallbackResetPermanentCode(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth) {
}
func (h *DummySuiteCallbackHandler) OnCallbackShareAgentChange(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackShare) {
}
func (h *DummySuiteCallbackHandler) OnCallbackShareChainChange(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackShare) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeContactUnkown(*wxcommon.XmlRxEnvelope, *SuiteCallbackData) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeContactUser(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackUser) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeContactDepart(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackDepart) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeContactTag(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackTag) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeExternalUser(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackExternalUser) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeExternalChat(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackExternalChat) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeExternalTag(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, *wxcommon.XmlCallbackExternalTag) {
}

const (
	SuiteCallbackTypeSuiteTicket        string = "suite_ticket"
	SuiteCallbackTypeCreateAuth         string = "create_auth"
	SuiteCallbackTypeChangeAuth         string = "change_auth"
	SuiteCallbackTypeCancelAuth         string = "cancel_auth"
	SuiteCallbackTypeResetPermanentCode string = "reset_permanent_code"
	SuiteCallbackTypeShareAgentChange   string = "share_agent_change"
	SuiteCallbackTypeShareChainChange   string = "share_chain_change"
)

type SuiteCallbackAuth struct {
	AuthCode string // create_auth, reset_permanent_code
	State    string // create_auth, change_auth
}

type SuiteCallbackShare struct {
	AppId   int
	CorpId  string
	AgentId int
}

type SuiteCallbackBase struct {
	SuiteId    string
	InfoType   string
	TimeStamp  int64
	AuthCorpId string
	ChangeType string
}

type SuiteCallbackData struct {
	wxcommon.CallbackBase
	SuiteId     string
	InfoType    string
	TimeStamp   int64
	AuthCorpId  string
	SuiteTicket string
	SuiteCallbackShare
	SuiteCallbackAuth
}

func NewCallbackHandler(cfg *wxcommon.SuiteCallbackConfig, enc *encryptor.WorkwxEncryptor, h SuiteCallbackHandler) gin.HandlerFunc {
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

		data := &SuiteCallbackData{}
		if err = xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(data); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		base := &SuiteCallbackBase{}
		base.SuiteId = data.SuiteId
		base.InfoType = data.InfoType
		base.TimeStamp = data.TimeStamp
		base.AuthCorpId = data.AuthCorpId
		base.ChangeType = data.ChangeType

		if data.InfoType == "" && data.Event != "" {
			base.InfoType = data.Event
			base.AuthCorpId = data.ToUserName
			base.TimeStamp = data.CreateTime
			data.InfoType = data.Event
			data.AuthCorpId = data.ToUserName
			data.TimeStamp = data.CreateTime
		}

		ev = ev.Str("infoType", data.InfoType)

		switch base.InfoType {
		case SuiteCallbackTypeSuiteTicket:
			h.OnCallbackSuiteTicket(&req, base, data.SuiteTicket)
		case SuiteCallbackTypeCreateAuth:
			h.OnCallbackCreateAuth(&req, base, &data.SuiteCallbackAuth)
		case SuiteCallbackTypeChangeAuth:
			h.OnCallbackChangeAuth(&req, base, &data.SuiteCallbackAuth)
		case SuiteCallbackTypeCancelAuth:
			h.OnCallbackCancelAuth(&req, base, &data.SuiteCallbackAuth)
		case SuiteCallbackTypeResetPermanentCode:
			h.OnCallbackResetPermanentCode(&req, base, &data.SuiteCallbackAuth)
		case SuiteCallbackTypeShareAgentChange:
			h.OnCallbackShareAgentChange(&req, base, &data.SuiteCallbackShare)
		case SuiteCallbackTypeShareChainChange:
			h.OnCallbackShareChainChange(&req, base, &data.SuiteCallbackShare)
		case wxcommon.CallbackTypeChangeContact:
			if err := onChangeContact(xml.NewDecoder(bytes.NewBuffer(payload.Msg)), &req, data, h); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
		case wxcommon.CallbackTypeChangeExternalContact:
			user := &wxcommon.XmlCallbackExternalUser{}
			if err := xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(user); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			h.OnCallbackChangeExternalUser(&req, base, user)
		case wxcommon.CallbackTypeChangeExternalChat:
			chat := &wxcommon.XmlCallbackExternalChat{}
			if err := xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(chat); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			h.OnCallbackChangeExternalChat(&req, base, chat)
		case wxcommon.CallbackTypeChangeExternalTag:
			tag := &wxcommon.XmlCallbackExternalTag{}
			if err := xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(tag); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			h.OnCallbackChangeExternalTag(&req, base, tag)
		default:
			h.OnCallbackChangeContactUnkown(&req, data)
		}

		ctx.String(http.StatusOK, "success")
	}
}
