package wxthird

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/glutwins/workwx/wxcommon"
)

const wxBaseURL = "https://qyapi.weixin.qq.com/cgi-bin"

type SuiteConfig struct {
	SuiteId        string
	SuiteSecret    string
	Token          string
	EncodingAESKey string
}

type WorkClient struct {
}

func (wx *WorkClient) getJSON(api string, resp wxcommon.WorkWxResp) error {
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

func (wx *WorkClient) postJSON(api string, req interface{}, resp wxcommon.WorkWxResp) error {
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
