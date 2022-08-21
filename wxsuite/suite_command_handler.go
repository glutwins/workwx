package wxsuite

import (
	"bytes"
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glutwins/workwx/internal/lowlevel/encryptor"
	"github.com/glutwins/workwx/internal/lowlevel/signature"
	"github.com/glutwins/workwx/wxcommon"
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
}

func (h *DummySuiteCallbackHandler) OnCallbackSuiteTicket(*wxcommon.XmlRxEnvelope, *SuiteCallbackBase, string) {
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
	SuiteCallbackBase
	SuiteTicket string
	SuiteCallbackShare
	SuiteCallbackAuth
}

func NewCallbackHandler(cfg *SuiteConfig, enc *encryptor.WorkwxEncryptor, h SuiteCallbackHandler) gin.HandlerFunc {
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

		data := &SuiteCallbackData{}
		if err = xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(data); err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		switch data.InfoType {
		case SuiteCallbackTypeSuiteTicket:
			h.OnCallbackSuiteTicket(&req, &data.SuiteCallbackBase, data.SuiteTicket)
		case SuiteCallbackTypeCreateAuth:
			h.OnCallbackCreateAuth(&req, &data.SuiteCallbackBase, &data.SuiteCallbackAuth)
		case SuiteCallbackTypeChangeAuth:
			h.OnCallbackChangeAuth(&req, &data.SuiteCallbackBase, &data.SuiteCallbackAuth)
		case SuiteCallbackTypeCancelAuth:
			h.OnCallbackCancelAuth(&req, &data.SuiteCallbackBase, &data.SuiteCallbackAuth)
		case SuiteCallbackTypeResetPermanentCode:
			h.OnCallbackResetPermanentCode(&req, &data.SuiteCallbackBase, &data.SuiteCallbackAuth)
		case SuiteCallbackTypeShareAgentChange:
			h.OnCallbackShareAgentChange(&req, &data.SuiteCallbackBase, &data.SuiteCallbackShare)
		case SuiteCallbackTypeShareChainChange:
			h.OnCallbackShareChainChange(&req, &data.SuiteCallbackBase, &data.SuiteCallbackShare)
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
			h.OnCallbackChangeExternalUser(&req, &data.SuiteCallbackBase, user)
		case wxcommon.CallbackTypeChangeExternalChat:
			chat := &wxcommon.XmlCallbackExternalChat{}
			if err := xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(chat); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			h.OnCallbackChangeExternalChat(&req, &data.SuiteCallbackBase, chat)
		case wxcommon.CallbackTypeChangeExternalTag:
			tag := &wxcommon.XmlCallbackExternalTag{}
			if err := xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(tag); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			h.OnCallbackChangeExternalTag(&req, &data.SuiteCallbackBase, tag)
		default:
			h.OnCallbackChangeContactUnkown(&req, data)
		}

		ctx.String(http.StatusOK, "success")
	}
}
