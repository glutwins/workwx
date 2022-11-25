package wxsuite

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/glutwins/workwx/store"
	"github.com/glutwins/workwx/wxcommon"
)

type GetSuiteTokenResp struct {
	wxcommon.CommonResp
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int    `json:"expires_in"`
}

type SuiteClient struct {
	wxcommon.SuiteClient
	SuiteSecret string
}

func NewSuiteClient(suiteId string, suiteSecret string, tokenCache store.TokenCache) *SuiteClient {
	sc := &SuiteClient{SuiteSecret: suiteSecret}
	sc.SuiteId = suiteId
	sc.TokenStore = tokenCache
	sc.GetAccessToken = sc.GetSuiteToken
	return sc
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
		if err := sc.PostJSON("/service/get_suite_token", map[string]interface{}{
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
	if err := sc.GetJSON(fmt.Sprintf("/service/get_pre_auth_code?suite_access_token=%s", token), resp); err != nil {
		return nil, err
	}
	return resp, err
}

type CorpSimple struct {
	CorpId   string `json:"corpid"`    // 授权方企业微信id
	CorpName string `json:"corp_name"` // 授权方企业名称，即企业简称
}

type CorpDetail struct {
	CorpSimple
	CorpType          string `json:"corp_type"`            // 授权方企业类型，认证号：verified, 注册号：unverified
	CorpSquareLogoURL string `json:"corp_square_logo_url"` // 授权方企业方形头像
	CorpUserMax       int    `json:"corp_user_max"`        // 授权方企业用户规模
	CorpFullName      string `json:"corp_full_name"`       // 授权方企业的主体名称(仅认证或验证过的企业有)，即企业全称。企业微信将逐步回收该字段，后续实际返回内容为企业名称，即auth_corp_info.corp_name
	VerifiedEndTime   int64  `json:"verified_end_time"`    // 认证到期时间
	SubjectType       int    `json:"subject_type"`         // 企业类型，1. 企业; 2. 政府以及事业单位; 3. 其他组织, 4.团队号
	CorpWxqrcode      string `json:"corp_wxqrcode"`        // 授权企业在微信插件（原企业号）的二维码，可用于关注微信插件
	CorpScale         string `json:"corp_scale"`           // 企业规模。当企业未设置该属性时，值为空
	CorpIndustry      string `json:"corp_industry"`        // 企业所属行业。当企业未设置该属性时，值为空
	CorpSubIndustry   string `json:"corp_sub_industry"`    // 企业所属子行业。当企业未设置该属性时，值为空
}

func (c CorpDetail) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *CorpDetail) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), c)
}

type SharedFrom struct {
	CorpId    string `json:"corpid"`
	ShareType int    `json:"share_type"`
}

type AgentPrivilege struct {
	Level      int      `json:"level"`
	AllowParty []int    `json:"allow_party"`
	AllowUser  []string `json:"allow_user"`
	AllowTag   []int    `json:"allow_tag"`
	ExtraParty []int    `json:"extra_party"`
	ExtraUser  []string `json:"extra_user"`
	ExtraTag   []int    `json:"extra_tag"`
}

type Agent struct {
	AgentId          int            `json:"agentid"`         // 授权方应用id
	Name             string         `json:"name"`            // 授权方应用名字
	RoundLogoURL     string         `json:"round_logo_url"`  // 授权方应用方形头像
	SquareLogoURL    string         `json:"square_logo_url"` // 授权方应用圆形头像
	AppId            int            `json:"appid"`
	AuthMode         int            `json:"auth_mode"`          // 授权模式，0为管理员授权；1为成员授权
	IsCustomizedApp  bool           `json:"is_customized_app"`  // 是否为代开发自建应用
	AuthFromThirdapp bool           `json:"auth_from_thirdapp"` // 来自第三方应用接口唤起,仅通过第三方应用添加自建应用 获取授权链接授权代开发自建应用时，才返回该字段
	Privilege        AgentPrivilege `json:"privilege"`          // 应用对应的权限
	SharedFrom       SharedFrom     `json:"shared_from"`        // 共享了应用的企业信息，仅当由企业互联或者上下游共享应用触发的安装时才返回
}

type AuthInfo struct {
	Agent []Agent `json:"agent"` // 授权的应用信息，注意是一个数组，但仅旧的多应用套件授权时会返回多个agent，对新的单应用授权，永远只返回一个agent
}

type GetPermanentCodeResp struct {
	wxcommon.CommonResp
	AccessToken    string     `json:"access_token"`     // 授权方（企业）access_token
	ExpiresIn      int        `json:"expires_in"`       // 授权方（企业）access_token超时时间（秒）
	PermanentCode  string     `json:"permanent_code"`   // 企业微信永久授权码,最长为512字节
	DealerCorpInfo CorpSimple `json:"dealer_corp_info"` // 代理服务商企业信息。应用被代理后才有该信息
	AuthCorpInfo   CorpDetail `json:"auth_corp_info"`   // 授权方企业信息
	AuthInfo       AuthInfo   `json:"auth_info"`        // 授权信息。如果是通讯录应用，且没开启实体应用，是没有该项的。通讯录应用拥有企业通讯录的全部信息读写权限
	State          string     `json:"state"`            // 安装应用时，扫码或者授权链接中带的state值
}

func (sc *SuiteClient) GetPermanentCode(authCode string) (*GetPermanentCodeResp, error) {
	resp := &GetPermanentCodeResp{}
	if err := sc.PostRespWithToken("/service/get_permanent_code?suite_access_token=%s", map[string]interface{}{
		"auth_code": authCode,
	}, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type GetAuthInfoResp struct {
	wxcommon.CommonResp
	AuthCorpInfo CorpDetail `json:"auth_corp_info"` // 授权方企业信息
	AuthInfo     AuthInfo   `json:"auth_info"`      // 授权信息。如果是通讯录应用，且没开启实体应用，是没有该项的。通讯录应用拥有企业通讯录的全部信息读写权限
}

func (sc *SuiteClient) GetAuthInfo(authCorpId, permanentCode string) (*GetAuthInfoResp, error) {
	resp := &GetAuthInfoResp{}
	if err := sc.PostRespWithToken("/service/get_auth_info?suite_access_token=%s", map[string]interface{}{
		"auth_corpid":    authCorpId,
		"permanent_code": permanentCode,
	}, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type GetUserinfo3rdResp struct {
	wxcommon.CommonResp
	CorpId     string
	UserId     string
	DeviceId   string
	UserTicket string `json:"user_ticket"`
	ExpiresIn  int    `json:"expires_in"`
	OpenUserID string `json:"open_userid"`
}

func (sc *SuiteClient) GetUserinfo3rd(code string) (*GetUserinfo3rdResp, error) {
	accessToken, err := sc.GetSuiteToken()
	if err != nil {
		return nil, err
	}
	resp := &GetUserinfo3rdResp{}
	var params = make(url.Values)
	params.Set("suite_access_token", accessToken)
	params.Set("code", code)
	if err := sc.GetJSON("/service/getuserinfo3rd?"+params.Encode(), resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (sc *SuiteClient) MiniprogramJsCode2Session(code string) (*wxcommon.MiniprogramJsCode2SessionResp, error) {
	accessToken, err := sc.GetSuiteToken()
	if err != nil {
		return nil, err
	}
	resp := &wxcommon.MiniprogramJsCode2SessionResp{}
	var params = make(url.Values)
	params.Set("suite_access_token", accessToken)
	params.Set("js_code", code)
	params.Set("grant_type", "authorization_code")
	if err := sc.GetJSON("/service/miniprogram/jscode2session?"+params.Encode(), resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (sc *SuiteClient) NewCorpClient(corpId string, corpSecret string, agentId int) *wxcommon.SuiteCorpClient {
	scc := &wxcommon.SuiteCorpClient{
		CorpId:      corpId,
		CorpSecret:  corpSecret,
		AgentId:     agentId,
		SuiteClient: sc.SuiteClient,
	}
	scc.GetAccessToken = func() (string, error) {
		token, err := scc.TokenStore.GetSuiteCorpAccessToken(scc.SuiteId, scc.CorpId)
		if err != nil {
			return "", err
		}
		if token == "" {
			// TODO: lock and reget from cache
			suiteToken, err := sc.GetSuiteToken()
			if err != nil {
				return "", err
			}
			resp := &wxcommon.GetCorpTokenResp{}
			if err := scc.PostJSON(fmt.Sprintf("/service/get_corp_token?suite_access_token=%s", suiteToken), map[string]interface{}{
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
	return scc
}

func (sc *SuiteClient) NewOwnCorpClient(corpId string, corpSecret string, agentId int) *wxcommon.SuiteCorpClient {
	scc := &wxcommon.SuiteCorpClient{
		CorpId:      corpId,
		CorpSecret:  corpSecret,
		AgentId:     agentId,
		SuiteClient: sc.SuiteClient,
	}
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
