package dgtessocr

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/otiai10/gosseract/v2"
)

func Ocr(ctx *dgctx.DgContext, imgFile string, languages ...string) (string, error) {
	// 创建一个客户端实例
	client := gosseract.NewClient()

	// 关闭客户端释放资源
	defer func(client *gosseract.Client) {
		err := client.Close()
		if err != nil {
			dglogger.Errorf(ctx, "Failed to close client: %v", err)
		}
	}(client)

	// 设置tesseract的语言（例如英语）
	err := client.SetLanguage(languages...)
	if err != nil {
		dglogger.Errorf(ctx, "Failed to set language: %v", err)
		return "", err
	}

	// 读取图像文件
	err = client.SetImage(imgFile)
	if err != nil {
		dglogger.Errorf(ctx, "Failed to read image file: %v", err)
		return "", err
	}

	// 执行OCR并获取识别出的文本
	text, err := client.Text()
	if err != nil {
		dglogger.Errorf(ctx, "Failed to perform OCR: %v", err)
		return "", err
	}

	return text, nil
}
