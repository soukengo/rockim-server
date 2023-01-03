package manager

import (
	"context"
	"github.com/emirpasic/gods/lists/arraylist"
	"golang.org/x/exp/slices"
	adminV1 "rockim/api/rockim/admin/manager/v1"
	adminTypes "rockim/api/rockim/admin/manager/v1/types"
	"rockim/api/rockim/service/platform/v1"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/biz/manager"
	"rockim/pkg/util/copier"
)

type TenantResourceService struct {
	uc *manager.TenantResourceUseCase
}

func NewTenantResourceService(uc *manager.TenantResourceUseCase) *TenantResourceService {
	return &TenantResourceService{uc: uc}
}

func (s *TenantResourceService) Create(ctx context.Context, in *adminV1.TenantResourceCreateRequest) (reply *adminV1.TenantResourceCreateResponse, err error) {
	req := &v1.TenantResourceCreateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.TenantResourceCreateResponse{}
	return
}

func (s *TenantResourceService) Update(ctx context.Context, in *adminV1.TenantResourceUpdateRequest) (reply *adminV1.TenantResourceUpdateResponse, err error) {
	req := &v1.TenantResourceUpdateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.TenantResourceUpdateResponse{}
	return
}

func (s *TenantResourceService) Delete(ctx context.Context, in *adminV1.TenantResourceDeleteRequest) (reply *adminV1.TenantResourceDeleteResponse, err error) {
	req := &v1.TenantResourceDeleteRequest{Id: in.Id}
	err = s.uc.Delete(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.TenantResourceDeleteResponse{}
	return
}

func (s *TenantResourceService) ListTree(ctx context.Context, in *adminV1.TenantResourceTreeRequest) (reply *adminV1.TenantResourceTreeResponse, err error) {
	list, err := s.uc.List(ctx)
	if err != nil {
		return
	}
	resources := make([]*adminTypes.TenantResource, len(list))
	for i, item := range list {
		resources[i] = convertTenantResource(item)
	}
	return &adminV1.TenantResourceTreeResponse{List: buildTenantResourceTree(resources)}, nil
}

func convertTenantResource(source *types.TenantResource) *adminTypes.TenantResource {
	return &adminTypes.TenantResource{
		Id:         source.Id,
		CreateTime: source.CreateTime,
		Category:   source.Category,
		Name:       source.Name,
		ParentId:   source.ParentId,
		Url:        source.Url,
		Icon:       source.Icon,
		Order:      source.Order,
	}
}

func buildTenantResourceTree(resourceList []*adminTypes.TenantResource) []*adminTypes.TenantResourceTree {
	if resourceList == nil || len(resourceList) == 0 {
		return []*adminTypes.TenantResourceTree{}
	}
	newList := arraylist.New()
	for _, v := range resourceList {
		newList.Add(v)
	}
	root := &adminTypes.TenantResourceTree{}
	root.Resource = &adminTypes.TenantResource{Id: "0"}
	root.Children = []*adminTypes.TenantResourceTree{}
	buildTenantSubResourceTree(root, newList)
	return root.Children
}

func buildTenantSubResourceTree(pTree *adminTypes.TenantResourceTree, resourceList *arraylist.List) {
	for i := 0; i < resourceList.Size(); i++ {
		item, _ := resourceList.Get(i)
		resource := item.(*adminTypes.TenantResource)
		tree := &adminTypes.TenantResourceTree{
			Resource: resource,
		}
		tree.Children = []*adminTypes.TenantResourceTree{}
		if pTree.Resource.Id == tree.Resource.ParentId {
			resourceList.Remove(i)
			buildTenantSubResourceTree(tree, resourceList)
			i--
			pTree.Children = append(pTree.Children, tree)
			slices.SortStableFunc[*adminTypes.TenantResourceTree](pTree.Children, func(o1, o2 *adminTypes.TenantResourceTree) bool {
				return o1.Resource.Order < o2.Resource.Order
			})
		}
	}
}
