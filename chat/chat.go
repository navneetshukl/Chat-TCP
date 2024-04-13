package chat

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

var Str string

type Server struct {
	listner      net.Listener
	connections  []net.Conn    // Will store the connection of total clients
	messages     chan Message  // this channel will listen for message
	addClient    chan net.Conn // this channel will add new client
	removeClient chan net.Conn // this channel will listen for removing the connection

	// Group related channels

	createGroup chan net.Conn
	addGroup    chan net.Conn
	exitGroup   chan net.Conn
}

type Message struct {
	conn    net.Conn
	message []byte
}

// NewServer function will create the instance of server type and return
func NewServer(port string) *Server {
	listner, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal("Error in starting the server ", err)
	}

	s := &Server{
		listner:      listner,
		connections:  make([]net.Conn, 0),
		messages:     make(chan Message),
		addClient:    make(chan net.Conn),
		removeClient: make(chan net.Conn),
		createGroup:  make(chan net.Conn),
		addGroup:     make(chan net.Conn),
		exitGroup:    make(chan net.Conn),
	}
	return s
}

func (s *Server) Start() {
	go s.startChannels()

	for {
		conn, err := s.listner.Accept()
		if err != nil {
			log.Fatal("Error in Accepting the connection ", err)
		}

		s.addClient <- conn
		conn.Write([]byte(Str))

		fmt.Println("Address is ", conn.RemoteAddr())
		go s.handleRequest(conn)
	}
}

// This function will simply listen for above three channels
func (s *Server) startChannels() {
	for {
		select {
		case message := <-s.messages:
			s.broadcastMessage(&message)
		case newClient := <-s.addClient:
			s.connections = append(s.connections, newClient)
			fmt.Println("Total connected clients are ", len(s.connections))
		case deadClient := <-s.removeClient:
			s.removeConn(deadClient)
			fmt.Println("Total connected clients are ", len(s.connections))
		}
	}
}

// This will handle the TCP request and perform the action accordingly

func (s *Server) handleRequest(conn net.Conn) {
	for {
		message := make([]byte, 1024)
		_, err := conn.Read(message)

		if err != nil {
			if err == io.EOF {
				s.removeClient <- conn
				fmt.Println("Client removed from address is ", conn.RemoteAddr())
				conn.Close()
				return

			}
			log.Fatal("Error in HandleRequest is ", err)
		}
		str := string(message)
		str = strings.TrimSpace(str)

		// fmt.Println(len(str))

		if str[0] == '1' {

			fmt.Println("Create group command")

		} else if str[0] == '2' {
			fmt.Println("Add group command")
		} else if str[0] == '3' {
			fmt.Println("Exit group command")
		} else {
			m := Message{
				conn:    conn,
				message: message,
			}
			s.messages <- m
		}
	}

}

// this function will send message to all clients
func (s *Server) broadcastMessage(m *Message) {
	for _, conn := range s.connections {

		/* Here i am handling the case that message from
		same client should not be echoed back to the client
		*/

		if m.conn != conn {
			_, err := conn.Write(m.message)
			if err != nil {
				log.Fatal("Error in writing to address ", conn.RemoteAddr())
			}
		}
	}
}

// This function will remove the connection
func (s *Server) removeConn(conn net.Conn) {
	var i int
	for i = range s.connections {
		if s.connections[i] == conn {
			break
		}
	}

	s.connections = append(s.connections[:i], s.connections[i+1:]...)
}
