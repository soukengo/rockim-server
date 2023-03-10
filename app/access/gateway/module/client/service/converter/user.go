package converter

import (
	clienttypes "rockimserver/apis/rockim/api/client/http/v1/types"
	"rockimserver/apis/rockim/service/user/v1/types"
)

func ToClientUser(user *types.User) *clienttypes.User {
	return &clienttypes.User{
		ProductId: user.ProductId,
		Account:   user.Account,
		Name:      user.Name,
		AvatarUrl: user.AvatarUrl,
		Fields:    user.Fields,
		Status:    user.Status,
	}
}
