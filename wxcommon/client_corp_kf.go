package wxcommon

func (scc *SuiteCorpClient) SyncKfMsg(token string, openKfId string, cursor string, limit int) (*SyncKfMsgResp, error) {
	resp := &SyncKfMsgResp{}
	if err := scc.PostRespWithToken("/kf/sync_msg?access_token=%s", SyncKfMsgReq{
		OpenKfId: openKfId,
		Token:    token,
		Cursor:   cursor,
		Limit:    limit,
	}, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (scc *SuiteCorpClient) GetKfServiceState(openKfId string, externalUserId string) (*GetServiceStateResp, error) {
	resp := &GetServiceStateResp{}
	if err := scc.PostRespWithToken("/kf/service_state/get?access_token=%s", map[string]interface{}{
		"open_kfid":       openKfId,
		"external_userid": externalUserId,
	}, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (scc *SuiteCorpClient) TransServiceState(openKfId string, externalUserId string, serviceUserId string, toState KfServiceState) (*TransServiceStateResp, error) {
	resp := &TransServiceStateResp{}
	if err := scc.PostRespWithToken("/kf/service_state/trans?access_token=%s", TransServiceStateReq{
		OpenKfId:       openKfId,
		ExternalUserId: externalUserId,
		ServiceState: ServiceState{
			ServicerUserId: serviceUserId,
			ServiceState:   toState,
		},
	}, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (scc *SuiteCorpClient) KfSendMsgOnEvent(msg KfMsgBody) (*KfSendMsgResp, error) {
	resp := &KfSendMsgResp{}
	if err := scc.PostRespWithToken("/kf/send_msg_on_event?access_token=%s", msg, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
func (scc *SuiteCorpClient) KfSendMsg(msg KfSendMsgRequest) (*KfSendMsgResp, error) {
	resp := &KfSendMsgResp{}
	if err := scc.PostRespWithToken("/kf/send_msg?access_token=%s", msg, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (scc *SuiteCorpClient) KfServiceList(openKfId string) (*KfServicerListResp, error) {
	resp := &KfServicerListResp{}
	if err := scc.GetRespWithToken("/kf/servicer/list?access_token=%s&open_kfid=%s", resp, openKfId); err != nil {
		return nil, err
	}
	return resp, nil
}
func (scc *SuiteCorpClient) KfServiceAdd(openKfId string, serviers []string, departments []string) (*KfEditServicerResp, error) {
	resp := &KfEditServicerResp{}
	if err := scc.PostRespWithToken("/kf/servicer/add?access_token=%s", map[string]interface{}{
		"open_kfid":          openKfId,
		"userid_list":        serviers,
		"department_id_list": departments,
	}, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (scc *SuiteCorpClient) KfServiceDel(openKfId string, serviers []string, departments []string) (*KfEditServicerResp, error) {
	resp := &KfEditServicerResp{}
	if err := scc.PostRespWithToken("/kf/servicer/del?access_token=%s", map[string]interface{}{
		"open_kfid":          openKfId,
		"userid_list":        serviers,
		"department_id_list": departments,
	}, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
