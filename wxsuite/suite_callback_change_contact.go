package wxsuite

import (
	"encoding/xml"

	"github.com/glutwins/workwx/wxcommon"
)

func onChangeContact(dec *xml.Decoder, req *wxcommon.XmlRxEnvelope, data *SuiteCallbackData, h SuiteCallbackHandler) error {
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

	switch data.ChangeType {
	case wxcommon.ChangeContactCreateUser, wxcommon.ChangeContactDeleteUser, wxcommon.ChangeContactUpdateUser:
		user := &wxcommon.XmlCallbackUser{}
		if err := dec.Decode(user); err != nil {
			return err
		}
		h.OnCallbackChangeContactUser(req, base, user)
	case wxcommon.ChangeContactCreateParty, wxcommon.ChangeContactUpdateParty, wxcommon.ChangeContactDeleteParty:
		depart := &wxcommon.XmlCallbackDepart{}
		if err := dec.Decode(depart); err != nil {
			return err
		}
		h.OnCallbackChangeContactDepart(req, base, depart)
	case wxcommon.ChangeContactUpdateTag:
		tag := &wxcommon.XmlCallbackTag{}
		if err := dec.Decode(tag); err != nil {
			return err
		}
		h.OnCallbackChangeContactTag(req, base, tag)
	default:
		h.OnCallbackChangeContactUnkown(req, data)
	}
	return nil
}
