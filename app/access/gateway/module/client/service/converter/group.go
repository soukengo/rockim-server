package converter

import (
	clienttypes "rockimserver/apis/rockim/api/client/v1/group/types"
	"rockimserver/apis/rockim/service/group/v1/types"
)

func ToClientGroup(source *types.Group) *clienttypes.Group {
	return &clienttypes.Group{
		Category:      source.Category,
		ProductId:     source.ProductId,
		CustomGroupId: source.CustomGroupId,
		Name:          source.Name,
		IconUrl:       source.IconUrl,
		Fields:        source.Fields,
		Status:        source.Status,
	}
}
