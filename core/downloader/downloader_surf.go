package downloader

import (
	"errors"
	"net/http"

	"github.com/l-dandelion/spider/core/downloader/surfer"
	"github.com/l-dandelion/spider/core/request"
	"github.com/l-dandelion/spider/core/spider"
)

type Surfer struct {
	surf    surfer.Surfer
	phantom surfer.Surfer
}

var SurferDownloader = &Surfer{
	surf: surfer.New(),
	// phantom: surfer.NewPhantom(config.PHANTOMJS, config.PHANTOMJS_TEMP),
	phantom: surfer.NewPhantom("phantom.js", "/Users/zhangliwei/godev/src/github.com/l-dandelion/spider/temp"),
}

func (self *Surfer) Download(ctx *spider.Context) *spider.Context {

	var resp *http.Response
	var err error

	switch ctx.Request.GetDownloaderID() {
	case request.SURF_ID:
		resp, err = self.surf.Download(ctx.Request)

	case request.PHANTOM_ID:
		resp, err = self.phantom.Download(ctx.Request)
	}

	if resp.StatusCode >= 400 {
		err = errors.New("响应状态 " + resp.Status)
	}

	ctx.SetResponse(resp).SetError(err)

	return ctx
}
