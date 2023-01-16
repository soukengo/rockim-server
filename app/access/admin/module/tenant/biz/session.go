package biz

import (
	"context"
	jwtAuth "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
	"rockim/app/access/admin/module/tenant/biz/types"
)

type SessionUseCase struct {
	resourceRepo SysTenantResourceRepo
}

func NewSessionUseCase(resourceRepo SysTenantResourceRepo) *SessionUseCase {
	return &SessionUseCase{resourceRepo: resourceRepo}
}

type Claims struct {
	jwt.RegisteredClaims
	User *SessionUser `json:"user"`
}

type SessionUser struct {
	TenantId  string `json:"tenantId"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
	IsAdmin   bool   `json:"isAdmin"`
}

func NewSessionUser(tenantId string, id string, name string, avatarUrl string, isAdmin bool) *SessionUser {
	return &SessionUser{TenantId: tenantId, ID: id, Name: name, AvatarUrl: avatarUrl, IsAdmin: isAdmin}
}

func (uc *SessionUseCase) Check(ctx context.Context) (u *SessionUser, err error) {
	token, ok := jwtAuth.FromContext(ctx)
	if !ok {
		err = ErrAuthorizationInvalid
		return
	}
	claims, ok := token.(*Claims)
	if !ok {
		err = ErrAuthorizationInvalid
		return
	}
	return claims.User, nil
}

func (uc *SessionUseCase) ListResources(ctx context.Context) (res []*types.SysTenantResource, err error) {
	_, err = uc.Check(ctx)
	if err != nil {
		return
	}
	return uc.resourceRepo.List(ctx)
}
