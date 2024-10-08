package wxsuite

import (
	"github.com/glutwins/workwx/store"
	"github.com/glutwins/workwx/wxcommon"
)

type ProviderConfig struct {
	CorpId         string `json:"corpid"`
	ProviderSecret string `json:"provider_secret"`
}

type ProviderAccessTokenResp struct {
	wxcommon.CommonResp
	ProviderAccessToken string `json:"provider_access_token"`
	ExpiresIn           int    `json:"expires_in"`
}

type ProviderClient struct {
	config *ProviderConfig
	cache  store.TokenCache
	wxcommon.WorkClient
}

func NewProviderClient(config *ProviderConfig, cache store.TokenCache) *ProviderClient {
	return &ProviderClient{config: config, cache: cache}
}

func (c *ProviderClient) GetProviderToken() (string, error) {
	accessToken, err := c.cache.GetProviderAccessToken(c.config.CorpId)
	if err != nil {
		return accessToken, nil
	}
	var resp ProviderAccessTokenResp
	if err := c.PostJSON("/service/get_provider_token", c.config, &resp); err != nil {
		return "", err
	}

	if resp.Err() != nil {
		return "", resp.Err()
	}

	c.cache.SetProviderAccessToken(c.config.CorpId, resp.ProviderAccessToken, resp.ExpiresIn)
	return resp.ProviderAccessToken, nil
}

type CustomizedAuthUrlResp struct {
	wxcommon.CommonResp
	QrcodeUrl string `json:"qrcode_url"`
	ExpiresIn int    `json:"expires_in"`
}

func (c *ProviderClient) GetCustomizedAuthUrl(state string, templateIdList []string) (*CustomizedAuthUrlResp, error) {
	accessToken, err := c.GetProviderToken()
	if err != nil {
		return nil, err
	}

	var resp CustomizedAuthUrlResp
	if err := c.PostJSON("/service/get_customized_auth_url?provider_access_token="+accessToken, map[string]interface{}{
		"state":           state,
		"templateid_list": templateIdList,
	}, &resp); err != nil {
		return nil, err
	}

	return &resp, resp.Err()
}
