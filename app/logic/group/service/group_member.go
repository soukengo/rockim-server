package service

import (
	"context"
	v1 "rockimserver/apis/rockim/service/group/v1"
	"rockimserver/app/logic/group/biz"
	"rockimserver/app/logic/group/biz/options"
)

type GroupMemberService struct {
	uc *biz.GroupMemberUseCase
	v1.UnimplementedGroupMemberAPIServer
}

func (s *GroupMemberService) IsMember(ctx context.Context, in *v1.GroupMemberCheckRequest) (out *v1.GroupMemberCheckResponse, err error) {
	isMember, err := s.uc.IsMember(ctx, &options.GroupMemberCheckOptions{
		ProductId: in.Base.ProductId,
		GroupId:   in.GroupId,
		Uid:       in.Uid,
	})
	if err != nil {
		return
	}
	out = &v1.GroupMemberCheckResponse{IsMember: isMember}
	return
}

func (s *GroupMemberService) Find(ctx context.Context, in *v1.GroupMemberFindRequest) (out *v1.GroupMemberFindResponse, err error) {
	member, err := s.uc.Find(ctx, &options.GroupMemberFindOptions{
		ProductId: in.Base.ProductId,
		GroupId:   in.GroupId,
		Uid:       in.Uid,
	})
	if err != nil {
		return
	}
	out = &v1.GroupMemberFindResponse{Member: member}
	return
}

func (s *GroupMemberService) List(ctx context.Context, in *v1.GroupMemberListRequest) (out *v1.GroupMemberListResponse, err error) {
	list, err := s.uc.List(ctx, &options.GroupMemberListOptions{
		ProductId: in.Base.ProductId,
		GroupId:   in.GroupId,
		Uids:      in.Uids,
	})
	if err != nil {
		return
	}
	out = &v1.GroupMemberListResponse{Members: list}
	return
}
func (s *GroupMemberService) PaginateUid(ctx context.Context, in *v1.GroupMemberIdPaginateRequest) (out *v1.GroupMemberIdPaginateResponse, err error) {
	list, paginated, err := s.uc.PaginateUid(ctx, &options.GroupMemberIdPaginateOptions{
		ProductId: in.Base.ProductId,
		GroupId:   in.GroupId,
	})
	if err != nil {
		return
	}
	out = &v1.GroupMemberIdPaginateResponse{Members: list, Paginate: paginated}
	return
}

func (s *GroupMemberService) ListGroupIdByUid(ctx context.Context, in *v1.GroupIdListByUidRequest) (out *v1.GroupIdListByUidResponse, err error) {
	groupIds, err := s.uc.ListGroupIdByUid(ctx, &options.GroupIdListByUidOptions{ProductId: in.Base.ProductId, Uid: in.Uid})
	if err != nil {
		return
	}
	out = &v1.GroupIdListByUidResponse{GroupIds: groupIds}
	return
}
