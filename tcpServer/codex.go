package tcpServer

import (
	"net"
	"github.com/leesper/tao"
	"fmt"
	"io/ioutil"
	"errors"
)

// TypeLengthValueCodec defines a special codec.
// Format: type-length-value |4 bytes|4 bytes|n bytes <= 8M|
type StringValue struct{}

// Decode decodes the bytes data into Message
func (codec StringValue) Decode(raw net.Conn) (tao.Message, error) {
	msgBytes ,num := readAll(raw)
	fmt.Printf("Decode  => res:%v num:%v\n",string(msgBytes),num)

	if num == 0 {
		return nil, errors.New("length cannot be 0")
	}

	// deserialize message from bytes
	unmarshaler := tao.GetUnmarshalFunc(1)

	if unmarshaler == nil {
		return nil, tao.ErrUndefined(1)
	}

	return unmarshaler(msgBytes)
}

// Encode encodes the message into bytes data.
func (codec StringValue) Encode(msg tao.Message) ([]byte, error) {
	data, err := msg.Serialize()
	fmt.Printf("Encode:%v err:", string(data), err.Error())
	if err != nil {
		return nil, err
	}

	return data, nil
}

func readAll(conn net.Conn) ([]byte, int) {
	buf, err := ioutil.ReadAll(conn)
	if err != nil {
		// Error Handler
		return nil, 0
	}
	fmt.Printf("readAll  =>  %v  %v\n", string(buf), len(buf))
	return  buf, len(buf)
}