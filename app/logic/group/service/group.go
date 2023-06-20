package service

import (
	"context"
	v1 "rockimserver/apis/rockim/service/group/v1"
	"rockimserver/app/logic/group/biz"
)

type GroupService struct {
	uc *biz.GroupUseCase
	v1.UnimplementedGroupAPIServer
}

func NewGroupService(uc *biz.GroupUseCase) *GroupService {
	return &GroupService{uc: uc}
}

func (s *GroupService) FindGroupId(ctx context.Context, in *v1.GroupIdFindRequest) (out *v1.GroupIdFindResponse, err error) {
	groupId, err := s.uc.FindGroupId(ctx, in.Base.ProductId, in.BizId)
	if err != nil {
		return
	}
	out = &v1.GroupIdFindResponse{GroupId: groupId}
	return
}

func (s *GroupService) Find(ctx context.Context, in *v1.GroupFindRequest) (out *v1.GroupFindResponse, err error) {
	groupId, err := s.uc.FindGroupId(ctx, in.Base.ProductId, in.BizId)
	if err != nil {
		return
	}
	group, err := s.uc.FindById(ctx, in.Base.ProductId, groupId)
	if err != nil {
		return
	}
	out = &v1.GroupFindResponse{Group: group}
	return
}

func (s *GroupService) FindById(ctx context.Context, in *v1.GroupFindByIdRequest) (out *v1.GroupFindByIdResponse, err error) {
	group, err := s.uc.FindById(ctx, in.Base.ProductId, in.GroupId)
	if err != nil {
		return
	}
	out = &v1.GroupFindByIdResponse{Group: group}
	return
}
