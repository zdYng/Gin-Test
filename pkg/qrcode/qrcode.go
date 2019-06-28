package qrcode

import (
	"github.com/boombuler/barcode/qr"
	"gin-blog/pkg/setting"
)

type QrCode struct{
	URL string
	Width int
	Height int
	Ext string
	Level qr.ErrorCorrectionLevel
	Mode qr.Encoding
}

const (
	EXT_JPG = ".jpg"
)

func NewQrCode(url string, width int,height int,level qr.ErrorCorrectionLevel, mode qr.Encoding)*QrCode{
	return &QrCode{
		URL: url,
		Width: width,
		Height: height,
		Level: level,
		Mode: mode,
		Ext: EXT_JPG,
	}
}

func GetQrCodePath() string{
	return setting.AppSetting.QrCodeSave
}