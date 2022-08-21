package wxcommon

import "net/url"

type XmlRxEnvelope struct {
	ToUserName string     `xml:"ToUserName"`
	AgentID    string     `xml:"AgentID"`
	Encrypt    string     `xml:"Encrypt"`
	Query      url.Values `xml:"-"`
}

type XmlCallbackUser struct {
	UserID         string
	OpenUserID     string
	NewUserID      string
	Name           string
	Department     string
	MainDepartment int
	IsLeaderInDept string
	DirectLeader   string
	Mobile         string
	Position       string
	Gender         int
	Email          string
	BizMail        string
	Avatar         string
	Alias          string
	Telephone      string
	ExtAttr        []ExtAttr `xml:">Item,omitempty"`
}

type XmlCallbackDepart struct {
	Id       int
	Name     string
	ParentId string
	Order    int
}

type XmlCallbackTag struct {
	TagId         int
	AddUserItems  string
	DelUserItems  string
	AddPartyItems string
	DelPartyItems string
}

type BatchJob struct {
	JobId   string
	JobType string
	ErrCode int
	ErrMsg  string
}

type XmlCallbackJob struct {
	BatchJob BatchJob
}

type XmlCallbackExternalUser struct {
	UserID         string
	ExternalUserID string
	State          string
	WelcomeCode    string
	Source         string
	FailReason     string
}

type XmlCallbackExternalChat struct {
	ChatId       string
	UpdateDetail string
	JoinScene    int
	QuitScene    int
	MemChangeCnt int
}

type XmlCallbackExternalTag struct {
	Id      string
	TagType string
}
