package wxcommon

func (scc *SuiteCorpClient) ExternalContactSendWelcomeMsg(req *ExternalContactSendWelcomeMsgReq) (*CommonResp, error) {
	resp := &CommonResp{}
	if err := scc.PostRespWithToken("/externalcontact/groupchat/get?access_token=%s", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactGet(externalUserId string, cursor string) (*ExternalContactGetResp, error) {
	resp := &ExternalContactGetResp{}
	if err := scc.GetRespWithToken("/externalcontact/list?access_token=%s&userid=%s", resp, externalUserId, cursor); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactList(userId string) (*ExternalContactListResp, error) {
	resp := &ExternalContactListResp{}
	if err := scc.GetRespWithToken("/externalcontact/list?access_token=%s&external_userid=%s&cursor=%s", resp, userId); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactBatchGetByUser(req *ExternalContactBatchGetByUserReq) (*ExternalContactBatchGetByUserResp, error) {
	resp := &ExternalContactBatchGetByUserResp{}
	if err := scc.PostRespWithToken("/externalcontact/batch/get_by_user?access_token=%s", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactRemark(req *ExternalContactRemarkReq) (*CommonResp, error) {
	resp := &CommonResp{}
	if err := scc.PostRespWithToken("/externalcontact/remark?access_token=%s", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactGetCorpTagList(req *ExternalContactGetCorpTagListReq) (*ExternalContactGetCorpTagListResp, error) {
	resp := &ExternalContactGetCorpTagListResp{}
	if err := scc.PostRespWithToken("/externalcontact/get_corp_tag_list?access_token=%s", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactAddCorpTag(req *ExternalContactAddCorpTagReq) (*ExternalContactAddCorpTagResp, error) {
	resp := &ExternalContactAddCorpTagResp{}
	if err := scc.PostRespWithToken("/externalcontact/add_corp_tag?access_token=%s", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactEditCorpTag(req *ExternalContactEditCorpTagReq) (*CommonResp, error) {
	resp := &CommonResp{}
	if err := scc.PostRespWithToken("/externalcontact/edit_corp_tag?access_token=%s", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactDelCorpTag(req *ExternalContactGetCorpTagListReq) (*CommonResp, error) {
	resp := &CommonResp{}
	if err := scc.PostRespWithToken("/externalcontact/del_corp_tag?access_token=%s", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactMarkTag(req *ExternalContactMarkTagReq) (*CommonResp, error) {
	resp := &CommonResp{}
	if err := scc.PostRespWithToken("/externalcontact/mark_tag?access_token=%s", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
