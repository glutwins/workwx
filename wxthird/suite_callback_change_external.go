package wxthird

type SuiteCallbackExternalUser struct {
	UserID         string
	ExternalUserID string
	State          string
	WelcomeCode    string
	Source         string
	FailReason     string
}

type SuiteCallbackExternalChat struct {
	ChatId       string
	UpdateDetail string
	JoinScene    int
	QuitScene    int
	MemChangeCnt int
}

type SuiteCallbackExternalTag struct {
	Id      string
	TagType string
}
