package wxthird

import (
	"fmt"

	"github.com/glutwins/workwx/store"
	"github.com/glutwins/workwx/wxcommon"
)

type GetSuiteTokenResp struct {
	wxcommon.CommonResp
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int    `json:"expires_in"`
}

type SuiteClient struct {
	SuiteId     string
	SuiteSecret string
	TokenStore  store.TokenCache
	WorkClient
}

func (sc *SuiteClient) GetSuiteToken() (string, error) {
	token, err := sc.TokenStore.GetSuiteAccessToken(sc.SuiteId)
	if err != nil {
		return "", err
	}
	if token == "" {
		// TODO: lock and reget from cach
		ticket, err := sc.TokenStore.GetSuiteTicket(sc.SuiteId)
		if err != nil {
			return "", err
		}
		resp := &GetSuiteTokenResp{}
		if err := sc.postJSON("/service/get_suite_token", map[string]interface{}{
			"suite_id":     sc.SuiteId,
			"suite_secret": sc.SuiteSecret,
			"suite_ticket": ticket,
		}, resp); err != nil {
			return "", err
		}
		sc.TokenStore.SetSuiteAccessToken(sc.SuiteId, resp.SuiteAccessToken, resp.ExpiresIn)
		return resp.SuiteAccessToken, nil
	}
	return token, nil
}

type GetPreAuthCodeResp struct {
	wxcommon.CommonResp
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}

func (sc *SuiteClient) GetPreAuthCode() (*GetPreAuthCodeResp, error) {
	token, err := sc.GetSuiteToken()
	if err != nil {
		return nil, err
	}
	resp := &GetPreAuthCodeResp{}
	if err := sc.getJSON(fmt.Sprintf("/service/get_pre_auth_code?suite_access_token=%s", token), resp); err != nil {
		return nil, err
	}
	return resp, err
}

func (sc *SuiteClient) NewCorpClient(corpId string, corpSecret string) *SuiteCorpClient {
	return &SuiteCorpClient{
		CorpId:      corpId,
		CorpSecret:  corpSecret,
		SuiteClient: *sc,
	}
}
