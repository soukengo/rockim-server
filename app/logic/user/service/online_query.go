package service

import (
	"context"
	"github.com/wxnacy/wgo/arrays"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/app/logic/user/biz"
	"rockimserver/app/logic/user/biz/options"
)

type OnlineQueryService struct {
	uc *biz.OnlineQueryUseCase
	v1.UnimplementedOnlineQueryAPIServer
}

func NewOnlineQueryService(uc *biz.OnlineQueryUseCase) *OnlineQueryService {
	return &OnlineQueryService{uc: uc}
}

func (s *OnlineQueryService) CheckOnline(ctx context.Context, in *v1.OnlineCheckRequest) (out *v1.OnlineCheckResponse, err error) {
	online, err := s.uc.CheckOnline(ctx, &options.OnlineCheckOptions{
		ProductId: in.Base.ProductId,
		Uid:       in.Uid,
	})
	if err != nil {
		return
	}
	out = &v1.OnlineCheckResponse{Online: online}
	return
}

func (s *OnlineQueryService) BatchCheckOnline(ctx context.Context, in *v1.OnlineBatchCheckRequest) (out *v1.OnlineBatchCheckResponse, err error) {
	onlineIds, err := s.uc.BatchCheckOnline(ctx, &options.OnlineBatchCheckOptions{
		ProductId: in.Base.ProductId,
		Uids:      in.Uids,
	})
	if err != nil {
		return
	}
	items := make([]*v1.OnlineBatchCheckResponse_Item, len(in.Uids))
	for i, uid := range in.Uids {
		online := arrays.ContainsString(onlineIds, uid) != -1
		items[i] = &v1.OnlineBatchCheckResponse_Item{
			Uid:    uid,
			Online: online,
		}
	}
	out = &v1.OnlineBatchCheckResponse{Items: items}
	return
}

func (s *OnlineQueryService) ListUser(ctx context.Context, in *v1.OnlineUserListRequest) (out *v1.OnlineUserListResponse, err error) {
	list, err := s.uc.ListUser(ctx, &options.OnlineUserListOptions{
		ProductId: in.Base.ProductId,
		Uids:      in.Uids,
	})
	if err != nil {
		return
	}
	out = &v1.OnlineUserListResponse{Users: list}
	return
}
