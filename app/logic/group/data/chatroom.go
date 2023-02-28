package data

import (
	"context"
	"rockimserver/apis/rockim/service/group/v1/types"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/data/cache"
	"rockimserver/app/logic/group/data/database"
)

type chatRoomRepo struct {
	db *database.ChatRoomData
	ch *cache.ChatRoomData
}

func NewChatRoomRepo(db *database.ChatRoomData, ch *cache.ChatRoomData) biz.ChatRoomRepo {
	return &chatRoomRepo{db: db, ch: ch}
}

func (r *chatRoomRepo) FindByCustomGroupId(ctx context.Context, productId string, customGroupId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *chatRoomRepo) FindById(ctx context.Context, productId string, groupId string) (*types.Group, error) {
	//TODO implement me
	panic("implement me")
}

func (r *chatRoomRepo) Create(ctx context.Context, group *types.Group) error {
	//TODO implement me
	panic("implement me")
}

func (r *chatRoomRepo) Delete(ctx context.Context, productId string, groupId string) error {
	//TODO implement me
	panic("implement me")
}
