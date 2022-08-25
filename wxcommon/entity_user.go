package wxcommon

type ExtAttrText struct {
	Value string `json:"value"`
}

type ExtAttrWeb struct {
	Title string `json:"title"`
	Url   string `json:"url"`
}

type ExtAttrMiniprogram struct {
	AppId    string `json:"appid"`
	PagePath string `json:"pagepath"`
	Title    string `json:"title"`
}

type UserExtAttr struct {
	Attrs []*ExtAttr `json:"attrs"`
}

type UserSimple struct {
	UserId     string `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department"`
	OpenUserId string `json:"open_userid"`
}

type UserExternalProfile struct {
	ExternalCorpName string         `json:"external_corp_name"`
	WechatChannels   WechatChannels `json:"wechat_channels"`
	ExternalAttr     []*ExtAttr     `json:"external_attr"`
}

type UserProfile struct {
	AvatarMediaid    string              `json:"avatar_mediaid"` // 创建成员时
	Avatar           string              `json:"avatar"`         // 获取成员时
	ThumbAvatar      string              `json:"thumb_avatar"`   // 获取成员时
	Alias            string              `json:"alias"`
	Mobile           string              `json:"mobile"`
	Order            []int               `json:"order"`
	Position         string              `json:"position"`
	Gender           string              `json:"gender"`
	Email            string              `json:"email"`
	BizMail          string              `json:"biz_mail"`
	IsLeaderInDept   []int8              `json:"is_leader_in_dept"`
	DirectLeader     []string            `json:"direct_leader"`
	Telephone        string              `json:"telephone"`
	Address          string              `json:"address"`
	MainDepartment   int                 `json:"main_department"`
	ExtAttr          UserExtAttr         `json:"extattr"`
	ExternalPosition string              `json:"external_position"`
	ExternalProfile  UserExternalProfile `json:"external_profile"`
}

type UserCreateReq struct {
	UserSimple
	UserProfile
	Enable   int  `json:"enable"`
	ToInvite bool `json:"to_invite"`
}

type UserSimpleListResp struct {
	CommonResp
	UserList []*UserSimple `json:"userlist"`
}

type UserGetUserinfoResp struct {
	CommonResp
	UserId   string
	DeviceId string
}

type UserDetail struct {
	UserSimple
	UserProfile
	Status int    `json:"status"` // 激活状态: 1=已激活，2=已禁用，4=未激活，5=退出企业
	QrCode string `json:"qr_code"`
}

type UserGetResp struct {
	CommonResp
	UserDetail
}

type UserUpdateReq struct {
	UserSimple
	UserProfile
	NewUserID string `json:"new_userid"`
	Enable    int    `json:"enable"`
}

type UserListResp struct {
	CommonResp
	UserList []*UserDetail `json:"userlist"`
}
