package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	sharedTypes "rockim/api/rockim/shared/types"
	"rockim/pkg/log"
	"rockim/pkg/util/copier"
	"rockim/pkg/util/encrypt"
	"time"
)

const (
	defaultPassword = "123456"
)

var (
	ErrPlatUserNotFound         = errors.NotFound(v1.ErrorReason_PLAT_USER_NOT_FOUND.String(), "账号不存在")
	ErrPlatUserAccountDuplicate = errors.Conflict(v1.ErrorReason_PLAT_USER_ACCOUNT_DUPLICATE.String(), "不能重复创建")
)

// PlatUserRepo is a PlatUser repo.
type PlatUserRepo interface {
	Create(ctx context.Context, record *types.PlatUser) error
	Update(ctx context.Context, record *types.PlatUser) error
	Delete(ctx context.Context, id string) error

	FindById(ctx context.Context, id string) (user *types.PlatUser, err error)
	FindIdByAccount(ctx context.Context, account string) (uid string, err error)
	Paging(ctx context.Context, req *PlatUserPagingRequest) (res *PlatUserPagingResponse, err error)

	ListRoleId(ctx context.Context, userId string) ([]string, error)
	SaveRoleId(ctx context.Context, userId string, roleIds []string) (err error)
}

// PlatUserUseCase is a PlatUser use case.
type PlatUserUseCase struct {
	repo PlatUserRepo
}

type PlatUserOptions struct {
	Name string
}

type PlatUserCreateRequest struct {
	Options  *PlatUserOptions
	Account  string
	Password string
}
type PlatUserUpdateRequest struct {
	Id      string
	Options *PlatUserOptions
}
type PlatUserDeleteRequest struct {
	Id string
}

type PlatUserPagingRequest struct {
	Paginate *sharedTypes.Paginating
	Keyword  string
}
type PlatUserPagingResponse struct {
	List     []*types.PlatUser
	Paginate *sharedTypes.Paginated
}

// NewPlatUserUseCase new a PlatUser use case.
func NewPlatUserUseCase(repo PlatUserRepo) *PlatUserUseCase {
	return &PlatUserUseCase{repo: repo}
}

func (uc *PlatUserUseCase) Create(ctx context.Context, req *PlatUserCreateRequest) (err error) {
	uid, err := uc.repo.FindIdByAccount(ctx, req.Account)
	if err != nil && !errors.IsNotFound(err) {
		return
	}
	if len(uid) > 0 {
		return ErrPlatUserAccountDuplicate
	}
	record := &types.PlatUser{Name: req.Options.Name, Account: req.Account, Password: req.Password}
	plainPwd := req.Password
	if len(plainPwd) == 0 {
		plainPwd = defaultPassword
	}
	record.Password = encrypt.MD5(plainPwd)
	record.Id = primitive.NewObjectID().Hex()
	record.CreateTime = time.Now().UnixMilli()
	return uc.repo.Create(ctx, record)
}
func (uc *PlatUserUseCase) Update(ctx context.Context, req *PlatUserUpdateRequest) (err error) {
	record, err := uc.repo.FindById(ctx, req.Id)
	if err != nil {
		return
	}
	err = copier.CopyWithOption(record, req.Options, copier.Option{IgnoreEmpty: true})
	if err != nil {
		return
	}
	record.UpdateTime = time.Now().UnixMilli()
	return uc.repo.Update(ctx, record)
}
func (uc *PlatUserUseCase) Delete(ctx context.Context, req *PlatUserDeleteRequest) (err error) {
	return uc.repo.Delete(ctx, req.Id)
}

func (uc *PlatUserUseCase) Find(ctx context.Context, account string) (user *types.PlatUser, err error) {
	uid, err := uc.repo.FindIdByAccount(ctx, account)
	if err != nil {
		log.WithContext(ctx).Errorf("FindIdByAccount account: %v error: %v", account, err)
		err = ErrPlatUserNotFound
		return
	}
	return uc.repo.FindById(ctx, uid)
}

func (uc *PlatUserUseCase) Paging(ctx context.Context, req *PlatUserPagingRequest) (res *PlatUserPagingResponse, err error) {
	return uc.repo.Paging(ctx, req)
}

func (uc *PlatUserUseCase) ListRoleId(ctx context.Context, userId string) (res []string, err error) {
	return uc.repo.ListRoleId(ctx, userId)
}

func (uc *PlatUserUseCase) SaveRoleId(ctx context.Context, userId string, roleIds []string) (err error) {
	return uc.repo.SaveRoleId(ctx, userId, roleIds)
}
