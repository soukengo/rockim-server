package service

import (
	"context"
	"github.com/wxnacy/wgo/arrays"
	v1 "rockimserver/apis/rockim/service/session/v1"
	"rockimserver/app/logic/session/biz"
	"rockimserver/app/logic/session/biz/options"
)

type ChannelQueryService struct {
	uc *biz.ChannelQueryUseCase
	v1.UnimplementedChannelQueryAPIServer
}

func NewChannelQueryService(uc *biz.ChannelQueryUseCase) *ChannelQueryService {
	return &ChannelQueryService{uc: uc}
}

func (s *ChannelQueryService) CheckOnline(ctx context.Context, in *v1.OnlineCheckRequest) (out *v1.OnlineCheckResponse, err error) {
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

func (s *ChannelQueryService) BatchCheckOnline(ctx context.Context, in *v1.OnlineBatchCheckRequest) (out *v1.OnlineBatchCheckResponse, err error) {
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

func (s *ChannelQueryService) ListUser(ctx context.Context, in *v1.UserChannelListRequest) (out *v1.UserChannelListResponse, err error) {
	list, err := s.uc.ListUser(ctx, &options.UserChannelListOptions{
		ProductId: in.Base.ProductId,
		Uids:      in.Uids,
	})
	if err != nil {
		return
	}
	out = &v1.UserChannelListResponse{Users: list}
	return
}
