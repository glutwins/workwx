package wxcommon

import "fmt"

type CommonResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (r *CommonResp) Error() string {
	return fmt.Sprintf("workwx-err:%d(%s)", r.ErrCode, r.ErrMsg)
}

func (r *CommonResp) Err() error {
	if r.ErrCode != 0 {
		return r
	}
	return nil
}

type WorkWxResp interface {
	Err() error
}
