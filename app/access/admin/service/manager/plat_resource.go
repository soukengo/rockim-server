package manager

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"rockim/api/rockim/admin/manager/v1/types"
)

func buildResourceTree(resourceList []*types.PlatResource) []*types.PlatResourceTree {
	if resourceList == nil || len(resourceList) == 0 {
		return []*types.PlatResourceTree{}
	}
	newList := arraylist.New()
	for _, v := range resourceList {
		newList.Add(v)
	}
	root := &types.PlatResourceTree{}
	root.Resource = &types.PlatResource{Id: "0"}
	root.Children = []*types.PlatResourceTree{}
	buildSubResourceTree(root, newList)
	return root.Children
}

func buildSubResourceTree(pTree *types.PlatResourceTree, resourceList *arraylist.List) {
	for i := 0; i < resourceList.Size(); i++ {
		item, _ := resourceList.Get(i)
		resource := item.(*types.PlatResource)
		tree := &types.PlatResourceTree{
			Resource: resource,
		}
		tree.Children = []*types.PlatResourceTree{}
		if pTree.Resource.Id == tree.Resource.ParentId {
			resourceList.Remove(i)
			buildSubResourceTree(tree, resourceList)
			i--
			pTree.Children = append(pTree.Children, tree)
		}
	}
}
