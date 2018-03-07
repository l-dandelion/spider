package request

import (
	"net/http"
	"time"
)

type Request struct {
	Url           string
	Method        string        //请求方法
	Header        http.Header   //请求头
	EnableCookie  bool          //是否使用cookie
	PostData      string        //表单数据
	DialTimeout   time.Duration //创建连接超时时间
	ConnTimeout   time.Duration //连接状态超时时间
	TryTimes      int           //尝试次数
	RetryPause    time.Duration //重试时间间隔
	RedirectTimes int           //重定向次数
	Priority      int           //优先级
	Proxy         string        //代理
	Unique        string        //标识
	DownloaderID  int           //下载器id
}

//下载器id编号
const (
	SURF_ID = iota
	PHANTOM_ID
)

func (self *Request) SetUrl(url string) *Request {
	self.Url = url
	return self
}

func (self *Request) GetUrl() string {
	return self.Url
}

func (self *Request) SetMethod(method string) *Request {
	self.Method = method
	return self
}

func (self *Request) GetMethod() string {
	return self.Method
}

func (self *Request) SetHeader(key, val string) *Request {
	self.SetHeader(key, val)
	return self
}

func (self *Request) SetHeaders(header http.Header) *Request {
	self.Header = header
	return self
}

func (self *Request) GetHeader() http.Header {
	return self.Header
}

func (self *Request) SetCookie(cookie string) *Request {
	self.Header.Set("cookie", cookie)
	return self
}

func (self *Request) SetEnableCookie(enable bool) *Request {
	self.EnableCookie = enable
	return self
}

func (self *Request) GetEnableCookie() bool {
	return self.EnableCookie
}

func (self *Request) SetPostData(postData string) *Request {
	self.PostData = postData
	return self
}

func (self *Request) GetPostData() string {
	return self.PostData
}

func (self *Request) SetDialTimeout(dialTimeout time.Duration) *Request {
	self.DialTimeout = dialTimeout
	return self
}

func (self *Request) GetDialTimeout() time.Duration {
	return self.DialTimeout
}

func (self *Request) SetConnTimeout(connTimeout time.Duration) *Request {
	self.ConnTimeout = connTimeout
	return self
}

func (self *Request) GetConnTimeout() time.Duration {
	return self.ConnTimeout
}

func (self *Request) SetTryTimes(tryTimes int) *Request {
	self.TryTimes = tryTimes
	return self
}

func (self *Request) GetTryTimes() int {
	return self.TryTimes
}

func (self *Request) SetRetryPause(retryPause time.Duration) *Request {
	self.RetryPause = retryPause
	return self
}

func (self *Request) GetRetryPause() time.Duration {
	return self.RetryPause
}

func (self *Request) SetRedirectTimes(redirectTimes int) *Request {
	self.RedirectTimes = redirectTimes
	return self
}

func (self *Request) GetRedirectTimes() int {
	return self.RedirectTimes
}

func (self *Request) SetPriority(priority int) *Request {
	self.Priority = priority
	return self
}

func (self *Request) GetPriority() int {
	return self.Priority
}

func (self *Request) SetProxy(proxy string) *Request {
	self.Proxy = proxy
	return self
}

func (self *Request) GetProxy() string {
	return self.Proxy
}

func (self *Request) GetUnique() string {
	return self.Unique
}

func (self *Request) SetDownloaderID(id int) *Request {
	self.DownloaderID = id
	return self
}

func (self *Request) GetDownloaderID() int {
	return self.DownloaderID
}
