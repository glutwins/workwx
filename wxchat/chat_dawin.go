//go:build darwin

package wxchat

import "io"

type WeWorkFinanceSdk struct {
}

func (sdk *WeWorkFinanceSdk) Destroy() {
}

func NewWeWorkFinanceSdk(corpId string, secret string, proxy string, auth string, timeout int) (*WeWorkFinanceSdk, error) {
	return &WeWorkFinanceSdk{}, nil
}

func (sdk *WeWorkFinanceSdk) GetChatMsg(seq int64, limit uint) (*EncryptData, error) {
	return nil, nil
}

func DecryptData(encryptMsg *EncryptDataChatData, privateKey string) (*PlainMsg, error) {
	return nil, nil
}

//GetMediaData 获取媒体文件
func (sdk *WeWorkFinanceSdk) GetMediaData(sdkFileId string, writer io.Writer) error {
	return nil
}
