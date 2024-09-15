package qrgen

import (
	"bytes"
	"image"
	"image/png"
)

// CreateImage 画像データとして操作可能なImageオブジェクトに変換する
func CreateImage(b []byte) (image.Image, error) {
	return png.Decode(bytes.NewReader(b))
}
