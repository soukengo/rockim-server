package tenant

import (
	"context"
	jwtAuth "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
	"rockim/api/rockim/service/platform/v1/types"
)

type SessionUseCase struct {
	resourceRepo TenantResourceRepo
}

func NewSessionUseCase(resourceRepo TenantResourceRepo) *SessionUseCase {
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

func (uc *SessionUseCase) ListResources(ctx context.Context) (res []*types.TenantResource, err error) {
	_, err = uc.Check(ctx)
	if err != nil {
		return
	}
	return uc.resourceRepo.List(ctx)
}
