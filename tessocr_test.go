package dgtessocr_test

import (
	dgctx "github.com/darwinOrg/go-common/context"
	dglogger "github.com/darwinOrg/go-logger"
	dgtessocr "github.com/darwinOrg/go-tessocr"
	"testing"
)

func TestOcr(t *testing.T) {
	ctx := &dgctx.DgContext{TraceId: "123"}
	text, err := dgtessocr.OcrImageFile(ctx, "test.jpg", "chi_sim")
	if err != nil {
		panic(err)
	}

	dglogger.Infof(ctx, "ocr text: %s", text)
}
