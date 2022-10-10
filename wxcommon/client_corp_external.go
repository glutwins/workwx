package wxcommon

func (scc *SuiteCorpClient) ExternalContactSendWelcomeMsg(req *ExternalContactSendWelcomeMsgReq) (*CommonResp, error) {
	resp := &CommonResp{}
	if err := scc.PostRespWithToken("/externalcontact/send_welcome_msg?access_token=%s", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactGet(externalUserId string, cursor string) (*ExternalContactGetResp, error) {
	resp := &ExternalContactGetResp{}
	if err := scc.GetRespWithToken("/externalcontact/get?access_token=%s&external_userid=%s&cursor=%s", resp, externalUserId, cursor); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactList(userId string) (*ExternalContactListResp, error) {
	resp := &ExternalContactListResp{}
	if err := scc.GetRespWithToken("/externalcontact/list?access_token=%s&userid=%s", resp, userId); err != nil {
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

func (scc *SuiteCorpClient) ExternalContactAddMsgTemplate(req *ExternalContactAddMsgTemplateReq) (*ExternalContactAddMsgTemplateResp, error) {
	resp := &ExternalContactAddMsgTemplateResp{}
	if err := scc.PostRespWithToken("/externalcontact/add_msg_template?access_token=%s", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (scc *SuiteCorpClient) UnionidToExternalUserid(req *UnionidToExternalUseridReq) (*UnionidToExternalUseridResp, error) {
	resp := &UnionidToExternalUseridResp{}
	if err := scc.PostRespWithToken("/idconvert/unionid_to_external_userid?access_token=%s", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (scc *SuiteCorpClient) ExternalUserIdToPendingId(req *ExternalUserIdToPendingIdReq) (*ExternalUserIdToPendingIdResp, error) {
	resp := &ExternalUserIdToPendingIdResp{}
	if err := scc.PostRespWithToken("/idconvert/batch/external_userid_to_pending_id?access_token=%s", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

//userid和partyid不可同时为空;
//此接口提供的数据以天为维度，查询的时间范围为[start_time,end_time]，即前后均为闭区间，支持的最大查询跨度为30天；
//用户最多可获取最近180天内的数据；
//如传入多个userid，则表示获取这些成员总体的联系客户数据
func (scc *SuiteCorpClient) ExternalContactGetBehaviorData(req *ExternalContactGetBehaviorDataReq) (*ExternalContactGetBehaviorDataResp, error) {
	resp := &ExternalContactGetBehaviorDataResp{}
	if err := scc.PostRespWithToken("/externalcontact/get_user_behavior_data?access_token=%s", req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
