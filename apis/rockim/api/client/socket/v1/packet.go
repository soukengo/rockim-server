package v1

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"rockimserver/pkg/component/server/socket/packet"
)

type PacketFactory struct {
	newFunc func() packet.IPacket
}

func (factory *PacketFactory) Offer(connId string) (p packet.IPacket) {
	return factory.newFunc()
}

func NewPacketFactory() packet.IFactory {
	return &PacketFactory{newFunc: func() packet.IPacket {
		return &Packet{}
	}}
}

const (
	firstFlag = 0x01
	lastFlag  = 0x10
)

type Packet struct {
	First     uint16
	Typ       uint8
	HeaderLen uint16
	BodyLen   uint32
	Header    []byte
	Body      []byte
	Last      uint16
}

func NewPacket(typ uint8, header []byte, body []byte) *Packet {
	return &Packet{
		First:     firstFlag,
		Typ:       typ,
		HeaderLen: uint16(len(header)),
		Header:    header,
		BodyLen:   uint32(len(body)),
		Body:      body,
		Last:      lastFlag,
	}
}

func (p *Packet) String() string {
	return fmt.Sprintf("First: %v,Typ: %v,Header: %v,Body: %v,Last: %v", p.First, p.Typ, string(p.Header), string(p.Body), p.Last)
}

type packetOptions struct {
	firstSize     uint16
	typSize       uint8
	headerLenSize uint16
	bodyLenSize   uint32
	lastSize      uint16

	firstOffset     int
	typOffset       int
	headerLenOffset int
	bodyLenOffset   int
	headerOffset    int

	packetHeaderSize int
}

var (
	opts = initOptions()
)

func initOptions() *packetOptions {
	opt := &packetOptions{
		firstSize:     2,
		typSize:       1,
		headerLenSize: 2,
		bodyLenSize:   4,
		lastSize:      uint16(2),
		firstOffset:   0,
	}
	opt.typOffset = opt.firstOffset + int(opt.firstSize)
	opt.headerLenOffset = opt.typOffset + int(opt.typSize)
	opt.bodyLenOffset = opt.headerLenOffset + int(opt.headerLenSize)
	opt.headerOffset = opt.bodyLenOffset + int(opt.bodyLenSize)
	opt.packetHeaderSize = int(opt.firstSize) + int(opt.typSize) + int(opt.headerLenSize) + int(opt.bodyLenSize)
	return opt
}

func (p *Packet) UnPackFrom(r io.Reader) (err error) {
	opt := opts
	bio := bufio.NewReader(r)
	packetLen := opt.packetHeaderSize
	buf, _ := bio.Peek(packetLen)
	if len(buf) < packetLen {
		return packet.ErrInvalidPacket
	}
	first := binary.BigEndian.Uint16(buf[opt.firstOffset:opt.typOffset])
	typ := buf[opt.typOffset]
	headerLen := binary.BigEndian.Uint16(buf[opt.headerLenOffset:opt.bodyLenOffset])
	bodyLen := binary.BigEndian.Uint32(buf[opt.bodyLenOffset:opt.headerOffset])
	var header []byte
	// 整体包长度
	packetLen = packetLen + int(headerLen) + int(bodyLen) + int(opt.lastSize)
	// 计算offset
	bodyOffset := opt.headerOffset + int(headerLen)
	lastOffset := uint32(bodyOffset) + bodyLen

	buf, _ = bio.Peek(packetLen)
	if len(buf) < packetLen {
		return packet.ErrInvalidPacket
	}
	if headerLen > 0 {
		header = buf[opt.headerOffset:bodyOffset]
	}
	var body []byte
	if bodyLen > 0 {
		body = buf[bodyOffset:lastOffset]
	}
	lastBuf := buf[lastOffset : lastOffset+uint32(opt.lastSize)]
	if err != nil {
		return
	}
	last := binary.BigEndian.Uint16(lastBuf)
	// discard
	_, _ = bio.Discard(packetLen)
	p.First = first
	p.Typ = typ
	p.HeaderLen = headerLen
	p.BodyLen = bodyLen
	p.Header = header
	p.Body = body
	p.Last = last
	return
}

func (p *Packet) PackTo(w io.Writer) (err error) {
	opt := opts
	bio := bufio.NewWriter(w)
	buf := make([]byte, opt.packetHeaderSize)
	binary.BigEndian.PutUint16(buf[opt.firstOffset:opt.typOffset], p.First)
	buf[opt.typOffset] = p.Typ
	binary.BigEndian.PutUint16(buf[opt.headerLenOffset:opt.bodyLenOffset], p.HeaderLen)
	binary.BigEndian.PutUint32(buf[opt.bodyLenOffset:opt.headerOffset], p.BodyLen)
	_, err = bio.Write(buf)
	if err != nil {
		return
	}
	if len(p.Header) > 0 {
		_, err = bio.Write(p.Header)
		if err != nil {
			return
		}
	}
	if len(p.Body) > 0 {
		_, err = bio.Write(p.Body)
		if err != nil {
			return
		}
	}
	lastBuf := make([]byte, opt.lastSize)
	binary.BigEndian.PutUint16(lastBuf, p.Last)
	_, err = bio.Write(lastBuf)
	if err != nil {
		return
	}
	err = bio.Flush()
	return
}
