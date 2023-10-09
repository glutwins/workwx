package wxcommon

type SyncKfMsgReq struct {
	Cursor   string `json:"cursor"`
	Token    string `json:"token"`
	OpenKfId string `json:"open_kfid"`
	Limit    int    `json:"limit"`
}

type SyncKfMsgResp struct {
	CommonResp
	NextCursor string  `json:"next_cursor"`
	HasMore    int     `json:"has_more"`
	MsgList    []KfMsg `json:"msg_list"`
}

type KfMsg struct {
	MsgId          string `json:"msgid"`
	OpenKfId       string `json:"open_kfid"`
	ExternalUserId string `json:"external_userid"`
	SendTime       int    `json:"send_time"`
	Origin         int    `json:"origin"`
	ServicerUserId string `json:"servicer_userid"`
	MsgType        string `json:"msgtype"`
}

type TransServiceStateReq struct {
	OpenKfId       string         `json:"open_kfid"`
	ExternalUserId string         `json:"external_userid"`
	ServiceState   KfServiceState `json:"service_state"`
	ServiceUserId  string         `json:"service_userid"`
}

type TransServiceStateResp struct {
	CommonResp
	MsgCode string `json:"msg_code"`
}
