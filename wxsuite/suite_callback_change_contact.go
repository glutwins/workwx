package wxsuite

import (
	"encoding/xml"

	"github.com/glutwins/workwx/wxcommon"
)

func onChangeContact(dec *xml.Decoder, req *wxcommon.XmlRxEnvelope, data *SuiteCallbackData, h SuiteCallbackHandler) error {
	switch data.ChangeType {
	case wxcommon.ChangeContactCreateUser, wxcommon.ChangeContactDeleteUser, wxcommon.ChangeContactUpdateUser:
		user := &wxcommon.XmlCallbackUser{}
		if err := dec.Decode(user); err != nil {
			return err
		}
		h.OnCallbackChangeContactUser(req, &data.SuiteCallbackBase, user)
	case wxcommon.ChangeContactCreateParty, wxcommon.ChangeContactUpdateParty, wxcommon.ChangeContactDeleteParty:
		depart := &wxcommon.XmlCallbackDepart{}
		if err := dec.Decode(depart); err != nil {
			return err
		}
		h.OnCallbackChangeContactDepart(req, &data.SuiteCallbackBase, depart)
	case wxcommon.ChangeContactUpdateTag:
		tag := &wxcommon.XmlCallbackTag{}
		if err := dec.Decode(tag); err != nil {
			return err
		}
		h.OnCallbackChangeContactTag(req, &data.SuiteCallbackBase, tag)
	default:
		h.OnCallbackChangeContactUnkown(req, data)
	}
	return nil
}
