package wxcommon

import (
	"encoding/json"
)

type SuiteEventBase struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Event        string
	EventKey     string
	AgentID      int
}

type SuiteKfEvent struct {
	SuiteEventBase
	Token    string
	OpenKfId string
}

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
	KfMsgContent
	KfEventBase
}

type TransServiceStateReq struct {
	OpenKfId       string `json:"open_kfid"`
	ExternalUserId string `json:"external_userid"`
	ServiceState
}

type TransServiceStateResp struct {
	CommonResp
	MsgCode string `json:"msg_code"`
}

type GetServiceStateResp struct {
	CommonResp
	ServiceState
}
type ServiceState struct {
	ServiceState   KfServiceState `json:"service_state"`
	ServicerUserId string         `json:"servicer_userid"`
}

type KfSendMsgRequest struct {
	ToUser   string `json:"touser"`
	OpenKfID string `json:"open_kfid"`
	KfMsgBody
}
type KfMsgBody struct {
	Code    string `json:"code,omitempty"`
	MsgID   string `json:"msgid,omitempty"`
	MsgType string `json:"msgtype"`
	KfMsgContent
}
type KfMsgContent struct {
	Text        *KfMsgText        `json:"text,omitempty"`
	Msgmenu     *KfMsgMenu        `json:"msgmenu,omitempty"`
	Image       *KfMsgImage       `json:"image,omitempty"`
	Voice       *KfMsgVoice       `json:"voice,omitempty"`
	Video       *KfMsgVideo       `json:"video,omitempty"`
	File        *KfMsgFile        `json:"file,omitempty"`
	Link        *KfMsgLink        `json:"link,omitempty"`
	Miniprogram *KfMsgMiniprogram `json:"miniprogram,omitempty"`
}
type KfSendMsgResp struct {
	CommonResp
	MsgID string `json:"msgid"`
}
type KfMsgText struct {
	Content string `json:"content"`
}
type KfMsgImage struct {
	MediaID string `json:"media_id"`
}
type KfMsgVoice struct {
	MediaID string `json:"media_id"`
}
type KfMsgVideo struct {
	MediaID string `json:"media_id"`
}
type KfMsgFile struct {
	MediaID string `json:"media_id"`
}
type KfMsgLink struct {
	MediaID      string `json:"media_id"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	Url          string `json:"url"`
	ThumbMediaID string `json:"thumb_media_id"`
}
type KfMsgMiniprogram struct {
	AppID        string `json:"appid"`
	Title        string `json:"title"`
	ThumbMediaID string `json:"thumb_media_id,omitempty"`
	Pagepath     string `json:"pagepath"`
}

type KfMsgMenu struct {
	HeadContent string          `json:"head_content,omitempty"`
	List        []KfMsgMenuItem `json:"list,omitempty"`
	TailContent string          `json:"tail_content,omitempty"`
}
type KfMsgMenuItem struct {
	Type        string                    `json:"type"`
	Click       *KfMsgMenuItemClick       `json:"click,omitempty"`
	View        *KfMsgMenuItemView        `json:"view,omitempty"`
	Miniprogram *KfMsgMenuItemMiniprogram `json:"miniprogram,omitempty"`
}

type KfMsgMenuItemClick struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

type KfMsgMenuItemView struct {
	Url     string `json:"url"`
	Content string `json:"content"`
}
type KfMsgMenuItemMiniprogram struct {
	AppId    string `json:"appid"`
	Pagepath string `json:"pagepath"`
	Content  string `json:"content"`
}

type KfServicerListResp struct {
	CommonResp
	List []KfServicer `json:"servicer_list"`
}

type KfServicer struct {
	UserId       string `json:"userid"`
	DepartmentId int    `json:"department_id"`
	Status       int    `json:"status"`
	StopType     int    `json:"stop_type"`
}

type KfEventBase struct {
	EventType string          `json:"event_type,omitempty"`
	Event     json.RawMessage `json:"event,omitempty"`
}

func ParseKfEvent[T any](raw json.RawMessage) (event *T, err error) {
	event = new(T)
	err = json.Unmarshal(raw, event)
	return event, err
}

type KfServiceStatusEvent struct {
	KfEventBase
	ServiceUserId string `json:"servicer_userid"`
	Status        int    `json:"status"`
	StopType      int    `json:"stop_type"`
	OpenKfId      string `json:"open_kfid"`
}

type KfEntetrSessionEvent struct {
	KfEventBase
	OpenKfId       string `json:"open_kfid"`
	ExternalUserId string `json:"external_userid"`
	Scene          string `json:"scene"`
	SceneParam     string `json:"scene_param"`
	WelcomeCode    string `json:"welcome_code"`
	WechatChannels struct {
		NickName string `json:"nickname"`
		Scene    int    `json:"scene"`
	} `json:"wechat_channels"`
}

type KfSessionStatusEvent struct {
	KfEventBase
	OpenKfId          string `json:"open_kfid"`
	ExternalUserId    string `json:"external_userid"`
	ChangeType        int    `json:"change_type"`
	OldServicerUserId string `json:"old_servicer_userid"`
	NewServicerUserId string `json:"new_servicer_userid"`
	MsgCode           string `json:"msg_code"`
}

type KfEditServicerResult struct {
	UserId       string `json:"userid"`
	DepartmentId string `json:"department_id"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}
type KfEditServicerResp struct {
	CommonResp
	List []KfEditServicerResult `json:"result_list"`
}

type KfGetCorpStatisticResp struct {
	CommonResp
	Statistic []KfCorpStatisticItem `json:"statistic_list"`
}

type KfCorpStatisticItem struct {
	StatTime  int64 `json:"stat_time"`
	Statistic struct {
		SessionCnt                int `json:"session_cnt"`
		CustomerCnt               int `json:"customer_cnt"`
		CustomerMsgCnt            int `json:"customer_msg_cnt"`
		UpgradeServiceCustomerCnt int `json:"upgrade_service_customer_cnt"`
		AiSessionReplyCnt         int `json:"ai_session_reply_cnt"`
		AiTransferRate            int `json:"ai_transfer_rate"`
		AiKnowledgeHitRate        int `json:"ai_knowledge_hit_rate"`
		MsgRejectedCustomerCnt    int `json:"msg_rejected_customer_cnt"`
	} `json:"statistic"`
}

type KfGetServicerStatisticResp struct {
	CommonResp
	Statistic []KfServicerStatisticItem `json:"statistic_list"`
}

type KfServicerStatisticItem struct {
	StatTime  int64 `json:"stat_time"`
	Statistic struct {
		SessionCnt                         int `json:"session_cnt"`
		CustomerCnt                        int `json:"customer_cnt"`
		CustomerMsgCnt                     int `json:"customer_msg_cnt"`
		ReplyRate                          int `json:"reply_rate"`
		FirstReplyAverageSec               int `json:"first_reply_average_sec"`
		SatisfactionInvestgateCnt          int `json:"satisfaction_investgate_cnt"`
		SatisfactionParticipationRate      int `json:"satisfaction_participation_rate"`
		SatisfiedRate                      int `json:"satisfied_rate"`
		MiddlingRate                       int `json:"middling_rate"`
		DissatisfiedRate                   int `json:"dissatisfied_rate"`
		UpgradeServiceCustomerCnt          int `json:"upgrade_service_customer_cnt"`
		UpgradeServiceMemberInviteCnt      int `json:"upgrade_service_member_invite_cnt"`
		UpgradeServiceMemberCustomerCnt    int `json:"upgrade_service_member_customer_cnt"`
		UpgradeServiceGroupchatInviteCnt   int `json:"upgrade_service_groupchat_invite_cnt"`
		UpgradeServiceGroupchatCustomerCnt int `json:"upgrade_service_groupchat_customer_cnt"`
		MsgRejectedCustomerCnt             int `json:"msg_rejected_customer_cnt"`
	}
}
