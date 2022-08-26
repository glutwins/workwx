package wxcommon

import (
	"fmt"
)

// MediaUpload 上传临时素材
// 参考 https://developer.work.weixin.qq.com/document/path/90253
// 接口不校验文件大小，请按照文档说明限制上传内容大小
func (scc *SuiteCorpClient) MediaUpload(mediaType string, media *MediaToUpload) (*MediaUploadResp, error) {
	token, err := scc.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp := &MediaUploadResp{}
	if err := scc.PostMedia(fmt.Sprintf("/media/upload?access_token=%s&type=%s", token, mediaType), media, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (scc *SuiteCorpClient) MediaUploadImg(mediaType string, media *MediaToUpload) (*MediaUploadImgResp, error) {
	token, err := scc.GetAccessToken()
	if err != nil {
		return nil, err
	}
	resp := &MediaUploadImgResp{}
	if err := scc.PostMedia(fmt.Sprintf("/media/uploadimg?access_token=%s", token), media, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
