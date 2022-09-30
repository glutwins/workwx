package wxcommon

import (
	"fmt"
)

const (
	MsgTypeText              = "text"
	MsgTypeImage             = "image"
	MsgTypeVoice             = "voice"
	MsgTypeVideo             = "video"
	MsgTypeFile              = "file"
	MsgTypeTextCard          = "textcard"
	MsgTypeNews              = "news"
	MsgTypeArticles          = "articles"
	MsgTypeMarkdown          = "markdown"
	MsgTypeTemplateCard      = "template_card"
	MsgTypeMiniprogramNotice = "miniprogram_notice"
)

type MessageText struct {
	Content string `json:"content"`
}

type MessageMedia struct {
	MediaId string `json:"media_id"`
}

type MessageDesc struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type MessageVideo struct {
	MessageMedia
	MessageDesc
}

type MessageTextCard struct {
	MessageDesc
	URL    string `json:"url"`
	Btntxt string `json:"btntxt"`
}

type MessageArticle struct {
	MessageDesc
	URL      string `json:"url"`
	PicURL   string `json:"picurl"`
	AppId    string `json:"appid"`
	PagePath string `json:"pagepath"`
}

type MessageNews struct {
	Articles []MessageArticle `json:"articles"`
}

/*
   "appid": "wx123123123123123",
   "page": "pages/index?userid=zhangsan&orderid=123123123",
   "title": "会议室预订成功通知",
   "description": "4月27日 16:16",
   "emphasis_first_item": true,
   "content_item": [
       {
           "key": "会议室",
           "value": "402"
       },
   ]*/
type ContentItem struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// MessageMiniprogramNotice
type MessageMiniprogramNotice struct {
	AppId             string        `json:"appid"`
	Page              string        `json:"page"`
	Title             string        `json:"title"`
	Description       string        `json:"description"`
	EmphasisFirstItem bool          `json:"emphasis_first_item"`
	ContentItem       []ContentItem `json:"content_item"`
}

type MessageTemplateCardSource struct {
	IconURL   string `json:"icon_url"`
	Desc      string `json:"desc"`
	DescColor string `json:"desc_color"`
}
type Action struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}
type MessageTemplateCardActionMenu struct {
	Desc       string   `json:"desc"`
	ActionList []Action `json:"action_list"`
}

type MessageTemplateCard struct {
	CardType string
}

type Message struct {
	ToUser                 string                    `json:"touser"`
	ToParty                string                    `json:"toparty"`
	ToTag                  string                    `json:"totag"`
	MsgType                string                    `json:"msgtype"`
	Text                   *MessageText              `json:"text,omitempty"`
	Image                  *MessageMedia             `json:"image,omitempty"`
	Voice                  *MessageMedia             `json:"voice,omitempty"`
	File                   *MessageMedia             `json:"file,omitempty"`
	Video                  *MessageVideo             `json:"video,omitempty"`
	TextCard               *MessageTextCard          `json:"textcard,omitempty"`
	News                   *MessageNews              `json:"news"`
	MiniprogramNotice      *MessageMiniprogramNotice `json:"miniprogram_notice"`
	AgentId                int                       `json:"agentid"`
	Safe                   int                       `json:"safe"`
	EnableIdTrans          int                       `json:"enable_id_trans"`
	EnableDuplicateCheck   int                       `json:"enable_duplicate_check"`
	DuplicateCheckInterval int                       `json:"duplicate_check_interval"`
}

type MessageSendResp struct {
	CommonResp
	InvalidUser    string `json:"invaliduser"`
	InvalidParty   string `json:"invalidparty"`
	InvalidTag     string `json:"invalidtag"`
	UnlicensedUser string `json:"unlicenseduser"`
	MsgId          string `json:"msgid"`
	ResponseCode   string `json:"response_code"`
}

func (scc *SuiteCorpClient) MessageSend(msg *Message) (*MessageSendResp, error) {
	token, err := scc.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp := &MessageSendResp{}
	if err := scc.PostJSON(fmt.Sprintf("/message/send?access_token=%s", token), msg, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
