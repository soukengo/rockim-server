package manager

import (
	"context"
	jwtAuth "github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/golang-jwt/jwt/v4"
	"rockim/api/rockim/service/platform/v1/types"
)

type SessionUseCase struct {
	repo         PlatUserRepo
	roleRepo     PlatRoleRepo
	resourceRepo PlatResourceRepo
}

func NewSessionUseCase(repo PlatUserRepo, roleRepo PlatRoleRepo, resourceRepo PlatResourceRepo) *SessionUseCase {
	return &SessionUseCase{repo: repo, roleRepo: roleRepo, resourceRepo: resourceRepo}
}

type Claims struct {
	jwt.RegisteredClaims
	User *SessionUser `json:"user"`
}

type SessionUser struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatarUrl"`
	IsAdmin   bool   `json:"isAdmin"`
}

func NewSessionUser(id string, name string, avatarUrl string, isAdmin bool) *SessionUser {
	return &SessionUser{ID: id, Name: name, AvatarUrl: avatarUrl, IsAdmin: isAdmin}
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

func (uc *SessionUseCase) ListResources(ctx context.Context) (res []*types.PlatResource, err error) {
	u, err := uc.Check(ctx)
	if err != nil {
		return
	}
	if u.IsAdmin {
		return uc.resourceRepo.List(ctx)
	}
	roleIds, err := uc.repo.ListRoleId(ctx, u.ID)
	if err != nil {
		return
	}
	if len(roleIds) == 0 {
		res = []*types.PlatResource{}
		return
	}
	resourceIds, err := uc.roleRepo.ListResourceId(ctx, roleIds)
	if err != nil {
		return
	}
	if len(resourceIds) == 0 {
		res = []*types.PlatResource{}
		return
	}
	return uc.resourceRepo.ListByIds(ctx, resourceIds)
}
