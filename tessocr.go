package dgtessocr

import (
	"errors"
	dgctx "github.com/darwinOrg/go-common/context"
	dgerr "github.com/darwinOrg/go-common/enums/error"
	dghttp "github.com/darwinOrg/go-httpclient"
	dglogger "github.com/darwinOrg/go-logger"
	"github.com/otiai10/gosseract/v2"
	"os"
)

var emptyImageError = errors.New("图片内容为空")

func OcrImageUrl(ctx *dgctx.DgContext, imageUrl string, languages ...string) (string, error) {
	imageBytes, err := dghttp.Client11.DoGet(ctx, imageUrl, nil, nil)
	if err != nil {
		dglogger.Errorf(ctx, "dghttp.Client11.DoGet error: %v", err)
		return "", err
	}
	if len(imageBytes) == 0 {
		dglogger.Errorf(ctx, "获取图片失败: %s", imageUrl)
		return "", dgerr.SYSTEM_ERROR
	}

	return OcrImageBytes(ctx, imageBytes, languages...)
}

func OcrImageFile(ctx *dgctx.DgContext, imageFile string, languages ...string) (string, error) {
	imageBytes, err := os.ReadFile(imageFile)
	if err != nil {
		dglogger.Errorf(ctx, "Failed to read image file: %s, error: %v", imageFile, err)
		return "", err
	}

	return OcrImageBytes(ctx, imageBytes, languages...)
}

func OcrImageBytes(ctx *dgctx.DgContext, imageBytes []byte, languages ...string) (string, error) {
	if len(imageBytes) == 0 {
		return "", emptyImageError
	}

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

	// 读取图像字节
	err = client.SetImageFromBytes(imageBytes)
	if err != nil {
		dglogger.Errorf(ctx, "Failed to read image bytes: %v", err)
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
