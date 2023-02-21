package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	v1 "rockimserver/apis/rockim/api/client/v1"
	"rockimserver/apis/rockim/shared/reasons"
	"rockimserver/app/access/gateway/module/client/biz"
	"rockimserver/app/access/gateway/module/client/biz/options"
	"rockimserver/app/access/gateway/module/client/service"
	"rockimserver/pkg/errors"
	"rockimserver/pkg/util/encrypt"
)

const (
	clientHeaderProductId   = "RockIM-Client-ProductId"
	clientHeaderSign        = "RockIM-Client-Sign"
	clientHeaderTimestamp   = "RockIM-Client-Timestamp"
	clientHeaderNonce       = "RockIM-Client-Nonce"
	clientHeaderAccessToken = "RockIM-Client-AccessToken"
)

var (
	ErrSignInvalid        = errors.BadRequest(reasons.OpenAPI_SIGN_INVALID.String(), "签名错误")
	ErrAccessTokenInvalid = errors.BadRequest(reasons.User_ACCESS_TOKEN_INVALID.String(), "访问令牌无效或已过期")
)

type ClientServiceGroup struct {
	productUc *biz.ProductUseCase
	authUc    *biz.AuthUseCase
	userSrv   *service.UserService
	authSrv   *service.AuthService
}

func NewClientServiceGroup(productUc *biz.ProductUseCase, authUc *biz.AuthUseCase, userSrv *service.UserService, authSrv *service.AuthService) *ClientServiceGroup {
	return &ClientServiceGroup{productUc: productUc, authUc: authUc, userSrv: userSrv, authSrv: authSrv}
}

func (g *ClientServiceGroup) Register(srv *http.Server) {
	srv.Use("/rockim.api.client.v1.*", g.checkSign(), g.checkAuth(), g.setAPIRequest(), validate.Validator())
	v1.RegisterAuthAPIHTTPServer(srv, g.authSrv)
	v1.RegisterUserAPIHTTPServer(srv, g.userSrv)
}

type ClientAPIRequest interface {
	SetBase(request *v1.APIRequest)
	GetBase() *v1.APIRequest
}

// setAPIRequest 设置APIRequest
func (g *ClientServiceGroup) setAPIRequest() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (resp any, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				err = ErrSignInvalid.NewWithMessage("缺少签名参数")
				return
			}
			r, ok := req.(ClientAPIRequest)
			if ok {
				if r.GetBase() == nil {
					r.SetBase(&v1.APIRequest{})
				}
				if len(r.GetBase().ProductId) == 0 {
					r.GetBase().ProductId = tr.RequestHeader().Get(clientHeaderProductId)
				}
			}
			return handler(ctx, req)
		}
	}
}

// checkSign 接口签名验证
func (g *ClientServiceGroup) checkSign() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (reply any, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				err = ErrSignInvalid.NewWithMessage("缺少签名参数")
				return
			}
			productId := tr.RequestHeader().Get(clientHeaderProductId)
			nonce := tr.RequestHeader().Get(clientHeaderNonce)
			timestamp := tr.RequestHeader().Get(clientHeaderTimestamp)
			sign := tr.RequestHeader().Get(clientHeaderSign)
			if len(productId) == 0 || len(nonce) == 0 || len(timestamp) == 0 || len(sign) == 0 {
				err = ErrSignInvalid.NewWithMessage("缺少签名参数")
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
				err = ErrSignInvalid.NewWithMessage("签名错误")
				return
			}
			return handler(ctx, req)
		}
	}
}

// checkAuth 验证访问令牌
func (g *ClientServiceGroup) checkAuth() middleware.Middleware {
	whiteList := make(map[string]struct{})
	whiteList[v1.OperationAuthAPILogin] = struct{}{}
	return selector.Server(func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (resp any, err error) {
			tr, ok := transport.FromServerContext(ctx)
			if !ok {
				err = ErrAccessTokenInvalid.NewWithMessage("缺少访问令牌参数")
				return
			}
			productId := tr.RequestHeader().Get(clientHeaderProductId)
			token := tr.RequestHeader().Get(clientHeaderAccessToken)
			if len(productId) == 0 || len(token) == 0 {
				err = ErrAccessTokenInvalid.NewWithMessage("缺少访问令牌参数")
				return
			}
			uid, err := g.authUc.CheckToken(ctx, &options.TokenCheckOptions{ProductId: productId, Token: token})
			if err != nil {
				return
			}
			ctx = context.WithValue(ctx, biz.ContextValueUID, uid)
			return handler(ctx, req)
		}
	}).Match(func(ctx context.Context, operation string) bool {
		_, ok := whiteList[operation]
		return !ok
	}).Build()
}
