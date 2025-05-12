package transport

import (
	nodeset "chatRPC/nodeset/rpc/serverStub"
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
	//write data to the desired address
	n, err := conn.WriteToUDP(args, toAddr)
	if err != nil {
		panic(err)
	}
	if n != len(args) {
		panic(fmt.Errorf("truncated send. Sent: %d, original: %d", n, len(args)))
	}
	//make buffer for the response
	buf := make([]byte, 2048)
	nBytes, _, err := conn.ReadFromUDP(buf)
	if err != nil {
		panic(err)
	}
	//return the buf with n-bytes read from UDP
	return buf[:nBytes]

}

// listen for request
func Listen() {
	for {
		//make buffer for requests
		buf := make([]byte, 2048)
		n, from, err := conn.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}
		args := buf[:n]

		//----can be modularized---
		//Dispatch to the function name
		response := nodeset.Add(args)
		//--------------------

		// //write the response to the connection
		_, err = conn.WriteToUDP(response, from)
		if err != nil {
			panic(err)
		}

	}

}

func GetAddress() string {
	return conn.LocalAddr().String()
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
}

var conn *net.UDPConn
