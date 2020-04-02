package netease

import (
	"encoding/json"
	// "fmt"

	// "encoding/json"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/go-resty/resty"
	jsoniter "github.com/json-iterator/go"
)

var jsonTool = jsoniter.ConfigCompatibleWithStandardLibrary

// ImClient .
type ImClient struct {
	AppKey    string
	AppSecret string
	Nonce     string

	mutex  *sync.Mutex
	client *resty.Client
}

// CreateImClient  创建im客户端，proxy留空表示不使用代理
func CreateImClient(appKey, appSecret, httpProxy string) *ImClient {
	c := &ImClient{AppKey: appKey, AppSecret: appSecret, Nonce: RandStringBytesMaskImprSrc(64), mutex: new(sync.Mutex)}
	c.client = resty.New()
	if len(httpProxy) > 0 {
		c.client.SetProxy(httpProxy)
	}

	c.client.SetHeader("Accept", "application/json;charset=utf-8")
	c.client.SetHeader("Content-Type", "application/x-www-form-urlencoded;charset=utf-8;")
	c.client.SetHeader("AppKey", c.AppKey)
	c.client.SetHeader("Nonce", c.Nonce)

	return c
}

func (c *ImClient) setCommonHead(req *resty.Request) {
	c.mutex.Lock() // 多线程并发访问map导致panic
	defer c.mutex.Unlock()

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	req.SetHeader("CurTime", timeStamp)
	req.SetHeader("CheckSum", ShaHashToHexStringFromString(c.AppSecret+c.Nonce+timeStamp))
}

func (c *ImClient) post(url string, fromData map[string]string, jsonResultKey string) (info *json.RawMessage, err error) {
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(fromData)

	resp, err := client.Post(url)
	info, err = handleResp(resp, err, jsonResultKey)
	return
}

func handleResp(resp *resty.Response, respErr error, jsonResultKey string) (info *json.RawMessage, err error) {
	if respErr != nil {
		return nil, respErr
	}

	var jsonRes map[string]*json.RawMessage

	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes[jsonResultKey], nil
}
