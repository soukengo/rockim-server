package tenant

import (
	"context"
	"github.com/emirpasic/gods/lists/arraylist"
	"golang.org/x/exp/slices"
	v1 "rockim/api/rockim/admin/tenant/v1"
	adminTypes "rockim/api/rockim/admin/tenant/v1/types"
	"rockim/api/rockim/service/platform/v1/types"
	"rockim/app/access/admin/biz/tenant"
)

type SessionService struct {
	uc *tenant.SessionUseCase
}

func NewSessionService(uc *tenant.SessionUseCase) *SessionService {
	return &SessionService{uc: uc}
}

func (s *SessionService) Check(ctx context.Context, in *v1.SessionCheckRequest) (reply *v1.SessionCheckResponse, err error) {
	session, err := s.uc.Check(ctx)
	if err != nil {
		return
	}
	reply = &v1.SessionCheckResponse{
		Session: &adminTypes.SessionUser{Id: session.ID, Name: session.Name, AvatarUrl: session.AvatarUrl},
	}
	return
}

func (s *SessionService) ListResource(ctx context.Context, in *v1.SessionListResourceRequest) (reply *v1.SessionListResourceResponse, err error) {
	list, err := s.uc.ListResources(ctx)
	if err != nil {
		return
	}
	resources := make([]*adminTypes.TenantResource, len(list))
	for i, item := range list {
		resources[i] = convertTenantResource(item)
	}
	reply = &v1.SessionListResourceResponse{
		List: buildTenantResourceTree(resources),
	}
	return
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
