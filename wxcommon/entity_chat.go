package wxcommon

type GroupChatMemberInvitor struct {
	UserId string `json:"userid"`
}

type GroupChatMember struct {
	UserId        string                 `json:"userid"`
	UnionId       string                 `json:"unionid"` // 仅限自建应用
	Type          int                    `json:"type"`
	JoinTime      int64                  `json:"join_time"`
	JoinScene     int                    `json:"join_scene"`
	Invitor       GroupChatMemberInvitor `json:"invitor"`
	GroupNickName string                 `json:"group_nickname"`
	Name          string                 `json:"name"`
}

type GroupChat struct {
	ChatId     string                    `json:"chat_id"`
	Name       string                    `json:"name"`
	Owner      string                    `json:"owner"`
	CreateTime int64                     `json:"create_time"`
	Notice     string                    `json:"notice"`
	MemberList []*GroupChatMember        `json:"member_list"`
	AdminList  []*GroupChatMemberInvitor `json:"admin_list"`
}

type ExternalContactGroupChatGetResp struct {
	CommonResp
	GroupChat GroupChat `json:"group_chat"`
}

type OwnFilter struct {
	UserIdList []string `json:"userid_list"`
}

type ExternalContactGroupChatListReq struct {
	StatusFilter int       `json:"status_filter"`
	OwnFilter    OwnFilter `json:"owner_filter"`
	Limit        int       `json:"limit"`
	Cursor       string    `json:"cursor"`
}

type GroupChatListItem struct {
	ChatId string `json:"chat_id"`
	Status int    `json:"status"`
}

type ExternalContactGroupChatListResp struct {
	CommonResp
	GroupChatList []*GroupChatListItem `json:"group_chat_list"`
	NextCursor    string               `json:"next_cursor"`
}

type ExternalContactOpenGIdToChatIdResp struct {
	CommonResp
	ChatId string `json:"chat_id"`
}

type ExternalContactGroupChatAddJoinWayReq struct {
	Scene          int      `json:"scene"`
	Remark         string   `json:"remark"`
	AutoCreateRoom int      `json:"auto_create_room"`
	RoomBaseName   string   `json:"room_base_name"`
	RoomBaseID     int      `json:"room_base_id"`
	ChatIDList     []string `json:"chat_id_list"`
	State          string   `json:"state"`
}

type ExternalContactGroupChatAddJoinWayResp struct {
	CommonResp
	ConfigID string `json:"config_id"`
}

type ExternalContactGroupChatGetJoinWayReq struct {
	ConfigID string `json:"config_id"`
}

type ExternalContactGroupChatGetJoinWayResp struct {
	CommonResp
	JoinWay JoinWayInfo `json:"join_way"`
}

type JoinWayInfo struct {
	ConfigID       string   `json:"config_id"`
	Scene          int      `json:"scene"`
	Remark         string   `json:"remark"`
	AutoCreateRoom int      `json:"auto_create_room"`
	RoomBaseName   string   `json:"room_base_name"`
	RoomBaseID     int      `json:"room_base_id"`
	ChatIDList     []string `json:"chat_id_list"`
	QrCode         string   `json:"qr_code"`
	State          string   `json:"state"`
}

type ExternalContactGroupChatUpdateJoinWayReq struct {
	ConfigID       string   `json:"config_id"`
	Scene          int      `json:"scene"`
	Remark         string   `json:"remark"`
	AutoCreateRoom int      `json:"auto_create_room"`
	RoomBaseName   string   `json:"room_base_name"`
	RoomBaseID     int      `json:"room_base_id"`
	ChatIDList     []string `json:"chat_id_list"`
	State          string   `json:"state"`
}

type ExternalContactGroupChatDelJoinWayReq struct {
	ConfigID string `json:"config_id"`
}
