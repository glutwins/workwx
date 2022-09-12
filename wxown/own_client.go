package wxown

import (
	"context"
	"fmt"

	"github.com/glutwins/workwx/store"
	"github.com/glutwins/workwx/wxcommon"
)

type SuiteClient wxcommon.SuiteClient

func NewSuiteClient(suiteId string, tokenCache store.TokenCache) *SuiteClient {
	sc := &SuiteClient{}
	sc.SuiteId = suiteId
	sc.TokenStore = tokenCache
	sc.GetAccessToken = func() (string, error) { return "", nil }
	return sc
}

func (sc *SuiteClient) WithContext(c context.Context) *SuiteClient {
	nsc := *sc
	nsc.Context = c
	return &nsc
}

func (sc *SuiteClient) NewCorpClient(corpId string, corpSecret string, agentId int) *wxcommon.SuiteCorpClient {
	scc := &wxcommon.SuiteCorpClient{CorpId: corpId, CorpSecret: corpSecret, AgentId: agentId, SuiteClient: wxcommon.SuiteClient(*sc)}
	scc.GetAccessToken = func() (string, error) {
		token, err := scc.TokenStore.GetSuiteCorpAccessToken(scc.SuiteId, scc.CorpId)
		if err != nil {
			return "", err
		}
		if token == "" {
			resp := &wxcommon.GetCorpTokenResp{}
			if err := scc.GetJSON(fmt.Sprintf("/gettoken?corpid=%s&corpsecret=%s", corpId, corpSecret), resp); err != nil {
				return "", err
			}
			scc.TokenStore.SetSuiteCorpAccessToken(scc.SuiteId, scc.CorpId, resp.AccessToken, resp.ExpiresIn)
			return resp.AccessToken, nil
		}
		return token, nil
	}
	return scc
}
