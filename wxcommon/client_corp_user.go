package wxcommon

import (
	"fmt"
)

func (scc *SuiteCorpClient) UserSimpleList(departmentId int) (*UserSimpleListResp, error) {
	resp := &UserSimpleListResp{}
	if err := scc.GetRespWithToken("/user/simplelist?access_token=%s&department_id=%d", resp, departmentId); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) UserGetUserinfo(code string) (*UserGetUserinfoResp, error) {
	token, err := scc.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp := &UserGetUserinfoResp{}
	if err := scc.GetJSON(fmt.Sprintf("/user/getuserinfo?access_token=%s&code=%s", token, code), resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) UserGet(userid string) (*UserGetResp, error) {
	token, err := scc.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp := &UserGetResp{}
	if err := scc.GetJSON(fmt.Sprintf("/user/get?access_token=%s&userid=%s", token, userid), resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) UserUpdate(req *UserUpdateReq) error {
	resp := &CommonResp{}
	if err := scc.PostRespWithToken("/user/update?access_token=%s", req, resp); err != nil {
		return err
	}
	return resp.Err()
}

func (scc *SuiteCorpClient) UserDelete(userid string) error {
	resp := &CommonResp{}
	if err := scc.GetRespWithToken("/user/delete?access_token=%s&userid=%s", resp, userid); err != nil {
		return err
	}
	return resp.Err()
}

func (scc *SuiteCorpClient) UserBatchDelete(userids []string) error {
	resp := &CommonResp{}
	if err := scc.PostRespWithToken("/user/delete?access_token=%s", map[string]interface{}{"useridlist": userids}, resp); err != nil {
		return err
	}
	return resp.Err()
}

func (scc *SuiteCorpClient) UserList(departmentId int) (*UserListResp, error) {
	resp := &UserListResp{}
	if err := scc.GetRespWithToken("/user/list?access_token=%s&department_id=%d", resp, departmentId); err != nil {
		return nil, err
	}
	return resp, nil
}
