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

func (scc *SuiteCorpClient) TransServiceState(openKfId string, externalUserId string, serviceUserId string, toState KfServiceState) (*TransServiceStateResp, error) {
	resp := &TransServiceStateResp{}
	if err := scc.PostRespWithToken("/kf/sync_msg?access_token=%s", TransServiceStateReq{
		OpenKfId:       openKfId,
		ExternalUserId: externalUserId,
		ServiceUserId:  serviceUserId,
		ServiceState:   toState,
	}, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
