package wxcommon

import "github.com/glutwins/workwx/store"

type SuiteClient struct {
	SuiteId    string
	TokenStore store.TokenCache
	WorkClient
}
