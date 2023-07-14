package service

import (
	"context"
	v1 "rockimserver/apis/rockim/service/session/v1"
	"rockimserver/app/logic/session/biz"
	"rockimserver/app/logic/session/biz/options"
)

type ChannelService struct {
	uc *biz.ChannelUseCase
	v1.UnimplementedChannelAPIServer
}

func NewChannelService(uc *biz.ChannelUseCase) *ChannelService {
	return &ChannelService{uc: uc}
}

func (s *ChannelService) Add(ctx context.Context, in *v1.ChannelAddRequest) (out *v1.ChannelAddResponse, err error) {
	err = s.uc.Add(ctx, &options.ChannelAddOptions{
		ProductId: in.Base.ProductId,
		ServerId:  in.ServerId,
		ChannelId: in.ChannelId,
		Uid:       in.Uid,
	})
	if err != nil {
		return
	}
	out = &v1.ChannelAddResponse{}
	return
}

func (s *ChannelService) Delete(ctx context.Context, in *v1.ChannelDeleteRequest) (out *v1.ChannelDeleteResponse, err error) {
	err = s.uc.Delete(ctx, &options.ChannelDeleteOptions{
		ProductId: in.Base.ProductId,
		ServerId:  in.ServerId,
		ChannelId: in.ChannelId,
	})
	if err != nil {
		return
	}
	out = &v1.ChannelDeleteResponse{}
	return
}

func (s *ChannelService) Refresh(ctx context.Context, in *v1.ChannelRefreshRequest) (out *v1.ChannelRefreshResponse, err error) {
	err = s.uc.Refresh(ctx, &options.ChannelRefreshOptions{
		ProductId: in.Base.ProductId,
		ServerId:  in.ServerId,
		ChannelId: in.ChannelId,
		Uid:       in.Uid,
	})
	if err != nil {
		return
	}
	out = &v1.ChannelRefreshResponse{}
	return
}
