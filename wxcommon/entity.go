package wxcommon

import (
	"fmt"
)

type CommonResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (r *CommonResp) Error() string {
	return fmt.Sprintf("workwx-err:%d(%s)", r.ErrCode, r.ErrMsg)
}

func (r *CommonResp) Err() error {
	if r.ErrCode != 0 {
		return r
	}
	return nil
}

type WorkWxResp interface {
	Err() error
}

type GetCorpTokenResp struct {
	CommonResp
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type TicketGetResp struct {
	CommonResp
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

type Text struct {
	Content string `json:"content"`
}

type Image struct {
	MediaId string `json:"media_id"`
	PicURL  string `json:"pic_url"`
}

type Link struct {
	Title  string `json:"title"`
	PicURL string `json:"picurl"`
	Desc   string `json:"desc"`
	URL    string `json:"url"`
}

type Miniprogram struct {
	Title      string `json:"title"`
	PicMediaId string `json:"pic_media_id"`
	AppId      string `json:"appid"`
	Page       string `json:"vedio"`
}

type Media struct {
	MediaId string `json:"media_id"`
}

type Attachment struct {
	MsgType     string       `json:"msgtype"`
	Image       *Image       `json:"image,omitempty"`
	Link        *Link        `json:"link,omitempty"`
	Miniprogram *Miniprogram `json:"miniprogram,omitempty"`
	Video       *Media       `json:"video,omitempty"`
	File        *Media       `json:"file,omitempty"`
}

type WechatChannels struct {
	NickName string `json:"nickname"` // 视频号
	Source   int    `json:"source"`   // 视频号添加场景，0-未知 1-视频号主页 2-视频号直播间
}

type ExtAttr struct {
	Name        string              `json:"name"`
	Type        int                 `json:"type"`
	Text        *ExtAttrText        `xml:",omitempty" json:"text,omitempty"`
	Web         *ExtAttrWeb         `xml:",omitempty" json:"web,omitempty"`
	Miniprogram *ExtAttrMiniprogram `xml:",omitempty" json:"miniprogram,omitempty"`
}
