package wxcommon

func (scc *SuiteCorpClient) ExternalContactSendWelcomeMsg(req *ExternalContactSendWelcomeMsgReq) (*CommonResp, error) {
	resp := &CommonResp{}
	if err := scc.PostRespWithToken("/externalcontact/groupchat/get?access_token=%s", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
