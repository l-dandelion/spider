package downloader

import (
	"github.com/l-dandelion/spider/core/spider"
)

// The Downloader interface.
// You can implement the interface by implement function Download.
// Function Download need to return Page instance pointer that has request result downloaded from Request.
type Downloader interface {
	Download(*spider.Context) *spider.Context
}
