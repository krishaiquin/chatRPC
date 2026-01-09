package transport

import (
	"bytes"
	"chatRPC/dlog"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync/atomic"
	"time"
)

// call to request
func Call(to string, funcName string, args []byte) []byte {
	//create UDP endpoint from string address
	toAddr, err := net.ResolveUDPAddr("udp", to)
	if err != nil {
		panic(err)
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	//write to buffer
	seqNum := seq.Add(1)
	pendingRequest[seqNum] = make(chan []byte, 1)
	if err = binary.Write(buf, binary.NativeEndian, seqNum); err != nil {
		panic(err)
	}
	if err = binary.Write(buf, binary.NativeEndian, request); err != nil {
		panic(err)
	}
	if err = binary.Write(buf, binary.NativeEndian, uint32(len(funcName))); err != nil {
		panic(err)
	}
	if err = binary.Write(buf, binary.NativeEndian, []byte(funcName)); err != nil {
		panic(err)
	}
	if err = binary.Write(buf, binary.NativeEndian, args); err != nil {
		panic(err)
	}

	//send request
	dlog.Printf("Sending request#%d to  [%s] - %s(%s)", seqNum, to, funcName, args)
	n, err := conn.WriteToUDP(buf.Bytes(), toAddr)
	if err != nil {
		panic(err)
	}
	if n != len(buf.Bytes()) {
		panic(fmt.Errorf("truncated send. Sent: %d, original: %d", n, len(args)))
	}

	//start the timeout timer
	timeout := make(chan bool)
	go func() {
		time.Sleep(responseTimeout)
		timeout <- true
	}()

	//see which happens first: timeout or receiving the result
	select {
	case <-timeout:
		dlog.Printf("Request has timed out! Exiting...")
	case res := <-pendingRequest[seqNum]:
		dlog.Printf("Result for request#%d has been sent to the application\n", seqNum)
		return res
	}

	return nil

}

// listen for request
func Listen() {

	for {
		//make buffer for requests
		res := make([]byte, 2048)
		n, from, err := conn.ReadFromUDP(res)
		if err != nil {
			panic(err)
		}
		res = res[:n]
		//put request to buf
		buf := bytes.NewBuffer(res)
		//extract data from buf
		var seqNum uint32
		var packetType byte
		if err = binary.Read(buf, binary.NativeEndian, &seqNum); err != nil {
			panic(err)
		}
		if err = binary.Read(buf, binary.NativeEndian, &packetType); err != nil {
			panic(err)
		}
		switch packetType {
		case request:
			dlog.Printf("Received request#%d from %s\n", seqNum, from.String())
			var funcLength uint32

			err = binary.Read(buf, binary.NativeEndian, &funcLength)
			funcName := make([]byte, funcLength)
			if err != nil {
				panic(err)
			}
			err = binary.Read(buf, binary.NativeEndian, &funcName)
			if err != nil {
				panic(err)
			}
			dlog.Printf("function name is: %s\n", funcName)
			serverStub := serverStubRegistry[string(funcName)]
			//dispatch
			response := serverStub(buf.Bytes())

			//make buffer
			buffer := bytes.NewBuffer(make([]byte, 0))
			if err = binary.Write(buffer, binary.NativeEndian, seqNum); err != nil {
				panic(err)
			}
			if err = binary.Write(buffer, binary.NativeEndian, result); err != nil {
				panic(err)
			}
			if err = binary.Write(buffer, binary.NativeEndian, response); err != nil {
				panic(err)
			}
			dlog.Printf("Sending response to [%s] - %s", from.String(), response)
			//write the response to the connection
			_, err = conn.WriteToUDP(buffer.Bytes(), from)
			if err != nil {
				panic(err)
			}

		case result:
			dlog.Printf("Received result for request#%d\n", seqNum)
			data, err := io.ReadAll(buf)
			if err != nil {
				panic(err)
			}
			pendingRequest[seqNum] <- data
		}

	}

}

func GetAddress() string {
	return conn.LocalAddr().String()
}

func RegisterServerStub(funcName string, serverStub func([]byte) []byte) {
	serverStubRegistry[funcName] = serverStub
}

func init() {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}
	//create the UDPConn
	conn, err = net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	//initialize
	serverStubRegistry = make(map[string]func([]byte) []byte)
	pendingRequest = make(map[uint32]chan []byte)
	responseTimeout = 15 * time.Second

}

var conn *net.UDPConn
var serverStubRegistry map[string]func([]byte) []byte
var pendingRequest map[uint32]chan []byte
var seq atomic.Uint32
var responseTimeout time.Duration

const (
	request byte = iota
	result
)
