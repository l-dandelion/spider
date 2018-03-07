package downloader

import (
	"github.com/l-dandelion/spider/core/request"
	"github.com/l-dandelion/spider/core/spider"
	"github.com/l-dandelion/spider/util/pool"
)

var (
	requestQueue   chan *spider.Context
	responseQueue  chan *spider.Context
	queueCap       = 100
	maxConcurrency = 100
	downloadPool   pool.Pool
	downloader     Downloader
)

func init() {
	downloader = SurferDownloader
	requestQueue = make(chan *spider.Context, queueCap)
	responseQueue = make(chan *spider.Context, queueCap)
	downloadPool.Init(maxConcurrency)
	go run()

}

func run() {
	for {
		ctx := <-requestQueue
		downloadPool.Add()
		go func(ctx *spider.Context) {
			defer downloadPool.Done()
			ctx = downloader.Download(ctx)
			responseQueue <- ctx
		}(ctx)
	}
}

func AddRequest(sp *spider.Spider, req *request.Request) {
	ctx := spider.GetContext(sp, req)
	requestQueue <- ctx
}

func GetResponse() *spider.Context {
	return <-responseQueue
}
