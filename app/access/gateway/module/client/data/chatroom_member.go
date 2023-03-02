package data

import (
	"context"
	"rockimserver/apis/rockim/service"
	v1 "rockimserver/apis/rockim/service/group/v1"
	"rockimserver/app/access/gateway/module/client/biz"
	"rockimserver/app/access/gateway/module/client/biz/options"
)

type chatRoomMemberRepo struct {
	ac v1.ChatRoomMemberAPIClient
}

func NewChatRoomMemberRepo(ac v1.ChatRoomMemberAPIClient) biz.ChatRoomMemberRepo {
	return &chatRoomMemberRepo{ac: ac}
}

func (r *chatRoomMemberRepo) Join(ctx context.Context, opts *options.ChatRoomJoinOptions) (err error) {
	_, err = r.ac.Join(ctx, &v1.ChatRoomJoinRequest{
		Base:    service.GenServiceRequest(opts.ProductId),
		GroupId: opts.GroupId,
		Uid:     opts.Uid,
	})
	if err != nil {
		return
	}
	return
}

func (r *chatRoomMemberRepo) Quit(ctx context.Context, opts *options.ChatRoomQuitOptions) (err error) {
	_, err = r.ac.Quit(ctx, &v1.ChatRoomQuitRequest{
		Base:    service.GenServiceRequest(opts.ProductId),
		GroupId: opts.GroupId,
		Uid:     opts.Uid,
	})
	if err != nil {
		return
	}
	return
}
