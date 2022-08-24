package wxcommon

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
