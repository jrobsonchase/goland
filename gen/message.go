package gen

import (
	"errors"
	"reflect"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

type WlHeader struct {
	Id   uint32
	Op   uint16
	Size uint16
}

func ReadHeader(bs []byte) (WlHeader, []byte, error) {
	if len(bs) < 8 {
		return WlHeader{}, nil, errors.New("ReadHeader: not enough data")
	}

	headerBytes := [8]byte{}
	for i, v := range bs[:8] {
		headerBytes[i] = v
	}

	return *(*WlHeader)(unsafe.Pointer(&headerBytes)), bs[8:], nil
}

type WlMessage struct {
	WlHeader
	Data []uint32
}

func ReadMessage(bs []byte) (msg WlMessage, rest []byte, err error) {
	msg.WlHeader, _, err = ReadHeader(bs)
	if err != nil {
		return
	}
	if len(bs) < int(msg.Size) {
		err = errors.New("ReadMessage: actual size does not match header size")
		return
	}
	msg.Data = make([]uint32, (msg.Size-8)/4)
	header := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	islice := *(*[]uint32)(unsafe.Pointer(header))
	copy(msg.Data, islice[2:])
	rest = bs[msg.Size:]
	return
}

func ParseMessages(bs []byte) ([]WlMessage, error) {
	var msgs []WlMessage
	var msg WlMessage
	var err error
	for {
		if len(bs) == 0 {
			return msgs, nil
		}
		msg, bs, err = ReadMessage(bs)
		if err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}
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

func StringArg(ws []uint32) WlString {
	bs := make([]byte, len(ws)*4)
	header := (*reflect.SliceHeader)(unsafe.Pointer(&ws))
	copy(bs, *(*[]byte)(unsafe.Pointer(header)))
	return WlString(bs)
}
func StringMsg(str WlString) []uint32 {
	bs := []byte(str)
	ws := make([]uint32, len(bs)/4)
	header := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	copy(ws, *(*[]uint32)(unsafe.Pointer(header)))
	return ws
}

func ArrayArg(ws []uint32) WlArray {
	bs := make([]byte, ws[0])
	header := (*reflect.SliceHeader)(unsafe.Pointer(&ws))
	header.Len *= 4
	header.Cap *= 4
	copy(bs, (*(*[]byte)(unsafe.Pointer(header)))[4:])
	return bs
}

func ArrayMsg(arr WlArray) []uint32 {
	l := len(arr)
	pad := 0
	if l%4 != 0 {
		arr = append(arr, make([]byte, 4-l%4)...)
		pad = 1
	}
	spew.Dump(arr)
	ws := make([]uint32, l/4+pad+1)
	ws[0] = uint32(l)
	header := (*reflect.SliceHeader)(unsafe.Pointer(&arr))
	header.Len /= 4
	header.Cap /= 4
	copy(ws[1:], *(*[]uint32)(unsafe.Pointer(header)))
	return ws
}
