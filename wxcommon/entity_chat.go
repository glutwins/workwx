package wxcommon

type GroupChatMemberInvitor struct {
	UserId string `json:"userid"`
}

type GroupChatMember struct {
	UserId        string                 `json:"userid"`
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
	GroupChat
}
