package wxcommon

const (
	CallbackTypeChangeContact         string = "change_contact"
	CallbackTypeChangeExternalContact string = "change_external_contact"
	CallbackTypeChangeExternalChat    string = "change_external_chat"
	CallbackTypeChangeExternalTag     string = "change_external_tag"
	CallbackTypeBatchJobResult        string = "batch_job_result"
)

const (
	ChangeContactCreateUser  string = "create_user"
	ChangeContactUpdateUser  string = "update_user"
	ChangeContactDeleteUser  string = "delete_user"
	ChangeContactCreateParty string = "create_party"
	ChangeContactUpdateParty string = "update_party"
	ChangeContactDeleteParty string = "delete_party"
	ChangeContactUpdateTag   string = "update_tag"
)

type KfServiceState int //客服会话状态

const (
	KfServiceStateUnSettled            KfServiceState = 0 //未处理
	KfServiceStateSettledByAI          KfServiceState = 1 //由智能助手接待
	KfServiceStateStaged               KfServiceState = 2 //待接入池排队中
	KfServiceStateSettledByServiceUser KfServiceState = 3 //由人工接待
	KfServiceStatePreStart             KfServiceState = 4 //已结束/未开始
)
