package wxcommon

type ExternalContactSendWelcomeMsgReq struct {
	WelcomeCode string        `json:"welcome_code"`
	Text        *Text         `json:"text,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
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
	OperUserID     string           `json:"oper_userid"`
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
	Id         string `json:"tag"`
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
	GroupId   string                          `json:"group_id"`
	GroupName string                          `json:"group_name"`
	Order     int                             `json:"order"`
	Tag       []ExternalContactEditCorpTagReq `json:"tag"`
}

type ExternalContactAddCorpTagResp struct {
	CommonResp
	TagGroup TagGroup `json:"tag_group"`
}

type ExternalContactEditCorpTagReq struct {
	Id    string `json:"tag"`
	Name  string `json:"name"`
	Order int    `json:"order"`
}

type ExternalContactMarkTagReq struct {
	UserId         string   `json:"userid"`
	ExternalUserId string   `json:"external_userid"`
	AddTag         []string `json:"add_tag,omitempty"`
	RemoveTag      []string `json:"remove_tag"`
}
