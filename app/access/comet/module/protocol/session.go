package protocol

import (
	"errors"
	"github.com/antlabs/timer"
	"github.com/soukengo/gopkg/component/transport/socket"
	"sync/atomic"
)

var (
	ErrInvalidSession = errors.New("invalid session")
)

type Session struct {
	ProductId               string
	Uid                     string
	LastHeartBeatTime       int64
	ServerHeartBeatInterval int64
	LastServerHeartBeatTime int64
	AuthCheckTimer          timer.TimeNoder
	HeartBeatCheckTimer     timer.TimeNoder

	authenticated int32 // 是否已认证
}

func (s *Session) MarkAuthenticated() {
	atomic.StoreInt32(&s.authenticated, 1)
}

func (s *Session) Authenticated() bool {
	return atomic.LoadInt32(&s.authenticated) > 0
}

func SessionFromChannel(ch socket.Channel) (out *Session, err error) {
	session := ch.Session()
	if session == nil {
		err = ErrInvalidSession
		return
	}
	out, ok := session.(*Session)
	if !ok {
		err = ErrInvalidSession
		return
	}
	return

}
