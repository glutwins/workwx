//go:build linux

package wxchat

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -Llib/linux -lWeWorkFinanceSdk_C
#include <stdlib.h>
#include "lib/linux/WeWorkFinanceSdk_C.h"
*/
import "C"
import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"unsafe"
) // 注意这个地方与上面注释的地方不能有空行，并且不能使用括号如import ("C" "fmt")

//UploadPartSize 分片上传大小
const UploadPartSize = 2 * 1024 * 1024

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

/*

//GetMediaData 获取媒体文件
func GetMediaData(c ctx.Context, request *MediaDataRequest, writer iox.Writer, retryLimit int) error {
	sdk := C.NewSdk()
	cropID := C.CString(request.CorpID)
	secret := C.CString(request.Secret)
	defer C.DestroySdk(sdk)
	defer C.free(unsafe.Pointer(cropID))
	defer C.free(unsafe.Pointer(secret))

	errCode := int(C.Init(sdk, cropID, secret))
	if int(errCode) != 0 {
		return errorx.New(CodeMsg[errCode])
	}

	sdkFileid := C.CString(request.SdkFileid)
	proxy := C.CString(request.Proxy)
	auth := C.CString(request.Auth)
	timeout := C.int(request.Timeout)

	defer C.free(unsafe.Pointer(sdkFileid))
	defer C.free(unsafe.Pointer(proxy))
	defer C.free(unsafe.Pointer(auth))

	isFinish := 0
	indexbuf := C.CString("")
	buffer := bytes.Buffer{}
	var mediaData C.MediaData_t
	retryCount := 0
	for isFinish == 0 {
		errCode = int(C.GetMediaData(sdk, indexbuf, sdkFileid, proxy, auth, timeout, &mediaData))
		if errCode != 0 {
			log.Warn(c).Interface("req", request).Msgf("%s 下载错误 [%d]%s", request.SdkFileid, errCode, CodeMsg[errCode])
			time.Sleep(500 * time.Millisecond)
			if retryCount >= retryLimit {
				C.free(unsafe.Pointer(indexbuf))
				return errorx.New(CodeMsg[errCode])
			}
			retryCount++
			continue
		}

		data := C.GoBytes(unsafe.Pointer(mediaData.data), mediaData.data_len)
		C.free(unsafe.Pointer(mediaData.data))

		buffer.Write(data)
		if buffer.Len() > UploadPartSize {
			n, err := writer.Write(buffer.Bytes())
			if err != nil || n != buffer.Len() {
				log.Error(c).Interface("req", request).Msgf("媒体文件 %s 写入错误：文件大小[%d]，写入大小[%d]", request.SdkFileid, buffer.Len(), n)
				C.free(unsafe.Pointer(indexbuf))
				if err != nil {
					log.Error(c).Interface("req", request).Msgf("媒体文件 %s 写入错误：%s", err.Error())
					return err
				}
				return errorx.New(fmt.Sprintf("文件数据写入不完整。数据长度[%d]，写入长度[%d]", buffer.Len(), n))
			}
			log.Info(c).Msgf("媒体文件 %s 下载信息：文件大小[%d]，写入大小[%d]", request.SdkFileid, buffer.Len(), n)
			buffer.Reset()
		}

		C.free(unsafe.Pointer(indexbuf))

		isFinish = int(mediaData.is_finish)
		indexbuf = mediaData.outindexbuf
	}
	C.free(unsafe.Pointer(indexbuf))

	if buffer.Len() > 0 {
		n, err := writer.Write(buffer.Bytes())
		if err != nil || n != buffer.Len() {
			log.Debug(c).Msgf("媒体文件 %s 写入错误：文件大小[%d]，写入大小[%d]", request.SdkFileid, buffer.Len(), n)
			C.free(unsafe.Pointer(indexbuf))
			if err != nil {
				log.Debug(c).Msgf("媒体文件 %s 写入错误：%s", err.Error())
				return err
			}
			return errorx.New(fmt.Sprintf("文件数据写入不完整。数据长度[%d]，写入长度[%d]", buffer.Len(), n))
		}
		log.Debug(c).Msgf("媒体文件 %s 下载信息：文件大小[%d]，写入大小[%d]", request.SdkFileid, buffer.Len(), n)
		buffer.Reset()
	}

	return nil
}
*/

func TrimHiddenCharacter(originStr string) string {
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
