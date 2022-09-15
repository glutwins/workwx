package wxown

import (
	"encoding/xml"

	"github.com/glutwins/workwx/wxcommon"
)

func onChangeContact(dec *xml.Decoder, req *wxcommon.XmlRxEnvelope, data *wxcommon.CallbackBase, h OwnCallbackHandler) error {
	switch data.ChangeType {
	case wxcommon.ChangeContactCreateUser, wxcommon.ChangeContactDeleteUser, wxcommon.ChangeContactUpdateUser:
		user := &wxcommon.XmlCallbackUser{}
		if err := dec.Decode(user); err != nil {
			return err
		}
		h.OnCallbackChangeContactUser(req, data, user)
	case wxcommon.ChangeContactCreateParty, wxcommon.ChangeContactUpdateParty, wxcommon.ChangeContactDeleteParty:
		depart := &wxcommon.XmlCallbackDepart{}
		if err := dec.Decode(depart); err != nil {
			return err
		}
		h.OnCallbackChangeContactDepart(req, data, depart)
	case wxcommon.ChangeContactUpdateTag:
		tag := &wxcommon.XmlCallbackTag{}
		if err := dec.Decode(tag); err != nil {
			return err
		}
		h.OnCallbackChangeContactTag(req, data, tag)
	default:
		h.OnCallbackChangeContactUnkown(req, data)
	}
	return nil
}
