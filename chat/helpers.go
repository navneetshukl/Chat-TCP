package chat

import (
	"fmt"
	"net"
)

func (s *Server) createGroups(conn net.Conn) {
	grpName := fmt.Sprintf("Group-%d", s.groupCount)
	if len(s.groups[grpName]) == 3 {
		s.groupCount++
		grpName = fmt.Sprintf("Group-%d", s.groupCount)
	}
	s.groups[grpName] = append(s.groups[grpName], conn)
	s.groupName[conn] = grpName
	fmt.Println("Total number of groups are ", len(s.groups))

	for key, value := range s.groups {
		fmt.Println("Key and value are ", key, value)
	}

}
