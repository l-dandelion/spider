package spider

import (
	"github.com/l-dandelion/spider/core/request"
	"net/http"
	"sync"
)

type Context struct {
	Spider   *Spider          //对应规则
	Rule     string           //规则名
	Request  *request.Request //请求
	Response *http.Response   //响应
	Err      error            //错误
}

//上下文对象池子
var (
	contextPool = &sync.Pool{
		New: func() interface{} {
			return &Context{}
		},
	}
)

//从池子中取一个上下文对象
func GetContext(spider *Spider, request *request.Request) *Context {
	ctx := contextPool.Get().(*Context)
	ctx.Spider = spider
	ctx.Request = request
	return ctx
}

//回收
func PutContext(ctx *Context) {
	ctx.Spider = nil
	ctx.Request = nil
	ctx.Response = nil
	ctx.Err = nil
	contextPool.Put(ctx)
}

//设置响应
func (self *Context) SetResponse(response *http.Response) *Context {
	self.Response = response
	return self
}

//设置错误
func (self *Context) SetError(err error) {
	self.Err = err
}
