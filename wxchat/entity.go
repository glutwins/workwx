package wxchat

import (
	"encoding/json"
	"fmt"
)

// CodeMsg 错误代码定义
var CodeMsg = map[int]string{
	10000: "参数错误，请求参数错误",
	10001: "网络错误，网络请求错误",
	10002: "数据解析失败",
	10003: "系统失败",
	10004: "密钥错误导致加密失败",
	10005: "fileid错误",
	10006: "解密失败",
	10007: "找不到消息加密版本的私钥，需要重新传入私钥对",
	10008: "解析encrypt_key出错",
	10009: "ip非法",
	10010: "数据过期",
	10011: "证书错误",
}

func NewSdkError(code int) error {
	return fmt.Errorf("sdk: code=%d, msg=%s", code, CodeMsg[code])
}

const (
	// MsgTypeText 消息类型: Text
	MsgTypeText = "text"
	// MsgTypeImage 消息类型: Image
	MsgTypeImage = "image"
	// MsgTypeRevoke 消息类型: Revoke
	MsgTypeRevoke = "revoke"
	// MsgTypeAgree 消息类型: Agree
	MsgTypeAgree = "agree"
	// MsgTypeVoice 消息类型: Voice
	MsgTypeVoice = "voice"
	// MsgTypeVideo 消息类型: Video
	MsgTypeVideo = "video"
	// MsgTypeCard 消息类型: Card
	MsgTypeCard = "card"
	// MsgTypeLocation 消息类型: Location
	MsgTypeLocation = "location"
	// MsgTypeEmotion 消息类型: Emotion
	MsgTypeEmotion = "emotion"
	// MsgTypeFile 消息类型: File
	MsgTypeFile = "file"
	// MsgTypeLink 消息类型: Link
	MsgTypeLink = "link"
	// MsgTypeWeapp 消息类型: Weapp
	MsgTypeWeapp = "weapp"
	// MsgTypeChatrecord 消息类型: Chatrecord
	MsgTypeChatrecord = "chatrecord"
	// MsgTypeTodo 消息类型: Todo
	MsgTypeTodo = "todo"
	// MsgTypeVote 消息类型: Vote
	MsgTypeVote = "vote"
	// MsgTypeCollect 消息类型: Collect
	MsgTypeCollect = "collect"
	// MsgTypeRedpacket 消息类型: Redpacket
	MsgTypeRedpacket = "redpacket"
	// MsgTypeMeeting 消息类型: Meeting
	MsgTypeMeeting = "meeting"
	// MsgTypeDocmsg 消息类型: Docmsg
	MsgTypeDocmsg = "docmsg"
	// MsgTypeMarkdown 消息类型: Markdown
	MsgTypeMarkdown = "markdown"
	// MsgTypeInfo 消息类型: Info
	MsgTypeInfo = "info"
	// MsgTypeCalendar 消息类型: Calendar
	MsgTypeCalendar = "calendar"
	// MsgTypeMixed 消息类型: Mixed
	MsgTypeMixed = "mixed"
	// MsgTypeMeetingVoiceCall 消息类型: MeetingVoiceCall
	MsgTypeMeetingVoiceCall = "meeting_voice_call"
	// MsgTypeVoipDocShare 消息类型: VoipDocShare
	MsgTypeVoipDocShare = "voip_doc_share"
)

// EncryptData 聊天加密数据
type EncryptData struct {
	ErrCode  int                   `json:"errcode"`
	ErrMsg   string                `json:"errmsg"`
	ChatData []EncryptDataChatData `json:"chatdata"`
}

// EncryptDataChatData 密文
type EncryptDataChatData struct {
	Seq              int64  `json:"seq"`
	MsgID            string `json:"msgid"`
	PublicKeyVer     int    `json:"publickey_ver"`
	EncryptRandomKey string `json:"encrypt_random_key"`
	EncryptChatMsg   string `json:"encrypt_chat_msg"`
}

// PlainMsg 聊天明文数据
type PlainMsg struct {
	MsgID            string           `json:"msgid" desc:"消息id，消息的唯一标识，企业可以使用此字段进行消息去重"`
	Action           string           `json:"action" desc:"消息动作，目前有send(发送消息)/recall(撤回消息)/switch(切换企业日志)三种类型"`
	From             string           `json:"from" desc:"消息发送方id 同一企业内容为userid，非相同企业为external_userid 消息如果是机器人发出，也为external_userid"`
	Tolist           []string         `json:"tolist" desc:"消息接收方列表，可能是多个，同一个企业内容为userid，非相同企业为external_userid 数组"`
	Roomid           string           `json:"roomid" desc:"群聊消息的群id 如果是单聊则为空"`
	Msgtime          int64            `json:"msgtime" desc:"消息发送时间戳，utc时间，ms单位"`
	Msgtype          string           `json:"msgtype" desc:"消息类型"`
	Text             Text             `json:"text" desc:"文本消息"`
	Image            Image            `json:"image" desc:"图片消息"`
	Revoke           Revoke           `json:"revoke" desc:"撤回消息"`
	Agree            Agree            `json:"agree" desc:"同意消息"`
	Voice            Voice            `json:"voice" desc:"语音消息"`
	Video            Video            `json:"video" desc:"视频消息"`
	Card             Card             `json:"card" desc:"名片消息"`
	Location         Location         `json:"location" desc:"位置消息"`
	Emotion          Emotion          `json:"emotion" desc:"表情消息"`
	File             File             `json:"file" desc:"文件消息"`
	Link             Link             `json:"link" desc:"链接"`
	Weapp            Weapp            `json:"weapp" desc:"小程序消息"`
	Chatrecord       Chatrecord       `json:"chatrecord" desc:"会话记录消息"`
	Todo             Todo             `json:"todo" desc:"待办消息"`
	Vote             Vote             `json:"vote" desc:"投票消息"`
	Collect          Collect          `json:"collect" desc:"填表消息"`
	Redpacket        Redpacket        `json:"redpacket" desc:"红包消息"`
	Meeting          Meeting          `json:"meeting" desc:"会议邀请消息"`
	Docmsg           Docmsg           `json:"doc" desc:"微信字段不一致，请勿随意修改! 在线文档消息"`
	Markdown         Markdown         `json:"markdown" desc:"MarkDown格式消息"`
	News             News             `json:"info" desc:"微信字段不一致，请勿随意修改! 图文消息"`
	Calendar         Calendar         `json:"calendar" desc:"日程消息"`
	Mixed            Mixed            `json:"mixed" desc:"混合消息"`
	MeetingVoiceCall MeetingVoiceCall `json:"meeting_voice_call" desc:"音频存档消息"`
	Voiceid          string           `json:"voiceid" desc:"是个坑! 只有meeting_voice_call类型存在 音频id"`
	VoipDocShare     VoipDocShare     `json:"voip_doc_share" desc:"音频共享文档消息"`
}

// Content 获取消息体对应类型的内容json
func (msg *PlainMsg) Content() string {
	var val interface{} = make(map[string]interface{})
	switch msg.Msgtype {
	case MsgTypeText:
		val = &msg.Text
	case MsgTypeImage:
		val = &msg.Image
	case MsgTypeRevoke:
		val = &msg.Revoke
	case MsgTypeAgree:
		val = &msg.Agree
	case MsgTypeVoice:
		val = &msg.Voice
	case MsgTypeVideo:
		val = &msg.Video
	case MsgTypeCard:
		val = &msg.Card
	case MsgTypeLocation:
		val = &msg.Location
	case MsgTypeEmotion:
		val = &msg.Emotion
	case MsgTypeFile:
		val = &msg.File
	case MsgTypeLink:
		val = &msg.Link
	case MsgTypeWeapp:
		val = &msg.Weapp
	case MsgTypeChatrecord:
		val = &msg.Chatrecord
	case MsgTypeTodo:
		val = &msg.Todo
	case MsgTypeVote:
		val = &msg.Vote
	case MsgTypeCollect:
		val = &msg.Collect
	case MsgTypeRedpacket:
		val = &msg.Redpacket
	case MsgTypeMeeting:
		val = &msg.Meeting
	case MsgTypeDocmsg:
		val = &msg.Docmsg
	case MsgTypeMarkdown:
		val = &msg.Markdown
	case MsgTypeInfo:
		val = &msg.News
	case MsgTypeCalendar:
		val = &msg.Calendar
	case MsgTypeMixed:
		val = &msg.Mixed
	case MsgTypeMeetingVoiceCall:
		msg.MeetingVoiceCall.Voiceid = msg.Voiceid
		val = &msg.MeetingVoiceCall
	case MsgTypeVoipDocShare:
		val = &msg.VoipDocShare
	}
	b, _ := json.Marshal(val)
	return string(b)
}

// SdkFileId 获取消息附件
func (msg *PlainMsg) SdkFileId() string {
	switch msg.Msgtype {
	case MsgTypeImage:
		return msg.Image.Sdkfileid
	case MsgTypeVoice:
		return msg.Voice.Sdkfileid
	case MsgTypeVideo:
		return msg.Video.Sdkfileid
	case MsgTypeEmotion:
		return msg.Emotion.Sdkfileid
	case MsgTypeFile:
		return msg.File.Sdkfileid
	case MsgTypeMeetingVoiceCall:
		return msg.MeetingVoiceCall.Sdkfileid
	case MsgTypeVoipDocShare:
		return msg.VoipDocShare.Sdkfileid
	}
	return ""
}

// Text 文本消息
type Text struct {
	Content string `json:"content" desc:"消息内容"`
}

// Image 图片消息
type Image struct {
	Sdkfileid string `json:"sdkfileid" desc:"媒体资源的id信息"`
	Md5sum    string `json:"md5sum" desc:"图片资源的md5值，供进行校验"`
	Filesize  int64  `json:"filesize" desc:"图片资源的文件大小"`
}

// Revoke 撤回消息
type Revoke struct {
	PreMsgid string `json:"pre_msgid" desc:"标识撤回的原消息的msgid"`
}

// Agree 同意消息
type Agree struct {
	Userid    string `json:"userid" desc:"同意/不同意协议者的userid，外部企业默认为external_userid"`
	AgreeTime int64  `json:"agree_time" desc:"同意/不同意协议的时间，utc时间，ms单位"`
}

// Voice 语音消息
type Voice struct {
	VoiceSize  int64  `json:"voice_size" desc:"语音消息大小"`
	PlayLength int64  `json:"play_length" desc:"播放长度"`
	Sdkfileid  string `json:"sdkfileid" desc:"媒体资源的id信息"`
	Md5sum     string `json:"md5sum" desc:"资源的md5值，供进行校验"`
}

// Video 视频消息
type Video struct {
	Sdkfileid  string `json:"sdkfileid" desc:"媒体资源的id信息"`
	Md5sum     string `json:"md5sum" desc:"资源的md5值，供进行校验"`
	Filesize   int64  `json:"filesize" desc:"资源的文件大小"`
	PlayLength int64  `json:"play_length" desc:"视频播放长度"`
}

// Card 卡片消息
type Card struct {
	Corpname string `json:"corpname" desc:"名片所有者所在的公司名称"`
	Userid   string `json:"userid" desc:"名片所有者的id，同一公司是userid，不同公司是external_userid"`
}

// Location 表情消息
type Location struct {
	Longitude float64 `json:"longitude" desc:"经度"`
	Latitude  float64 `json:"latitude" desc:"纬度"`
	Address   string  `json:"address" desc:"地址信息"`
	Title     string  `json:"title" desc:"位置信息的title"`
	Zoom      float64 `json:"zoom" desc:"缩放比例"`
}

// Emotion 表情消息
type Emotion struct {
	Type      int64  `json:"type" desc:"表情类型，png或者gif.1表示gif 2表示png"`
	Width     int64  `json:"width" desc:"表情图片宽度"`
	Height    int64  `json:"height" desc:"表情图片高度"`
	Sdkfileid string `json:"sdkfileid" desc:"媒体资源的id信息"`
	Md5sum    string `json:"md5sum" desc:"资源的md5值，供进行校验"`
	Imagesize int64  `json:"imagesize" desc:"资源的文件大小"`
}

// File 文件消息
type File struct {
	Sdkfileid string `json:"sdkfileid" desc:"媒体资源的id信息"`
	Md5sum    string `json:"md5sum" desc:"资源的md5值，供进行校验"`
	Filename  string `json:"filename" desc:"文件名称"`
	Fileext   string `json:"fileext" desc:"文件类型后缀"`
	Filesize  int64  `json:"filesize" desc:"文件大小"`
}

// Link 链接消息
type Link struct {
	Title       string `json:"title" desc:"消息标题"`
	Description string `json:"description" desc:"消息描述"`
	LinkURL     string `json:"link_url" desc:"链接url地址"`
	ImageURL    string `json:"image_url" desc:"链接图片url"`
}

// Weapp 小程序消息
type Weapp struct {
	Title       string `json:"title" desc:"消息标题"`
	Description string `json:"description" desc:"消息描述"`
	Username    string `json:"username" desc:"用户名称"`
	Displayname string `json:"displayname" desc:"小程序名称"`
}

// Chatrecord 会话记录消息
type Chatrecord struct {
	Title string           `json:"title" desc:"聊天记录标题"`
	Item  []ChatrecordItem `json:"item" desc:"消息记录内的消息内容，批量数据"`
}

// ChatrecordItem 消息记录内的消息内容，批量数据
type ChatrecordItem struct {
	Type         string `json:"type" desc:"每条聊天记录的具体消息类型"`
	Msgtime      int64  `json:"msgtime" desc:"消息时间，utc时间，ms单位"`
	Content      string `json:"content" desc:"消息内容。Json串，内容为对应类型的json"`
	FromChatroom bool   `json:"from_chatroom" desc:"是否来自群会话"`
}

// Todo 待办消息
type Todo struct {
	Title   string `json:"title" desc:"代办的来源文本"`
	Content string `json:"content" desc:"代办的具体内容"`
}

// Vote 投票消息
type Vote struct {
	Votetitle string   `json:"votetitle" desc:"投票主题"`
	Voteitem  []string `json:"voteitem" desc:"投票选项，可能多个内容"`
	Votetype  int64    `json:"votetype" desc:"投票类型.101发起投票、102参与投票"`
	Voteid    string   `json:"voteid" desc:"投票id，方便将参与投票消息与发起投票消息进行前后对照"`
}

// Collect 填表消息
type Collect struct {
	RoomName   string          `json:"room_name" desc:"填表消息所在的群名称"`
	Creator    string          `json:"creator" desc:"创建者在群中的名字"`
	CreateTime string          `json:"create_time" desc:"create_time"`
	Title      string          `json:"title" desc:"表名"`
	Details    []CollectDetail `json:"details" desc:"表内容"`
}

// CollectDetail 表内容
type CollectDetail struct {
	ID   int64  `json:"id" desc:"表项id"`
	Ques string `json:"ques" desc:"表项名称"`
	Type string `json:"type" dsc:"表项类型，有Text(文本),Number(数字),Date(日期),Time(时间)"`
}

// Redpacket 红包消息
type Redpacket struct {
	Type        int64  `json:"type" desc:"红包消息类型。1 普通红包、2 拼手气群红包、3 激励群红包"`
	Wish        string `json:"wish" desc:"红包祝福语"`
	Totalcnt    int64  `json:"totalcnt" desc:"红包总个数"`
	Totalamount int64  `json:"totalamount" desc:"红包总金额。单位为分"`
}

// Meeting 会议邀请消息
type Meeting struct {
	Topic       string `json:"topic" desc:"会议主题"`
	Starttime   int64  `json:"starttime" desc:"会议开始时间"`
	Endtime     int64  `json:"endtime" desc:"会议结束时间"`
	Address     string `json:"address" desc:"会议地址"`
	Remarks     string `json:"remarks" desc:"会议备注"`
	Meetingtype int64  `json:"meetingtype" desc:"会议消息类型。101发起会议邀请消息、102处理会议邀请消息"`
	Meetingid   int64  `json:"meetingid" desc:"会议id。方便将发起、处理消息进行对照"`
	Status      int64  `json:"status" desc:"会议邀请处理状态。1 参加会议、2 拒绝会议、3 待定、4 未被邀请、5 会议已取消、6 会议已过期、7 不在房间内。只有meetingtype为102的时候此字段才有内容。"`
}

// Docmsg 在线文档消息
type Docmsg struct {
	Title      string `json:"title" desc:"在线文档名称"`
	LinkURL    string `json:"link_url" desc:"在线文档链接"`
	DocCreator string `json:"doc_creator" desc:"在线文档创建者。本企业成员创建为userid；外部企业成员创建为external_userid"`
}

// Markdown MarkDown格式消息
type Markdown struct {
	Content string `json:"content" desc:"markdown消息内容，目前为机器人发出的消息"`
}

// News 图文消息
type News struct {
	Info NewsInfo `json:"info" desc:"图文消息的内容"`
}

// NewsInfo 图文消息的内容
type NewsInfo struct {
	Item []NewsInfoItem `json:"item" desc:"图文消息数组，每个item结构包含title、description、url、picurl等结构"`
}

// NewsInfoItem 图文消息数组
type NewsInfoItem struct {
	Title       string `json:"title" desc:"图文消息标题"`
	Description string `json:"description" desc:"图文消息描述"`
	URL         string `json:"url" desc:"图文消息点击跳转地址"`
	Picurl      string `json:"picurl" desc:"图文消息配图的url"`
}

// Calendar 日程消息
type Calendar struct {
	Title        string   `json:"title" desc:"日程主题"`
	Creatorname  string   `json:"creatorname" desc:"日程组织者"`
	Attendeename []string `json:"attendeename" desc:"日程参与人。数组"`
	Starttime    int64    `json:"starttime" desc:"日程开始时间。Utc时间，单位秒"`
	Endtime      int64    `json:"endtime" desc:"日程结束时间。Utc时间，单位秒"`
	Place        string   `json:"place" desc:"日程地点"`
	Remarks      string   `json:"remarks" desc:"日程备注"`
}

// Mixed 混合消息
type Mixed struct {
	Item []MixedItem `json:"item" desc:"数组其中每个元素由type与content组成，type和content均为String类型。JSON解析content后即可获取对应type类型的消息内容"`
}

// MixedItem 混合消息详情
type MixedItem struct {
	Type       string                 `json:"type" desc:"消息类型"`
	Content    string                 `json:"content" desc:"消息内容"`
	ContentMap map[string]interface{} `json:"content_map" desc:"消息内容map"`
}

// MeetingVoiceCall 音频存档消息
type MeetingVoiceCall struct {
	Voiceid         string                            `json:"voiceid" desc:"不是微信返回的将外层的放入 音频id"`
	Endtime         int64                             `json:"endtime" desc:"音频结束时间"`
	Sdkfileid       string                            `json:"sdkfileid" desc:"音频媒体下载的id"`
	Demofiledata    []MeetingVoiceCallDemofiledata    `json:"demofiledata" desc:"文档分享对象"`
	Sharescreendata []MeetingVoiceCallSharescreendata `json:"sharescreendata" desc:"屏幕共享对象"`
}

// MeetingVoiceCallDemofiledata 文档分享对象
type MeetingVoiceCallDemofiledata struct {
	Filename     string `json:"filename" desc:"文档共享名称"`
	Demooperator string `json:"demooperator" desc:"文档共享操作用户的id"`
	Starttime    int64  `json:"starttime" desc:"文档共享开始时间"`
	Endtime      int64  `json:"endtime" desc:"文档共享结束时间"`
}

// MeetingVoiceCallSharescreendata 屏幕共享对象
type MeetingVoiceCallSharescreendata struct {
	Share     string `json:"share" desc:"屏幕共享用户的id"`
	Starttime int64  `json:"starttime" desc:"屏幕共享开始时间"`
	Endtime   int64  `json:"endtime" desc:"屏幕共享结束时间"`
}

// VoipDocShare 音频共享文档消息
type VoipDocShare struct {
	Filename  string `json:"filename" desc:"文档共享文件名称"`
	Md5sum    string `json:"md5sum" desc:"共享文件的md5值"`
	Filesize  int64  `json:"filesize" desc:"共享文件的大小"`
	Sdkfileid string `json:"sdkfileid" desc:"共享文件的sdkfile，通过此字段进行媒体数据下载"`
}
