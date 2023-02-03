package converter

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"golang.org/x/exp/slices"
	apiTypes "rockimserver/apis/rockim/api/admin/tenant/v1/types"
	"rockimserver/app/access/admin/module/tenant/biz/types"
)

func ToSysTenantResource(source *types.SysTenantResource) *apiTypes.SysTenantResource {
	return &apiTypes.SysTenantResource{
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

func BuildTenantResourceTree(resourceList []*apiTypes.SysTenantResource) []*apiTypes.SysTenantResourceTree {
	if resourceList == nil || len(resourceList) == 0 {
		return []*apiTypes.SysTenantResourceTree{}
	}
	newList := arraylist.New()
	for _, v := range resourceList {
		newList.Add(v)
	}
	root := &apiTypes.SysTenantResourceTree{}
	root.Resource = &apiTypes.SysTenantResource{Id: "0"}
	root.Children = []*apiTypes.SysTenantResourceTree{}
	buildTenantSubResourceTree(root, newList)
	return root.Children
}

func buildTenantSubResourceTree(pTree *apiTypes.SysTenantResourceTree, resourceList *arraylist.List) {
	for i := 0; i < resourceList.Size(); i++ {
		item, _ := resourceList.Get(i)
		resource := item.(*apiTypes.SysTenantResource)
		tree := &apiTypes.SysTenantResourceTree{
			Resource: resource,
		}
		tree.Children = []*apiTypes.SysTenantResourceTree{}
		if pTree.Resource.Id == tree.Resource.ParentId {
			resourceList.Remove(i)
			buildTenantSubResourceTree(tree, resourceList)
			i--
			pTree.Children = append(pTree.Children, tree)
			slices.SortStableFunc[*apiTypes.SysTenantResourceTree](pTree.Children, func(o1, o2 *apiTypes.SysTenantResourceTree) bool {
				return o1.Resource.Order < o2.Resource.Order
			})
		}
	}
}
