package manager

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/conf"
	"rockim/pkg/util/encrypt"
	"time"
)

type PlatUserRepo interface {
	FindByAccount(ctx context.Context, account string) (*types.PlatUser, error)
}

type AuthUseCase struct {
	authCfg *conf.Auth
	repo    PlatUserRepo
}

func NewAuthUseCase(authCfg *conf.Auth, repo PlatUserRepo) *AuthUseCase {
	return &AuthUseCase{authCfg: authCfg, repo: repo}
}

func (uc *AuthUseCase) Login(ctx context.Context, account string, password string) (token string, err error) {
	user, err := uc.repo.FindByAccount(ctx, account)
	if err != nil {
		return
	}
	if encrypt.MD5(password) != user.Password {
		err = ErrPasswordIncorrect
		return
	}
	var session = &SessionUser{ID: user.Id, Name: user.Name, AvatarUrl: user.AvatarUrl}
	claims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		Claims{RegisteredClaims: jwt.RegisteredClaims{
			ID:        user.Id,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(uc.authCfg.Jwt.Expires)),
		}, User: session},
	)
	return claims.SignedString([]byte(uc.authCfg.Jwt.AppKey))
}
