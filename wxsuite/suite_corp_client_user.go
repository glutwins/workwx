package wxsuite

import (
	"fmt"

	"github.com/glutwins/workwx/wxcommon"
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
	wxcommon.CommonResp
	UserList []*UserSimple `json:"userlist"`
}

func (scc *SuiteCorpClient) UserSimpleList(departmentId int) (*UserSimpleListResp, error) {
	resp := &UserSimpleListResp{}
	token, err := scc.GetCorpToken()
	if err != nil {
		return nil, err
	}
	if err := scc.getJSON(fmt.Sprintf("/user/simplelist?access_token=%s&department_id=%d", token, departmentId), resp); err != nil {
		return nil, err
	}

	return resp, err
}
