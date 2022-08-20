package wxcommon

import (
	"fmt"
)

type UserCreateReq struct {
	UserSimple
	Alias    string `json:"alias"`
	Mobile   string `json:"mobile"`
	Order    []int  `json:"order"`
	Position string `json:"position"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
}

type UserSimple struct {
	UserId     string `json:"userid"`
	Name       string `json:"name"`
	Department []int  `json:"department"`
	OpenUserId string `json:"open_userid"`
}

type UserSimpleListResp struct {
	CommonResp
	UserList []*UserSimple `json:"userlist"`
}

func (scc *SuiteCorpClient) UserSimpleList(departmentId int) (*UserSimpleListResp, error) {
	resp := &UserSimpleListResp{}
	token, err := scc.TokenHandler()
	if err != nil {
		return nil, err
	}
	if err := scc.GetJSON(fmt.Sprintf("/user/simplelist?access_token=%s&department_id=%d", token, departmentId), resp); err != nil {
		return nil, err
	}

	return resp, err
}

type UserGetUserinfoResp struct {
	CommonResp
	UserId   string
	DeviceId string
}

func (scc *SuiteCorpClient) UserGetUserinfo(code string) (*UserGetUserinfoResp, error) {
	token, err := scc.TokenHandler()
	if err != nil {
		return nil, err
	}
	resp := &UserGetUserinfoResp{}
	if err := scc.GetJSON(fmt.Sprintf("/user/getuserinfo?access_token=%s&code=%s", token, code), resp); err != nil {
		return nil, err
	}

	return resp, nil
}

type UserGetResp struct {
	CommonResp
	UserId           string      `json:"userid"`
	Name             string      `json:"name"`
	Mobile           string      `json:"mobile"`
	Department       []int32     `json:"department"`
	Order            []int32     `json:"order"`
	Position         string      `json:"position"`
	Gender           string      `json:"gender"`
	Email            string      `json:"email"`
	BizMail          string      `json:"biz_mail"`
	IsLeaderInDept   []int8      `json:"is_leader_in_dept"`
	DirectLeader     []string    `json:"direct_leader"`
	Avatar           string      `json:"avatar"`
	ThumbAvatar      string      `json:"thumb_avatar"`
	Telephone        string      `json:"telephone"`
	Alias            string      `json:"alias"`
	Address          string      `json:"address"`
	OpenUserID       string      `json:"open_userid"`
	MainDepartment   int         `json:"main_department"`
	ExternalPosition string      `json:"external_position"`
	ExternalProfile  interface{} `json:"external_profile"`
}

func (scc *SuiteCorpClient) UserGet(userid string) (*UserGetResp, error) {
	token, err := scc.TokenHandler()
	if err != nil {
		return nil, err
	}
	resp := &UserGetResp{}
	if err := scc.GetJSON(fmt.Sprintf("/user/get?access_token=%s&userid=%s", token, userid), resp); err != nil {
		return nil, err
	}

	return resp, nil
}
