package gen

import (
	"errors"
	"net"
	"reflect"
	"unsafe"

	"golang.org/x/sys/unix"
)

type WlHeader struct {
	Id   uint32
	Op   uint16
	Size uint16
}

func parseHeader(bs []byte) (WlHeader, []byte, error) {
	if len(bs) < 8 {
		return WlHeader{}, nil, errors.New("ReadHeader: not enough data")
	}

	headerBytes := bs[:8]
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&headerBytes))

	return *(*WlHeader)(unsafe.Pointer(sliceHeader.Data)), bs[8:], nil
}

func marshHeader(head WlHeader) []byte {
	bs := make([]byte, 8)
	copy(bs, (*(*[8]byte)(unsafe.Pointer(&head)))[:])
	return bs
}

type WlMessage struct {
	WlHeader
	Data []byte
}

type WlWireMessage struct {
	Messages []WlMessage
	FDs      []int
}

func parseOneMessage(bs []byte) (msg WlMessage, rest []byte, err error) {
	msg.WlHeader, _, err = parseHeader(bs)
	if err != nil {
		return
	}
	if len(bs) < int(msg.Size) {
		err = errors.New("ReadMessage: actual size does not match header size")
		return
	}
	msg.Data = make([]byte, msg.Size-8)
	copy(msg.Data, bs[8:])
	rest = bs[msg.Size:]
	return
}

func marshOneMessage(msg WlMessage) []byte {
	bs := make([]byte, msg.Size)
	copy(bs, marshHeader(msg.WlHeader))
	copy(bs[8:], msg.Data)
	return bs
}

func parseMessages(bs []byte) ([]WlMessage, error) {
	var msgs []WlMessage
	var msg WlMessage
	var err error
	for {
		if len(bs) == 0 {
			return msgs, nil
		}
		msg, bs, err = parseOneMessage(bs)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
}

func marshMessages(msgs []WlMessage) []byte {
	var bs []byte
	for _, v := range msgs {
		bs = append(bs, marshOneMessage(v)...)
	}
	return bs
}

func parseFDs(oob []byte) ([]int, error) {
	scms, err := unix.ParseSocketControlMessage(oob)
	if err != nil {
		return nil, err
	}
	fdss := make([][]int, len(scms))
	total := 0
	for i, v := range scms {
		fdss[i], err = unix.ParseUnixRights(&v)
		if err != nil {
			return nil, err
		}
		total += len(fdss[i])
	}
	fds := make([]int, 0, total)
	for _, v := range fdss {
		for _, fd := range v {
			fds = append(fds, fd)
		}
	}

	return fds, nil
}

func marshFDs(fds []int) []byte {
	return unix.UnixRights(fds...)
}

type (
	WlFixed  float32
	WlString string
	WlObject uint32
	WlInt    int32
	WlUint   uint32
	WlNewId  uint32
	WlArray  []byte
	WlFd     uintptr
)

type unixMsg struct {
	pkt, oob []byte
}

func ReadMsg(conn *net.UnixConn) (*WlWireMessage, error) {
	pkt := make([]byte, 4096)
	oob := make([]byte, 4096)
	n, oobn, _, _, err := conn.ReadMsgUnix(pkt, oob)
	if err != nil {
		return nil, err
	}
	pkt = pkt[:n]
	oob = oob[:oobn]
	msgs, err := parseMessages(pkt)
	if err != nil {
		return nil, err
	}
	fds, err := parseFDs(oob)
	if err != nil {
		return nil, err
	}
	return &WlWireMessage{
		Messages: msgs,
		FDs:      fds,
	}, nil
}

func SendMsg(conn *net.UnixConn, wmsg *WlWireMessage) error {
	_, _, err := conn.WriteMsgUnix(marshMessages(wmsg.Messages), marshFDs(wmsg.FDs), nil)
	if err != nil {
		return err
	}
	return nil
}
