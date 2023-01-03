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

type PlatResourceService struct {
	uc *manager.PlatResourceUseCase
}

func NewPlatResourceService(uc *manager.PlatResourceUseCase) *PlatResourceService {
	return &PlatResourceService{uc: uc}
}

func (s *PlatResourceService) Create(ctx context.Context, in *adminV1.PlatResourceCreateRequest) (reply *adminV1.PlatResourceCreateResponse, err error) {
	req := &v1.PlatResourceCreateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Create(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.PlatResourceCreateResponse{}
	return
}

func (s *PlatResourceService) Update(ctx context.Context, in *adminV1.PlatResourceUpdateRequest) (reply *adminV1.PlatResourceUpdateResponse, err error) {
	req := &v1.PlatResourceUpdateRequest{}
	err = copier.Copy(req, in)
	if err != nil {
		return
	}
	err = s.uc.Update(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.PlatResourceUpdateResponse{}
	return
}

func (s *PlatResourceService) Delete(ctx context.Context, in *adminV1.PlatResourceDeleteRequest) (reply *adminV1.PlatResourceDeleteResponse, err error) {
	req := &v1.PlatResourceDeleteRequest{Id: in.Id}
	err = s.uc.Delete(ctx, req)
	if err != nil {
		return
	}
	reply = &adminV1.PlatResourceDeleteResponse{}
	return
}

func (s *PlatResourceService) ListTree(ctx context.Context, in *adminV1.PlatResourceTreeRequest) (reply *adminV1.PlatResourceTreeResponse, err error) {
	list, err := s.uc.List(ctx)
	if err != nil {
		return
	}
	resources := make([]*adminTypes.PlatResource, len(list))
	for i, item := range list {
		resources[i] = convertResource(item)
	}
	return &adminV1.PlatResourceTreeResponse{List: buildResourceTree(resources)}, nil
}

func convertResource(source *types.PlatResource) *adminTypes.PlatResource {
	return &adminTypes.PlatResource{
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

func buildResourceTree(resourceList []*adminTypes.PlatResource) []*adminTypes.PlatResourceTree {
	if resourceList == nil || len(resourceList) == 0 {
		return []*adminTypes.PlatResourceTree{}
	}
	newList := arraylist.New()
	for _, v := range resourceList {
		newList.Add(v)
	}
	root := &adminTypes.PlatResourceTree{}
	root.Resource = &adminTypes.PlatResource{Id: "0"}
	root.Children = []*adminTypes.PlatResourceTree{}
	buildSubResourceTree(root, newList)
	return root.Children
}

func buildSubResourceTree(pTree *adminTypes.PlatResourceTree, resourceList *arraylist.List) {
	for i := 0; i < resourceList.Size(); i++ {
		item, _ := resourceList.Get(i)
		resource := item.(*adminTypes.PlatResource)
		tree := &adminTypes.PlatResourceTree{
			Resource: resource,
		}
		tree.Children = []*adminTypes.PlatResourceTree{}
		if pTree.Resource.Id == tree.Resource.ParentId {
			resourceList.Remove(i)
			buildSubResourceTree(tree, resourceList)
			i--
			pTree.Children = append(pTree.Children, tree)
			slices.SortStableFunc[*adminTypes.PlatResourceTree](pTree.Children, func(o1, o2 *adminTypes.PlatResourceTree) bool {
				return o1.Resource.Order < o2.Resource.Order
			})
		}
	}
}
