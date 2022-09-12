package wxcommon

func (scc *SuiteCorpClient) ExternalContactGroupChatList(req *ExternalContactGroupChatListReq) (*ExternalContactGroupChatListResp, error) {
	resp := &ExternalContactGroupChatListResp{}
	if err := scc.PostRespWithToken("/externalcontact/groupchat/list?access_token=%s", req, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactGroupChatGet(chatId string) (*ExternalContactGroupChatGetResp, error) {
	resp := &ExternalContactGroupChatGetResp{}
	if err := scc.PostRespWithToken("/externalcontact/groupchat/get?access_token=%s", map[string]interface{}{
		"chat_id":   chatId,
		"need_name": 1,
	}, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) ExternalContactOpenGIdToChatId(openGId string) (*ExternalContactOpenGIdToChatIdResp, error) {
	resp := &ExternalContactOpenGIdToChatIdResp{}
	if err := scc.PostRespWithToken("/externalcontact/opengid_to_chatid?access_token=%s", map[string]interface{}{
		"opengid": openGId,
	}, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
