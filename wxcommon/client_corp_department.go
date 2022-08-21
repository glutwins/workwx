package wxcommon

import "fmt"

type Department struct {
	ID               int32    `json:"id"`
	Name             string   `json:"name"`
	NameEn           string   `json:"name_en"`
	DepartmentLeader []string `json:"department_leader"`
	ParentId         int32    `json:"parentid"`
	Order            int32    `json:"order"`
}

type DepartmentListResp struct {
	CommonResp
	Department []*Department `json:"department"`
}

func (scc *SuiteCorpClient) DepartmentList(id int) (*DepartmentListResp, error) {
	token, err := scc.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp := &DepartmentListResp{}
	if err := scc.GetJSON(fmt.Sprintf("/department/list?access_token=%s&id=%d", token, id), resp); err != nil {
		return nil, err
	}

	return resp, nil
}
