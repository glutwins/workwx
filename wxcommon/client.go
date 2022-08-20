package wxcommon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const wxBaseURL = "https://qyapi.weixin.qq.com/cgi-bin"

type TokenHandler func() (string, error)

type WorkClient struct {
	TokenHandler TokenHandler
}

func (wx *WorkClient) GetJSON(api string, resp WorkWxResp) error {
	r, err := http.Get(wxBaseURL + api)
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("wxbizhttp:%d(%s)", r.StatusCode, r.Status)
	}

	if err = json.NewDecoder(r.Body).Decode(resp); err != nil {
		return err
	}

	return resp.Err()
}

func (wx *WorkClient) PostJSON(api string, req interface{}, resp WorkWxResp) error {
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	r, err := http.Post(wxBaseURL+api, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("wxbizhttp:%d(%s)", r.StatusCode, r.Status)
	}

	if err = json.NewDecoder(r.Body).Decode(resp); err != nil {
		return err
	}

	return resp.Err()
}