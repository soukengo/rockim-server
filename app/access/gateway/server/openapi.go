package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "rockimserver/apis/rockim/api/openapi/v1"
	"rockimserver/app/access/gateway/module/openapi/biz"
	"rockimserver/app/access/gateway/module/openapi/service"
	"rockimserver/pkg/util/encrypt"
)

const (
	openAPIHeaderProductId = "RockIM-OpenAPI-ProductId"
	openAPIHeaderSign      = "RockIM-OpenAPI-Sign"
	openAPIHeaderTimestamp = "RockIM-OpenAPI-Timestamp"
	openAPIHeaderNonce     = "RockIM-OpenAPI-Nonce"
)

type OpenApiServiceGroup struct {
	productUc   *biz.ProductUseCase
	userSrv     *service.UserService
	authSrv     *service.AuthService
	chatRoomSrv *service.ChatRoomService
}

func NewOpenApiServiceGroup(productUc *biz.ProductUseCase, userSrv *service.UserService, authSrv *service.AuthService, chatRoomSrv *service.ChatRoomService) *OpenApiServiceGroup {
	return &OpenApiServiceGroup{productUc: productUc, userSrv: userSrv, authSrv: authSrv, chatRoomSrv: chatRoomSrv}
}

func (g *OpenApiServiceGroup) Register(srv *http.Server) {
	srv.Use("/rockim.api.openapi.v1.*", g.checkSign(), g.setAPIRequest(), validate.Validator())
	v1.RegisterAuthAPIHTTPServer(srv, g.authSrv)
	v1.RegisterUserAPIHTTPServer(srv, g.userSrv)
	v1.RegisterChatRoomAPIHTTPServer(srv, g.chatRoomSrv)
}

type OpenAPIRequest interface {
	SetBase(request *v1.APIRequest)
	GetBase() *v1.APIRequest
}

// setAPIRequest 设置APIRequest
func (g *OpenApiServiceGroup) setAPIRequest() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (resp any, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				err = biz.ErrSignInvalid.NewWithMessage("缺少签名参数")
				return
			}
			r, ok := req.(OpenAPIRequest)
			if ok {
				if r.GetBase() == nil {
					r.SetBase(&v1.APIRequest{})
				}
				if len(r.GetBase().ProductId) == 0 {
					r.GetBase().ProductId = tr.RequestHeader().Get(openAPIHeaderProductId)
				}
			}
			return handler(ctx, req)
		}
	}
}

// checkSign 接口签名验证
func (g *OpenApiServiceGroup) checkSign() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				err = biz.ErrSignInvalid.NewWithMessage("缺少签名参数")
				return
			}
			productId := tr.RequestHeader().Get(openAPIHeaderProductId)
			nonce := tr.RequestHeader().Get(openAPIHeaderNonce)
			timestamp := tr.RequestHeader().Get(openAPIHeaderTimestamp)
			sign := tr.RequestHeader().Get(openAPIHeaderSign)
			if len(productId) == 0 || len(nonce) == 0 || len(timestamp) == 0 || len(sign) == 0 {
				err = biz.ErrSignInvalid.NewWithMessage("缺少签名参数")
				return
			}
			params := map[string]string{
				"productId": productId,
				"nonce":     nonce,
				"timestamp": timestamp,
			}
			p, err := g.productUc.FindById(ctx, productId)
			if err != nil {
				return
			}
			if !encrypt.CheckSign(params, p.ProductKey, sign) {
				err = biz.ErrSignInvalid.NewWithMessage("签名错误")
				return
			}
			return handler(ctx, req)
		}
	}
}
