// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net"
// )

// type Message struct {
// 	conn    net.Conn
// 	message []byte
// }

// // Will store the connection of total clients
// var connections []net.Conn

// // this channel will listen for messages
// var messages = make(chan Message)

// // this channel will add new client
// var addClient = make(chan net.Conn)

// // this channel will listen for removing the connection
// var removeClient = make(chan net.Conn)

// func main() {

// 	server, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer server.Close()

// 	go startChannels()

// 	for {
// 		conn, err := server.Accept()
// 		if err != nil {
// 			log.Fatal("Error in Accepting connection ", err)
// 		}

// 		addClient <- conn

// 		fmt.Println("Address is ", conn.RemoteAddr())

// 		go handleRequest(conn)
// 	}

// }

// // This function will simply listen for above three channels
// func startChannels() {
// 	for {
// 		select {
// 		case message := <-messages:
// 			broadcastMessage(&message)
// 		case newClient := <-addClient:
// 			connections = append(connections, newClient)
// 			fmt.Println("Total connected clients are ", len(connections))
// 		case deadClient := <-removeClient:
// 			removeConn(deadClient)
// 			fmt.Println("Total coonected clients are ", len(connections))

// 		}
// 	}
// }

// // This will handle the TCP request and perform the action accordingly
// func handleRequest(conn net.Conn) {
// 	for {
// 		message := make([]byte, 1024)

// 		_, err := conn.Read(message)
// 		if err != nil {
// 			if err == io.EOF {
// 				removeClient <- conn
// 				fmt.Println("Client removed from address is ", conn.RemoteAddr())
// 				conn.Close()
// 				return
// 			}
// 			log.Fatal("Error in Handlerequest is ", err)
// 		}
// 		m := Message{
// 			conn:    conn,
// 			message: message,
// 		}
// 		messages <- m
// 	}
// }

// // this function will send message to all clients
// func broadcastMessage(m *Message) {
// 	for _, conn := range connections {

// 		/* Here i am handling the case that message from
// 		same client should not be echoed back to the client
// 		*/
// 		if m.conn != conn {
// 			_, err := conn.Write(m.message)

// 			if err != nil {
// 				log.Fatal("Error in writing to address ", conn.RemoteAddr())
// 			}
// 		}

// 	}
// }

// // This function will remove the connection
// func removeConn(conn net.Conn) {
// 	var i int

// 	for i = range connections {
// 		if connections[i] == conn {
// 			break
// 		}
// 	}

// 	connections = append(connections[:i], connections[i+1:]...)
// }

package main

import "github.com/chat/chat"

func main() {
	server := chat.NewServer(":8080")
	server.Start()
}
