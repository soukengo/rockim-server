package nbio

import (
	log "github.com/golang/glog"
	"github.com/google/uuid"
	"github.com/lesismal/nbio"
	"net/url"
	"rockimserver/pkg/component/server/socket/network"
	"rockimserver/pkg/component/server/socket/network/tcp"
	"rockimserver/pkg/component/server/socket/packet"
	"sync"
)

type nbioServer struct {
	id        string
	cfg       *tcp.Config
	endpoint  *url.URL
	parser    *packet.Parser
	handler   network.Handler
	eng       *nbio.Engine
	startOnce *sync.Once
}

func NewServer(cfg *tcp.Config, parser *packet.Parser) network.Server {
	g := nbio.NewEngine(nbio.Config{
		Network:            "tcp",
		Addrs:              []string{cfg.Addr},
		ReadBufferSize:     cfg.ReadBufSize,
		MaxWriteBufferSize: cfg.WriteBufSize,
	})
	s := &nbioServer{
		id:        uuid.New().String(),
		endpoint:  &url.URL{Scheme: "tcp", Path: cfg.Addr},
		parser:    parser,
		handler:   network.DefaultHandler,
		cfg:       cfg,
		startOnce: new(sync.Once),
		eng:       g,
	}
	g.OnOpen(s.onOpen)
	g.OnClose(s.onClose)
	g.OnData(s.onData)
	return s

}

func (s *nbioServer) Id() string {
	return s.id
}
func (s *nbioServer) Start() (err error) {
	s.startOnce.Do(func() {
		go func() {
			err = s.eng.Start()
			if err != nil {
				panic(err)
			}
		}()
		log.Infof("started tcp server with endpoint %s", s.endpoint)
	})
	return
}

func (s *nbioServer) Close() error {
	s.eng.Stop()
	return nil
}

func (s *nbioServer) SetHandler(handler network.Handler) {
	s.handler = handler
}

func (s *nbioServer) onOpen(c *nbio.Conn) {
	conn := newConn(c, s.parser)
	c.SetSession(conn)
	s.handler.OnConnect(conn)
}

func (s *nbioServer) onClose(c *nbio.Conn, err error) {
	conn := s.parseConn(c)
	if conn == nil {
		return
	}
	conn.markClosed()
	s.handler.OnDisConnect(conn)
}

func (s *nbioServer) onData(c *nbio.Conn, data []byte) {
	conn := s.parseConn(c)
	if conn == nil {
		c.Close()
		return
	}
	conn.buf.Write(data)
	for {
		p, err := conn.Read()
		if err != nil {
			if err != packet.ErrInvalidPacket {
				log.Errorf("reader error: %v", err)
			}
			return
		}
		s.handler.OnReceived(conn, p)
	}

}

func (s *nbioServer) parseConn(c *nbio.Conn) *nbioConn {
	conn, ok := c.Session().(*nbioConn)
	if !ok {
		return nil
	}
	return conn
}
