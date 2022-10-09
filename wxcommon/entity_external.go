package wxcommon

type ExternalContactSendWelcomeMsgReq struct {
	WelcomeCode string        `json:"welcome_code"`
	Text        *Text         `json:"text,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

type ExternalContactAddMsgTemplateReq struct {
	ChatType       string        `json:"chat_type"`
	ExternalUserId []string      `json:"external_userid"`
	Sender         string        `json:"sender"`
	Text           *Text         `json:"text,omitempty"`
	Attachments    []*Attachment `json:"attachments,omitempty"`
}

type ExternalContactAddMsgTemplateResp struct {
	CommonResp
	FailList []string `json:"fail_list"`
	MsgId    string   `json:"msgid"`
}

type ExternalContactGetGroupMsgListV2Req struct {
	ChatType   string `json:"chat_type"`
	StartTime  int64  `json:"start_time"`
	EndTime    int64  `json:"end_time"`
	Creator    string `json:"creator"`
	FilterType int    `json:"filter_type"`
	Limit      int    `json:"limit"`
	Cursor     string `json:"cursor"`
}

type ExternalContactGetGroupMsg struct {
	MsgId       string        `json:"msgid"`
	Creator     string        `json:"creator"`
	CreateTime  string        `json:"create_time"`
	CreateType  int           `json:"create_type"`
	Text        *Text         `json:"text,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

type ExternalContactGetGroupMsgListV2Resp struct {
	CommonResp
	NextCursor   string                        `json:"next_cursor"`
	GroupMsgList []*ExternalContactGetGroupMsg `json:"group_msg_list"`
}

type ExternalContactGetGroupMsgTaskReq struct {
	MsgId  string `json:"msgid"`
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type ExternalContactGetGroupMsgTaskItem struct {
	UserId   string `json:"userid"`
	Status   int    `json:"status"`
	SendTime int    `json:"send_time"`
}

type ExternalContactGetGroupMsgTaskResp struct {
	CommonResp
	NextCursor string                                `json:"next_cursor"`
	TaskList   []*ExternalContactGetGroupMsgTaskItem `json:"task_list"`
}

type ExternalContactGetGroupMsgSendResultReq struct {
	ExternalContactGetGroupMsgTaskReq
	UserId string `json:"userid"`
}

type ExternalContactGetGroupMsgSend struct {
	ExternalUserId string `json:"external_userid"`
	ChatId         string `json:"chat_id"`
	UserId         string `json:"userid"`
	Status         int    `json:"status"`
	SendTime       int64  `json:"send_time"`
}

type ExternalContactGetGroupMsgSendResultResp struct {
	CommonResp
	NextCursor string                            `json:"next_cursor"`
	SendList   []*ExternalContactGetGroupMsgSend `json:"send_list"`
}

type ExternalProfile struct {
	ExternalAttr []*ExtAttr `json:"external_attr"`
}

type FollowUserTag struct {
	GroupName string `json:"group_name"`
	TagName   string `json:"tag_name"`
	TagId     string `json:"tag_id"`
	Type      int    `json:"type"`
}

type FollowUser struct {
	UserId         string           `json:"userid"`
	Remark         string           `json:"remark"`
	Description    string           `json:"description"`
	CreateTime     int64            `json:"createtime"`
	Tags           []*FollowUserTag `json:"tags,omitempty"`
	TagId          []string         `json:"tag_id"`
	RemarkCorpName string           `json:"remark_corp_name"`
	RemarkMobiles  []string         `json:"remark_mobiles,omitempty"` // 该成员对此客户备注的手机号码，代开发自建应用需要管理员授权才可以获取，第三方不可获取
	OperUserId     string           `json:"oper_userid"`
	AddWay         int              `json:"add_way"`
	WechatChannels WechatChannels   `json:"wechat_channels"`
	State          string           `json:"state"`
}

type ExternalContact struct {
	ExternalUserId  string          `json:"external_userid"`
	Name            string          `json:"name"`
	Position        string          `json:"position"`
	Avatar          string          `json:"avatar"`
	CorpName        string          `json:"corp_name"`
	CorpFullName    string          `json:"corp_full_name"`
	Type            int             `json:"type"`
	Gender          int8            `json:"gender"`
	UnionId         string          `json:"unionid"` // 仅限自建应用
	ExternalProfile ExternalProfile `json:"external_profile"`
}

type ExternalContactWithFollowUser struct {
	ExternalContact ExternalContact `json:"external_contact"`
	FollowInfo      FollowUser      `json:"follow_info"`
}

type ExternalContactGetResp struct {
	CommonResp
	ExternalContact ExternalContact `json:"external_contact"`
	FollowUser      []*FollowUser   `json:"follow_user"`
}

type ExternalContactListResp struct {
	CommonResp
	ExternalUserId []string `json:"external_userid"`
}

type ExternalContactBatchGetByUserReq struct {
	UserIdList []string `json:"userid_list"` // 上限100
	Cursor     string   `json:"cursor"`
	Limit      int      `json:"limit"` // 上限100，默认50
}

type ExternalContactBatchGetByUserResp struct {
	CommonResp
	ExternalContactList []*ExternalContactWithFollowUser `json:"external_contact_list"`
	NextCursor          string                           `json:"next_cursor"`
}

type ExternalContactRemarkReq struct {
	UserId           string   `json:"userid"`
	ExternalUserId   string   `json:"external_userid"`
	Remark           string   `json:"remark"`
	Description      string   `json:"description"`
	RemarkCompany    string   `json:"remark_company"`           // 只在外部联系人为微信用户时有效
	RemarkMobiles    []string `json:"remark_mobiles,omitempty"` // 如果要清除，需填写一个空字符串
	RemarkPicMediaId string   `json:"remark_pic_mediaid"`
}

type ExternalContactGetCorpTagListReq struct {
	TagId   []string `json:"tag_id"` // group_id有值时忽略
	GroupId []string `json:"group_id"`
}

type Tag struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted"`
}

type TagGroup struct {
	GroupId    string `json:"group_id"`
	GroupName  string `json:"group_name"`
	CreateTime int64  `json:"create_time"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted"`
	Tag        []*Tag `json:"tag"`
}

type ExternalContactGetCorpTagListResp struct {
	CommonResp
	TagGroup []*TagGroup `json:"tag_group"`
}

type ExternalContactAddCorpTagReq struct {
	GroupId   string                          `json:"group_id,omitempty"`
	GroupName string                          `json:"group_name,omitempty"`
	Order     int                             `json:"order,omitempty"`
	Tag       []ExternalContactEditCorpTagReq `json:"tag,omitempty"`
}

type ExternalContactAddCorpTagResp struct {
	CommonResp
	TagGroup TagGroup `json:"tag_group"`
}

type ExternalContactEditCorpTagReq struct {
	Id    string `json:"tag,omitempty"`
	Name  string `json:"name,omitempty"`
	Order int    `json:"order,omitempty"`
}

type ExternalContactMarkTagReq struct {
	UserId         string   `json:"userid"`
	ExternalUserId string   `json:"external_userid"`
	AddTag         []string `json:"add_tag,omitempty"`
	RemoveTag      []string `json:"remove_tag,omitempty"`
}

type UnionidToExternalUseridReq struct {
	UnionId     string `json:"unionid"`
	OpenId      string `json:"openid"`
	SubjectType int    `json:"subject_type"`
}

type ExternalPendingId struct {
	ExternalUserid string `json:"external_userid"`
	PendingId      string `json:"pending_id"`
}

type UnionidToExternalUseridResp struct {
	CommonResp
	ExternalPendingId
}

type ExternalUserIdToPendingIdReq struct {
	ChatId         string   `json:"chat_id,omitempty"`
	ExternalUserId []string `json:"external_userid"`
}

type ExternalUserIdToPendingIdResp struct {
	CommonResp
	Result []ExternalPendingId `json:"result"`
}
