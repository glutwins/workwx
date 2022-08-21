package wxsuite

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/glutwins/workwx/internal/lowlevel/encryptor"
	"github.com/glutwins/workwx/internal/lowlevel/signature"
)

type SuiteCallbackHandler interface {
	OnCallbackSuiteTicket(*XmlRxEnvelope, *SuiteCallbackBase, string)
	OnCallbackCreateAuth(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth)
	OnCallbackChangeAuth(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth)
	OnCallbackCancelAuth(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth)
	OnCallbackResetPermanentCode(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth)
	OnCallbackShareAgentChange(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackShare)
	OnCallbackShareChainChange(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackShare)
	OnCallbackChangeContactUnkown(*XmlRxEnvelope, *SuiteCallbackData)
	OnCallbackChangeContactUser(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackUser)
	OnCallbackChangeContactDepart(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackDepart)
	OnCallbackChangeContactTag(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackTag)
	OnCallbackChangeExternalUser(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackExternalUser)
	OnCallbackChangeExternalChat(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackExternalChat)
	OnCallbackChangeExternalTag(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackExternalTag)
}

type DummySuiteCallbackHandler struct {
}

func (h *DummySuiteCallbackHandler) OnCallbackSuiteTicket(*XmlRxEnvelope, *SuiteCallbackBase, string) {
}
func (h *DummySuiteCallbackHandler) OnCallbackCreateAuth(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeAuth(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth) {
}
func (h *DummySuiteCallbackHandler) OnCallbackCancelAuth(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth) {
}
func (h *DummySuiteCallbackHandler) OnCallbackResetPermanentCode(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackAuth) {
}
func (h *DummySuiteCallbackHandler) OnCallbackShareAgentChange(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackShare) {
}
func (h *DummySuiteCallbackHandler) OnCallbackShareChainChange(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackShare) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeContactUnkown(*XmlRxEnvelope, *SuiteCallbackData) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeContactUser(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackUser) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeContactDepart(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackDepart) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeContactTag(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackTag) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeExternalUser(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackExternalUser) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeExternalChat(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackExternalChat) {
}
func (h *DummySuiteCallbackHandler) OnCallbackChangeExternalTag(*XmlRxEnvelope, *SuiteCallbackBase, *SuiteCallbackExternalTag) {
}

const (
	SuiteCallbackTypeSuiteTicket           string = "suite_ticket"
	SuiteCallbackTypeCreateAuth            string = "create_auth"
	SuiteCallbackTypeChangeAuth            string = "change_auth"
	SuiteCallbackTypeCancelAuth            string = "cancel_auth"
	SuiteCallbackTypeResetPermanentCode    string = "reset_permanent_code"
	SuiteCallbackTypeShareAgentChange      string = "share_agent_change"
	SuiteCallbackTypeShareChainChange      string = "share_chain_change"
	SuiteCallbackTypeChangeContact         string = "change_contact"
	SuiteCallbackTypeChangeExternalContact string = "change_external_contact"
	SuiteCallbackTypeChangeExternalChat    string = "change_external_chat"
	SuiteCallbackTypeChangeExternalTag     string = "change_external_tag"
)

type XmlRxEnvelope struct {
	ToUserName string     `xml:"ToUserName"`
	AgentID    string     `xml:"AgentID"`
	Encrypt    string     `xml:"Encrypt"`
	Query      url.Values `xml:"-"`
}

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
		var req XmlRxEnvelope
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
		case SuiteCallbackTypeChangeContact:
			if err := onChangeContact(xml.NewDecoder(bytes.NewBuffer(payload.Msg)), &req, data, h); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
		case SuiteCallbackTypeChangeExternalContact:
			user := &SuiteCallbackExternalUser{}
			if err := xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(user); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			h.OnCallbackChangeExternalUser(&req, &data.SuiteCallbackBase, user)
		case SuiteCallbackTypeChangeExternalChat:
			chat := &SuiteCallbackExternalChat{}
			if err := xml.NewDecoder(bytes.NewBuffer(payload.Msg)).Decode(chat); err != nil {
				ctx.Status(http.StatusBadRequest)
				return
			}
			h.OnCallbackChangeExternalChat(&req, &data.SuiteCallbackBase, chat)
		case SuiteCallbackTypeChangeExternalTag:
			tag := &SuiteCallbackExternalTag{}
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
