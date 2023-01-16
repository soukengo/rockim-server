package service

import (
	"context"
	adminV1 "rockim/api/rockim/admin/manager/v1"
	apiTypes "rockim/api/rockim/admin/manager/v1/types"
	"rockim/app/access/admin/module/manager/biz"
	"rockim/app/access/admin/module/manager/biz/options"
	"rockim/app/access/admin/module/manager/service/converter"
	"rockim/pkg/util/copier"
)

type SysResourceService struct {
	uc *biz.SysResourceUseCase
}

func NewSysResourceService(uc *biz.SysResourceUseCase) *SysResourceService {
	return &SysResourceService{uc: uc}
}

func (s *SysResourceService) Create(ctx context.Context, in *adminV1.SysResourceCreateRequest) (out *adminV1.SysResourceCreateResponse, err error) {
	req := &options.SysResourceCreateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	out = &adminV1.SysResourceCreateResponse{}
	return
}

func (s *SysResourceService) Update(ctx context.Context, in *adminV1.SysResourceUpdateRequest) (out *adminV1.SysResourceUpdateResponse, err error) {
	req := &options.SysResourceUpdateOptions{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	out = &adminV1.SysResourceUpdateResponse{}
	return
}

func (s *SysResourceService) Delete(ctx context.Context, in *adminV1.SysResourceDeleteRequest) (out *adminV1.SysResourceDeleteResponse, err error) {
	err = s.uc.Delete(ctx, in.Id)
	if err != nil {
		return
	}
	out = &adminV1.SysResourceDeleteResponse{}
	return
}

func (s *SysResourceService) ListTree(ctx context.Context, in *adminV1.SysResourceTreeRequest) (out *adminV1.SysResourceTreeResponse, err error) {
	list, err := s.uc.List(ctx)
	if err != nil {
		return
	}
	resources := make([]*apiTypes.SysResource, len(list))
	for i, item := range list {
		resources[i] = converter.ToManagerSysResource(item)
	}
	return &adminV1.SysResourceTreeResponse{List: converter.BuildResourceTree(resources)}, nil
}
