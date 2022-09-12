package wxcommon

import (
	"context"

	"github.com/glutwins/workwx/store"
)

type SuiteClient struct {
	SuiteId    string
	TokenStore store.TokenCache
	WorkClient
}

func (sc *SuiteClient) SuiteClientWithContext(c context.Context) *SuiteClient {
	var nsc = *sc
	nsc.Context = c
	return &nsc
}
