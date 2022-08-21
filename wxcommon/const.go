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
