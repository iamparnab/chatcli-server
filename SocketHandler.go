package main

import (
	"encoding/json"
	"fmt"
	"net"
)

// SocketHandler handles socket sonnections
func SocketHandler(conn net.Conn, allUsers map[string]net.Conn) {
	for {
		userData := QueryOneType{}

		/**
		 * https://golang.org/pkg/encoding/json/#NewDecoder
		 */
		decoder := json.NewDecoder(conn)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&userData)

		if err != nil {
			go fmt.Println("Error", err, conn.RemoteAddr())
			respObj := QueryZeroType{
				Q:            0,
				Ok:           false,
				ErrorMessage: "400: Bad Request",
			}
			json.NewEncoder(conn).Encode(respObj)
		} else {
			switch userData.Q {
			case 0:
				addNewEntry(userData, allUsers, conn)
				totalUsers := GetUserCount(allUsers)
				userNames := GetUserNames(allUsers, totalUsers)
				userObj := AllUserType{
					UserCount: totalUsers,
					UserNames: userNames,
				}
				respObj := QueryZeroType{
					Q:     0,
					Ok:    true,
					Users: userObj,
				}
				json.NewEncoder(conn).Encode(respObj)
			case 1:
				receiverCon, ok := allUsers[userData.Receiver]
				if !ok {
					totalUsers := GetUserCount(allUsers)
					userNames := GetUserNames(allUsers, totalUsers)
					respObj := QueryZeroType{
						Q:            0,
						Ok:           false,
						Sender:       userData.Sender,
						ErrorMessage: "404: User " + userData.Receiver + " does not exist",
						Users: AllUserType{
							UserCount: totalUsers,
							UserNames: userNames,
						},
					}
					json.NewEncoder(conn).Encode(respObj)
				} else {
					respObj := QueryOneType{
						Q:        1,
						Ok:       true,
						Sender:   userData.Sender,
						Receiver: userData.Receiver,
						Message:  userData.Message,
					}
					json.NewEncoder(receiverCon).Encode(respObj)
				}
			}
		}
	}
}

func addNewEntry(userData QueryOneType, allUsers map[string]net.Conn, conn net.Conn) {
	allUsers[userData.Sender] = conn
}
