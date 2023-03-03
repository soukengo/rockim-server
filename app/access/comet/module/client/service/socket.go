package service

import (
	"context"
	v1 "rockimserver/apis/rockim/api/client/v1/socket"
)

type SocketService struct {
	v1.UnimplementedSocketAPIServer
}

func (s *SocketService) Connect(ctx context.Context, channelId string) (err error) {
	return
}
func (s *SocketService) DisConnect(ctx context.Context, channelId string) (err error) {
	return
}

func (s *SocketService) Auth(ctx context.Context, request *v1.AuthRequest) (*v1.AuthResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SocketService) HeartBeat(ctx context.Context, request *v1.HeartBeatRequest) (*v1.HeartBeatResponse, error) {
	//TODO implement me
	panic("implement me")
}
