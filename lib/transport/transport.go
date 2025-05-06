package transport

import (
	message "chatRPC/message/rpc/serverStub"
	"fmt"
	"net"
)

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
		panic(fmt.Errorf("Truncated send. Sent: %d, original: %d", n, len(args)))
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
		//Dispath to server
		fmt.Println("Received Request!")
		//call the serverStub
		response := message.Echo(args)
		//--------------------

		//write the response to the connection
		_, err = conn.WriteToUDP(response, from)
		if err != nil {
			panic(err)
		}

	}

}

func LocalAddr() string {
	return conn.LocalAddr().String()
}

var conn *net.UDPConn
