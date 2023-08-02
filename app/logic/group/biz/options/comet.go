package options

import (
	"github.com/golang/protobuf/proto"
	comettypes "rockimserver/apis/rockim/service/comet/v1/types"
	"rockimserver/apis/rockim/shared/enums"
)

type PushOptions struct {
	ProductId string
	Operation enums.Comet_PushOperation
	Uids      []string
	Data      proto.Message
}
type PushRoomOptions struct {
	ProductId string
	Operation enums.Comet_PushOperation
	Room      *comettypes.Room
	Data      proto.Message
}

type ControlOptions struct {
	ProductId   string
	ControlType enums.Comet_ControlType
	Uids        []string
	Data        proto.Message
}
