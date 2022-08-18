package wxthird

import (
	"fmt"

	"github.com/glutwins/workwx/wxcommon"
)

type SuiteCorpClient struct {
	CorpId     string
	CorpSecret string
	SuiteClient
}

type GetCorpTokenResp struct {
	wxcommon.CommonResp
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (scc *SuiteCorpClient) GetCorpToken() (string, error) {
	token, err := scc.TokenStore.GetSuiteCorpAccessToken(scc.SuiteId, scc.CorpId)
	if err != nil {
		return "", err
	}
	if token == "" {
		// TODO: lock and reget from cache
		suiteToken, err := scc.GetSuiteToken()
		if err != nil {
			return "", err
		}
		resp := &GetCorpTokenResp{}
		if err := scc.postJSON(fmt.Sprintf("/service/get_corp_token?suite_access_token=%s", suiteToken), map[string]interface{}{
			"auth_corpid":    scc.CorpId,
			"permanent_code": scc.CorpSecret,
		}, resp); err != nil {
			return "", err
		}
		scc.TokenStore.SetSuiteCorpAccessToken(scc.SuiteId, scc.CorpId, resp.AccessToken, resp.ExpiresIn)
		return resp.AccessToken, nil
	}
	return token, nil
}
