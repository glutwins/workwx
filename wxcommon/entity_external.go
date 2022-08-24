package wxcommon

type ExternalContactSendWelcomeMsgReq struct {
	WelcomeCode string        `json:"welcome_code"`
	Text        *Text         `json:"text,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}
