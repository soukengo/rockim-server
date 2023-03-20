package socket

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
	V1 = uint16(1)
)

type Packet struct {
	Version   uint16
	Typ       uint8
	HeaderLen uint16
	BodyLen   uint32
	Header    []byte
	Body      []byte
}

func NewPacket(typ uint8, header []byte, body []byte) *Packet {
	return &Packet{
		Version:   V1,
		Typ:       typ,
		HeaderLen: uint16(len(header)),
		Header:    header,
		BodyLen:   uint32(len(body)),
		Body:      body,
	}
}

func (p *Packet) String() string {
	return fmt.Sprintf("Version: %v,Typ: %v,Header: %v,Body: %v", p.Version, p.Typ, string(p.Header), string(p.Body))
}

type packetOptions struct {
	versionSize   uint16
	typSize       uint8
	headerLenSize uint16
	bodyLenSize   uint32

	versionOffset   int
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
		versionSize:   2,
		typSize:       1,
		headerLenSize: 2,
		bodyLenSize:   4,
		versionOffset: 0,
	}
	opt.typOffset = opt.versionOffset + int(opt.versionSize)
	opt.headerLenOffset = opt.typOffset + int(opt.typSize)
	opt.bodyLenOffset = opt.headerLenOffset + int(opt.headerLenSize)
	opt.headerOffset = opt.bodyLenOffset + int(opt.bodyLenSize)
	opt.packetHeaderSize = int(opt.versionSize) + int(opt.typSize) + int(opt.headerLenSize) + int(opt.bodyLenSize)
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
	version := binary.BigEndian.Uint16(buf[opt.versionOffset:opt.typOffset])
	typ := buf[opt.typOffset]
	headerLen := binary.BigEndian.Uint16(buf[opt.headerLenOffset:opt.bodyLenOffset])
	bodyLen := binary.BigEndian.Uint32(buf[opt.bodyLenOffset:opt.headerOffset])
	var header []byte
	// 整体包长度
	packetLen = packetLen + int(headerLen) + int(bodyLen)
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
	// discard
	_, _ = bio.Discard(packetLen)
	p.Version = version
	p.Typ = typ
	p.HeaderLen = headerLen
	p.BodyLen = bodyLen
	p.Header = header
	p.Body = body
	return
}

func (p *Packet) PackTo(w io.Writer) (err error) {
	opt := opts
	bio := bufio.NewWriter(w)
	buf := make([]byte, opt.packetHeaderSize)
	binary.BigEndian.PutUint16(buf[opt.versionOffset:opt.typOffset], p.Version)
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
	err = bio.Flush()
	return
}
