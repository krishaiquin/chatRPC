package transport

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

// call to request
func Call(to string, funcName string, args []byte) []byte {
	//create UDP endpoint from string address
	toAddr, err := net.ResolveUDPAddr("udp", to)
	if err != nil {
		panic(err)
	}

	buf := new(bytes.Buffer)
	//write to buffer
	err = binary.Write(buf, binary.NativeEndian, uint32(len(funcName)))
	if err != nil {
		panic(err)
	}
	err = binary.Write(buf, binary.NativeEndian, []byte(funcName))
	if err != nil {
		panic(err)
	}
	err = binary.Write(buf, binary.NativeEndian, args)
	if err != nil {
		panic(err)
	}
	//write buf to the desired address
	n, err := conn.WriteToUDP(args, toAddr)
	if err != nil {
		panic(err)
	}
	if n != len(args) {
		panic(fmt.Errorf("truncated send. Sent: %d, original: %d", n, len(args)))
	}
	//make buffer or the response
	response := make([]byte, 2048)
	n, _, err = conn.ReadFromUDP(response)
	if err != nil {
		panic(err)
	}
	//return the buf with n-bytes read from UDP
	return response[:n]

}

// listen for request
func Listen() {
	for {
		//make buffer for requests
		request := make([]byte, 2048)
		n, from, err := conn.ReadFromUDP(request)
		if err != nil {
			panic(err)
		}
		request = request[:n]
		//put request to buf
		buf := bytes.NewBuffer(request)
		//extract data from buf
		var funcLength uint32
		var args []byte
		err = binary.Read(buf, binary.NativeEndian, &funcLength)
		if err != nil {
			panic(err)
		}
		err = binary.Read(buf, binary.NativeEndian, &args)
		if err != nil {
			panic(err)
		}
		funcName := args[:funcLength]
		args = args[funcLength:]
		serverStub := serverStubRegistry[string(funcName)]
		//dispatch
		response := serverStub(args)

		//write the response to the connection
		_, err = conn.WriteToUDP(response, from)
		if err != nil {
			panic(err)
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

	//initialize function registry
	serverStubRegistry = make(map[string]func([]byte) []byte)

}

var conn *net.UDPConn
var serverStubRegistry map[string]func([]byte) []byte
