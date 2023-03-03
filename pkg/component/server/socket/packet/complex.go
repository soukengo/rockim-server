package packet

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

type ComplexPacket struct {
	first     uint16
	op        uint16
	headerLen uint16
	bodyLen   uint32
	header    []byte
	body      []byte
	last      uint16
}

func NewComplexPacket(first uint16, op uint16, header []byte, body []byte, last uint16) *ComplexPacket {
	return &ComplexPacket{
		first:     first,
		op:        op,
		headerLen: uint16(len(header)),
		header:    header,
		bodyLen:   uint32(len(body)),
		body:      body,
		last:      last,
	}
}

func (p *ComplexPacket) String() string {
	return fmt.Sprintf("first: %v,op: %v,header: %v,body: %v,last: %v", p.first, p.op, string(p.header), string(p.body), p.last)
}

type complexPacketOptions struct {
	firstSize     uint16
	opSize        uint16
	headerLenSize uint16
	bodyLenSize   uint32
	lastSize      uint16

	firstOffset     int
	opOffset        int
	headerLenOffset int
	bodyLenOffset   int
	headerOffset    int

	packetHeaderSize int
}

var (
	complexOpt = initComplexPacketOptions()
)

func initComplexPacketOptions() *complexPacketOptions {
	opt := &complexPacketOptions{
		firstSize:     2,
		opSize:        2,
		headerLenSize: 2,
		bodyLenSize:   4,
		lastSize:      uint16(2),
		firstOffset:   0,
	}
	opt.opOffset = opt.firstOffset + int(opt.firstSize)
	opt.headerLenOffset = opt.opOffset + int(opt.opSize)
	opt.bodyLenOffset = opt.headerLenOffset + int(opt.headerLenSize)
	opt.headerOffset = opt.bodyLenOffset + int(opt.bodyLenSize)
	opt.packetHeaderSize = int(opt.firstSize) + int(opt.opSize) + int(opt.headerLenSize) + int(opt.bodyLenSize)
	return opt
}

func (p *ComplexPacket) UnPackFrom(r io.Reader) (err error) {
	opt := complexOpt
	bio := bufio.NewReader(r)
	packetLen := opt.packetHeaderSize
	buf, _ := bio.Peek(packetLen)
	if len(buf) < packetLen {
		return ErrInvalidPacket
	}
	first := binary.BigEndian.Uint16(buf[opt.firstOffset:opt.opOffset])
	op := binary.BigEndian.Uint16(buf[opt.opOffset:opt.headerLenOffset])
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
		return ErrInvalidPacket
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
	p.first = first
	p.op = op
	p.headerLen = headerLen
	p.bodyLen = bodyLen
	p.header = header
	p.body = body
	p.last = last
	return
}

func (p *ComplexPacket) PackTo(w io.Writer) (err error) {
	opt := complexOpt
	bio := bufio.NewWriter(w)
	buf := make([]byte, opt.packetHeaderSize)
	binary.BigEndian.PutUint16(buf[opt.firstOffset:opt.opOffset], p.first)
	binary.BigEndian.PutUint16(buf[opt.opOffset:opt.headerLenOffset], p.op)
	binary.BigEndian.PutUint16(buf[opt.headerLenOffset:opt.bodyLenOffset], p.headerLen)
	binary.BigEndian.PutUint32(buf[opt.bodyLenOffset:opt.headerOffset], p.bodyLen)
	_, err = bio.Write(buf)
	if err != nil {
		return
	}
	if len(p.header) > 0 {
		_, err = bio.Write(p.header)
		if err != nil {
			return
		}
	}
	if len(p.body) > 0 {
		_, err = bio.Write(p.body)
		if err != nil {
			return
		}
	}
	lastBuf := make([]byte, opt.lastSize)
	binary.BigEndian.PutUint16(lastBuf, p.last)
	_, err = bio.Write(lastBuf)
	if err != nil {
		return
	}
	err = bio.Flush()
	return
}
