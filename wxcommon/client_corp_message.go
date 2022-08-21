package wxcommon

import (
	"fmt"
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

type Message struct {
	ToUser                 string           `json:"touser"`
	ToParty                string           `json:"toparty"`
	ToTag                  string           `json:"totag"`
	MsgType                string           `json:"msgtype"`
	Text                   *MessageText     `json:"text,omitempty"`
	Image                  *MessageMedia    `json:"image,omitempty"`
	Voice                  *MessageMedia    `json:"voice,omitempty"`
	File                   *MessageMedia    `json:"file,omitempty"`
	Video                  *MessageVideo    `json:"video,omitempty"`
	TextCard               *MessageTextCard `json:"textcard,omitempty"`
	News                   *MessageNews     `json:"news"`
	AgentId                int              `json:"agentid"`
	Safe                   int              `json:"safe"`
	EnableIdTrans          int              `json:"enable_id_trans"`
	EnableDuplicateCheck   int              `json:"enable_duplicate_check"`
	DuplicateCheckInterval int              `json:"duplicate_check_interval"`
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
