package http

import (
	"context"
	"github.com/go-resty/resty/v2"
	"github.com/golang/protobuf/proto"
	"net/http"
	"rockim/pkg/errors"
	log "rockim/pkg/log"
	"rockim/pkg/util/json"
	"rockim/pkg/util/strings"
	"time"
)

var (
	client     = resty.New().SetTimeout(time.Second * 30)
	loggerName = "http"
)

// Post send POST request
func Post(ctx context.Context, addr string, params map[string]string) (res []byte, err error) {
	start := time.Now().UnixMilli()
	reqId := strings.RandStr(20)
	logger := log.Use(loggerName).WithContext(ctx)
	logger.Infof("Post[%s] request addr: %s, params: %v", reqId, addr, params)
	resp, err := newRequest().SetFormData(params).Post(addr)
	if err != nil {
		return
	}
	res = resp.Body()
	if resp.StatusCode() != http.StatusOK {
		err = errors.NewCode(resp.StatusCode())
	}
	end := time.Now().UnixMilli()
	logger.Infof("Post[%s] cost : %v", reqId, end-start)
	return
}

// PostJSON send POST request with JSON body
func PostJSON(ctx context.Context, addr string, data string) (res []byte, err error) {
	return PostJSONWithHeader(ctx, addr, data, nil)
}

// PostJSONWithHeader send POST request with headers
func PostJSONWithHeader(ctx context.Context, addr string, data string, headers map[string]string) (res []byte, err error) {
	if len(headers) == 0 {
		headers = make(map[string]string)
	}
	start := time.Now().UnixMilli()
	headers["content-type"] = "application/json"
	headers["accept"] = "application/json"
	logger := log.Use(loggerName).WithContext(ctx)
	reqId := strings.RandStr(20)
	logger.Infof("PostJSONWithHeader[%s] request addr: %s, data: %v", reqId, addr, data)
	resp, err := newRequest().SetHeaders(headers).SetBody(data).Post(addr)
	if err != nil {
		return
	}
	res = resp.Body()
	if resp.StatusCode() != http.StatusOK {
		err = errors.NewCode(resp.StatusCode())
	}
	end := time.Now().UnixMilli()
	logger.Infof("PostJSONWithHeader[%s] cost : %v", reqId, end-start)
	return
}

// PostProtobufWithHeader send Protobuf request with headers
func PostProtobufWithHeader(ctx context.Context, addr string, data proto.Message, headers map[string]string) (res []byte, err error) {
	if len(headers) == 0 {
		headers = make(map[string]string)
	}
	start := time.Now().UnixMilli()
	headers["content-type"] = "application/x-protobuf"
	headers["Accept"] = "application/x-protobuf"
	logger := log.Use(loggerName).WithContext(ctx)
	reqId := strings.RandStr(20)
	logger.Infof("PostProtobufWithHeader[%s] request addr: %s, data: %v", reqId, addr, json.TryToJSONString(data))
	body, err := proto.Marshal(data)
	if err != nil {
		return
	}
	req := newRequest().SetHeaders(headers).SetBody(body)
	resp, err := req.Post(addr)
	if err != nil {
		return
	}
	res = resp.Body()
	if resp.StatusCode() != http.StatusOK {
		err = errors.NewCode(resp.StatusCode())
	}
	end := time.Now().UnixMilli()
	logger.Infof("PostProtobufWithHeader[%s] cost : %v", reqId, end-start)
	return
}

// Get send GET REQUEST
func Get(ctx context.Context, addr string) (res []byte, err error) {
	start := time.Now().UnixMilli()
	logger := log.Use(loggerName).WithContext(ctx)
	reqId := strings.RandStr(20)
	logger.Infof("Get[%s] request addr: %s", reqId, addr)
	resp, err := newRequest().Get(addr)
	if err != nil {
		return
	}
	res = resp.Body()
	if resp.StatusCode() != http.StatusOK {
		err = errors.NewCode(resp.StatusCode())
	}
	end := time.Now().UnixMilli()
	logger.Infof("Get[%s] cost : %v", reqId, end-start)
	return
}

func newRequest() *resty.Request {
	return client.R().EnableTrace()
}
