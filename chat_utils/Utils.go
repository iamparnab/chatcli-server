package chatutils

import "net"

//GetUserCount type
func GetUserCount(allUsers map[string]net.Conn) int {
	return len(allUsers)
}

//GetUserNames returns all uers
func GetUserNames(allUsers map[string]net.Conn, totalUsers int) []string {
	var userNames = make([]string, 0)
	for keys := range allUsers {
		userNames = append(userNames, keys)
	}
	return userNames
}
