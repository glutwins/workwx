package wxcommon

import "context"

type ClientLogger interface {
	Println(c context.Context, api string, req interface{}, resp WorkWxResp, err error)
}
