package biz

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/soukengo/gopkg/component/auth"
	"github.com/soukengo/gopkg/util/encrypt"
	"time"
)

type AuthUseCase struct {
	authCfg *auth.Config
	repo    TenantRepo
}

func NewAuthUseCase(authCfg *auth.Config, repo TenantRepo) *AuthUseCase {
	return &AuthUseCase{authCfg: authCfg, repo: repo}
}

func (uc *AuthUseCase) Login(ctx context.Context, email string, password string) (token string, err error) {
	tenant, err := uc.repo.FindByEmail(ctx, email)
	if err != nil {
		return
	}
	if encrypt.MD5(password) != tenant.Password {
		err = ErrPasswordIncorrect
		return
	}
	var session = NewSessionUser(tenant.Id, tenant.Id, tenant.Name, "", true)
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		Claims{RegisteredClaims: jwt.RegisteredClaims{
			ID:        tenant.Id,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(uc.authCfg.Jwt.Expires)),
		}, User: session},
	)
	return claims.SignedString([]byte(uc.authCfg.Jwt.AppKey))
}
