package biz

import "context"

type TokenRepo interface {
	CreateToken(ctx context.Context, productId string, uid string, token string) error
	FindByToken(ctx context.Context, productId string, token string) (string, error)
	DeleteToken(ctx context.Context, productId string, token string) error
}

type LoginRequest struct {
}

type AuthUseCase struct {
	repo TokenRepo
}

func NewAuthUseCase(repo TokenRepo) *AuthUseCase {
	return &AuthUseCase{repo: repo}
}

func (uc *AuthUseCase) Login() {

}
