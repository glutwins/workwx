//go:build windows

package wxchat

/*
#cgo CFLAGS: -x c++ -fpermissive -I./
#cgo LDFLAGS: -Llib/windows -lWeWorkFinanceSdk_C
#include <stdlib.h>
#include "lib/windows/WeWorkFinanceSdk_C.h"
*/
import "C"
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"io"
	"unsafe"
)
// 注意这个地方与上面注释的地方不能有空行，并且不能使用括号如import ("C" "fmt")
// 需要安装g++，windows可用https://github.com/skeeto/w64devkit/releases

type WeWorkFinanceSdk struct {
	sdk     *C.WeWorkFinanceSdk_t
	proxy   *C.char
	auth    *C.char
	timeout C.int
}

func (sdk *WeWorkFinanceSdk) Destroy() {
	C.free(unsafe.Pointer(sdk.proxy))
	C.free(unsafe.Pointer(sdk.auth))

	C.DestroySdk(sdk.sdk)
}

func NewWeWorkFinanceSdk(corpId string, secret string, proxy string, auth string, timeout int) (*WeWorkFinanceSdk, error) {
	sdk := &WeWorkFinanceSdk{sdk: C.NewSdk(), proxy: C.CString(proxy), auth: C.CString(auth), timeout: C.int(timeout)}
	cCropId := C.CString(corpId)
	cSecret := C.CString(secret)
	defer C.free(unsafe.Pointer(cCropId))
	defer C.free(unsafe.Pointer(cSecret))
	errCode := int(C.Init(sdk.sdk, cCropId, cSecret))
	if errCode != 0 {
		return nil, NewSdkError(errCode)
	}
	return sdk, nil
}

func (sdk *WeWorkFinanceSdk) GetChatMsg(seq int64, limit uint) (*EncryptData, error) {
	var res C.Slice_t
	errCode := int(C.GetChatData(sdk.sdk, C.ulonglong(seq), C.uint(limit), sdk.proxy, sdk.auth, sdk.timeout, &res))
	if errCode != 0 {
		return nil, NewSdkError(errCode)
	}
	encryptData := new(EncryptData)
	data := C.GoBytes(unsafe.Pointer(res.buf), res.len)
	defer C.free(unsafe.Pointer(res.buf))
	err := json.Unmarshal(data, encryptData)
	return encryptData, err
}

func DecryptData(encryptMsg *EncryptDataChatData, privateKey string) (*PlainMsg, error) {
	block, _ := pem.Decode([]byte(privateKey))
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes) // 还原数据
	if err != nil {
		return nil, err
	}
	randomKey, _ := base64.StdEncoding.DecodeString(encryptMsg.EncryptRandomKey)
	key, err := rsa.DecryptPKCS1v15(rand.Reader, priv, randomKey)
	if err != nil {
		return nil, err
	}

	var msg C.Slice_t
	msgKey := C.CString(string(key))
	encryptChatMsg := C.CString(encryptMsg.EncryptChatMsg)
	errCode := int(C.DecryptData(msgKey, encryptChatMsg, &msg))

	C.free(unsafe.Pointer(msgKey))
	C.free(unsafe.Pointer(encryptChatMsg))
	defer C.free(unsafe.Pointer(msg.buf))

	if errCode != 0 {
		return nil, NewSdkError(errCode)
	}

	plainMsg := new(PlainMsg)
	data := C.GoBytes(unsafe.Pointer(msg.buf), msg.len)
	err = json.Unmarshal(data, plainMsg)
	return plainMsg, err
}

//GetMediaData 获取媒体文件
func (sdk *WeWorkFinanceSdk) GetMediaData(sdkFileId string, writer io.Writer) error {
	cSdkFileid := C.CString(sdkFileId)
	defer C.free(unsafe.Pointer(cSdkFileid))

	isFinish := 0
	indexbuf := C.CString("")
	defer func() {
		if indexbuf != nil {
			C.free(unsafe.Pointer(indexbuf))
			indexbuf = nil
		}
	}()
	var mediaData C.MediaData_t
	for isFinish == 0 {
		if errCode := int(C.GetMediaData(sdk.sdk, indexbuf, cSdkFileid, sdk.proxy, sdk.auth, sdk.timeout, &mediaData)); errCode != 0 {
			return NewSdkError(errCode)
		}

		data := C.GoBytes(unsafe.Pointer(mediaData.data), mediaData.data_len)
		C.free(unsafe.Pointer(mediaData.data))

		if _, err := writer.Write(data); err != nil {
			return err
		}

		isFinish = int(mediaData.is_finish)

		C.free(unsafe.Pointer(indexbuf))
		indexbuf = mediaData.outindexbuf
	}

	return nil
}

func trimHiddenCharacter(originStr string) string {
	srcRunes := []rune(originStr)
	dstRunes := make([]rune, 0, len(srcRunes))
	for _, c := range srcRunes {
		if c >= 0 && c <= 31 {
			continue
		}
		if c == 127 {
			continue
		}
		dstRunes = append(dstRunes, c)
	}
	return string(dstRunes)
}
