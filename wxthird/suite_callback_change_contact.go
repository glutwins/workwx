package wxthird

import (
	"encoding/xml"

	"github.com/glutwins/workwx/wxcommon"
)

type SuiteCallbackExtAttrText struct {
	Value string
}

type SuiteCallbackExtAttrWeb struct {
	Title string
	Url   string
}

type SuiteCallbackExtAttrItem struct {
	Name string
	Type int
	Text []*SuiteCallbackExtAttrText
	Web  []*SuiteCallbackExtAttrWeb
}

type SuiteCallbackUser struct {
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
	ExtAttr        []*SuiteCallbackExtAttrItem
}

type SuiteCallbackDepart struct {
	Id       int
	Name     string
	ParentId string
	Order    int
}

type SuiteCallbackTag struct {
	TagId         int
	AddUserItems  string
	DelUserItems  string
	AddPartyItems string
	DelPartyItems string
}

func onChangeContact(dec *xml.Decoder, req *XmlRxEnvelope, data *SuiteCallbackData, h SuiteCallbackHandler) error {
	switch data.ChangeType {
	case wxcommon.ChangeContactCreateUser, wxcommon.ChangeContactDeleteUser, wxcommon.ChangeContactUpdateUser:
		user := &SuiteCallbackUser{}
		if err := dec.Decode(user); err != nil {
			return err
		}
		h.OnCallbackChangeContactUser(req, &data.SuiteCallbackBase, user)
	case wxcommon.ChangeContactCreateParty, wxcommon.ChangeContactUpdateParty, wxcommon.ChangeContactDeleteParty:
		depart := &SuiteCallbackDepart{}
		if err := dec.Decode(depart); err != nil {
			return err
		}
		h.OnCallbackChangeContactDepart(req, &data.SuiteCallbackBase, depart)
	case wxcommon.ChangeContactUpdateTag:
		tag := &SuiteCallbackTag{}
		if err := dec.Decode(tag); err != nil {
			return err
		}
		h.OnCallbackChangeContactTag(req, &data.SuiteCallbackBase, tag)
	default:
		h.OnCallbackChangeContactUnkown(req, data)
	}
	return nil
}
