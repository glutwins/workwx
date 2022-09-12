package wxcommon

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

const wxBaseURL = "https://qyapi.weixin.qq.com/cgi-bin"

type TokenHandler func() (string, error)

type WorkClient struct {
	GetAccessToken TokenHandler
	Context        context.Context
	Logger         ClientLogger
}

func (wx *WorkClient) context() context.Context {
	if wx.Context != nil {
		return wx.Context
	}
	return context.Background()
}

func (wx *WorkClient) GetJSON(api string, resp WorkWxResp) error {
	var err error
	defer func() {
		if wx.Logger != nil {
			wx.Logger.Println(wx.context(), wxBaseURL+api, nil, resp, err)
		}
	}()
	var r *http.Response
	r, err = http.Get(wxBaseURL + api)
	if err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		err = fmt.Errorf("wxbizhttp:%d(%s)", r.StatusCode, r.Status)
		return err
	}

	if err = json.NewDecoder(r.Body).Decode(resp); err != nil {
		return err
	}

	return resp.Err()
}

func (wx *WorkClient) GetRespWithToken(api string, resp WorkWxResp, args ...interface{}) error {
	token, err := wx.GetAccessToken()
	if err != nil {
		return err
	}

	var argsWithToken = []interface{}{token}
	argsWithToken = append(argsWithToken, args...)

	return wx.GetJSON(fmt.Sprintf(api, argsWithToken...), resp)
}

func (wx *WorkClient) PostJSON(api string, req interface{}, resp WorkWxResp) error {
	var err error
	defer func() {
		if wx.Logger != nil {
			wx.Logger.Println(wx.context(), wxBaseURL+api, nil, resp, err)
		}
	}()

	var b []byte
	if b, err = json.Marshal(req); err != nil {
		return err
	}

	var r *http.Response
	if r, err = http.Post(wxBaseURL+api, "application/json", bytes.NewBuffer(b)); err != nil {
		return err
	}

	if r.StatusCode != http.StatusOK {
		err = fmt.Errorf("wxbizhttp:%d(%s)", r.StatusCode, r.Status)
		return err
	}

	if err = json.NewDecoder(r.Body).Decode(resp); err != nil {
		return err
	}

	return resp.Err()
}

func (wx *WorkClient) PostMedia(api string, media *MediaToUpload, resp WorkWxResp) error {
	buf := bytes.NewBuffer(nil)
	mv := multipart.NewWriter(buf)
	wr, err := mv.CreateFormFile("media", media.filename)
	if err != nil {
		return err
	}

	_, err = io.Copy(wr, media.r)
	if err != nil {
		return err
	}

	r, err := http.Post(wxBaseURL+api, mv.FormDataContentType(), buf)
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

func (wx *WorkClient) PostRespWithToken(api string, req interface{}, resp WorkWxResp, args ...interface{}) error {
	token, err := wx.GetAccessToken()
	if err != nil {
		return err
	}

	var argsWithToken = []interface{}{token}
	argsWithToken = append(argsWithToken, args...)

	return wx.PostJSON(fmt.Sprintf(api, argsWithToken...), req, resp)
}
