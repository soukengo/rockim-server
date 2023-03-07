package service

import (
	"context"
	v1 "rockimserver/apis/rockim/service/user/v1"
	"rockimserver/app/logic/user/biz"
	"rockimserver/app/logic/user/biz/options"
)

type OnlineService struct {
	uc *biz.OnlineUseCase
	v1.UnimplementedOnlineAPIServer
}

func NewOnlineService(uc *biz.OnlineUseCase) *OnlineService {
	return &OnlineService{uc: uc}
}

func (s *OnlineService) Add(ctx context.Context, in *v1.OnlineAddRequest) (out *v1.OnlineAddResponse, err error) {
	uid, err := s.uc.Add(ctx, &options.OnlineAddOptions{
		ProductId: in.Base.ProductId,
		ServerId:  in.ServerId,
		ChannelId: in.ChannelId,
		Token:     in.Token,
	})
	if err != nil {
		return
	}
	out = &v1.OnlineAddResponse{Uid: uid}
	return
}

func (s *OnlineService) Delete(ctx context.Context, in *v1.OnlineDeleteRequest) (out *v1.OnlineDeleteResponse, err error) {
	err = s.uc.Delete(ctx, &options.OnlineDeleteOptions{
		ProductId: in.Base.ProductId,
		ServerId:  in.ServerId,
		ChannelId: in.ChannelId,
	})
	if err != nil {
		return
	}
	out = &v1.OnlineDeleteResponse{}
	return
}

func (s *OnlineService) Refresh(ctx context.Context, in *v1.OnlineRefreshRequest) (out *v1.OnlineRefreshResponse, err error) {
	err = s.uc.Refresh(ctx, &options.OnlineRefreshOptions{
		ProductId: in.Base.ProductId,
		ServerId:  in.ServerId,
		ChannelId: in.ChannelId,
	})
	if err != nil {
		return
	}
	out = &v1.OnlineRefreshResponse{}
	return
}
