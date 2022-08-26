package wxcommon

import (
	"bytes"
	"io"
	"os"
	"path"
)

type MediaToUpload struct {
	filename string
	filesize int64
	r        io.Reader
}

func NewMediaFromFile(filename string) (*MediaToUpload, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	return &MediaToUpload{filename: path.Base(filename), filesize: stat.Size(), r: f}, nil
}

func NewMediaFromBuffer(filename string, buf *bytes.Buffer) (*MediaToUpload, error) {
	return &MediaToUpload{filename: path.Base(filename), filesize: int64(buf.Len()), r: buf}, nil
}

type MediaUploadResp struct {
	CommonResp
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt int64  `json:"created_at"`
}

type MediaUploadImgResp struct {
	CommonResp
	URL string `json:"url"`
}
