package biz

import (
	"context"
	clienttypes "rockimserver/apis/rockim/api/client/v1/types"
)

type PushRepo interface {
	PushUsers(ctx context.Context, productId string, uids []string, packet *clienttypes.IMMessagePacket) error
	PushGroup(ctx context.Context, productId string, groupId string, packet *clienttypes.IMMessagePacket) error
}
