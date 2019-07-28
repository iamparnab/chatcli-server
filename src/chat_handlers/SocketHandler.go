package chathandlers

import (
	chatmodels "chat_models"
	chatutils "chat_utils"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"regexp"
)

// SocketHandler handles socket sonnections
func SocketHandler(conn net.Conn, allUsers map[string]net.Conn) {
	for {
		userData := chatmodels.QueryOneType{}
		tempBuffer := make([]byte, 1024)
		n, err := conn.Read(tempBuffer)
		// Check for client disconnect
		if err == io.EOF {
			go fmt.Println("Client ", conn.RemoteAddr(), " is disconnected")
			break
		}

		//check for HTTP request
		tempBuffer = tempBuffer[0:n]
		if matched, err := regexp.Match(`Connection:\s*keep-alive`, tempBuffer); matched {
			if err != nil {
				go fmt.Println(err)
			}
			break
		}

		go fmt.Println("Source: ", conn.RemoteAddr(), " Byte size: ", n)
		if err != nil {
			go fmt.Println("Error", err, conn.RemoteAddr())
			respObj := chatmodels.QueryZeroType{
				Q:            0,
				Ok:           false,
				ErrorMessage: "400: Bad Request",
			}
			if encodedData, err := json.Marshal(respObj); err != nil {
				writeToSocket(conn, encodedData)
			}
		} else {
			json.Unmarshal(tempBuffer, &userData)

			switch userData.Q {
			case 0:
				addNewEntry(userData, allUsers, conn)
				totalUsers := chatutils.GetUserCount(allUsers)
				userNames := chatutils.GetUserNames(allUsers, totalUsers)
				userObj := chatmodels.AllUserType{
					UserCount: totalUsers,
					UserNames: userNames,
				}
				respObj := chatmodels.QueryZeroType{
					Q:     0,
					Ok:    true,
					Users: userObj,
				}

				if encodedData, err := json.Marshal(respObj); err == nil {
					writeToSocket(conn, encodedData)
				}

			case 1:
				receiverCon, ok := allUsers[userData.Receiver]
				if !ok {
					totalUsers := chatutils.GetUserCount(allUsers)
					userNames := chatutils.GetUserNames(allUsers, totalUsers)
					respObj := chatmodels.QueryZeroType{
						Q:            0,
						Ok:           false,
						Sender:       userData.Sender,
						ErrorMessage: "404: User " + userData.Receiver + " does not exist",
						Users: chatmodels.AllUserType{
							UserCount: totalUsers,
							UserNames: userNames,
						},
					}
					if encodedData, err := json.Marshal(respObj); err == nil {
						writeToSocket(conn, encodedData)
					}
				} else {
					respObj := chatmodels.QueryOneType{
						Q:        1,
						Ok:       true,
						Sender:   userData.Sender,
						Receiver: userData.Receiver,
						Message:  userData.Message,
					}
					if encodedData, err := json.Marshal(respObj); err == nil {
						writeToSocket(receiverCon, encodedData)
					}
				}
			}
		}
	}
	conn.Close()
}

func addNewEntry(userData chatmodels.QueryOneType, allUsers map[string]net.Conn, conn net.Conn) {
	allUsers[userData.Sender] = conn
}
func writeToSocket(conn net.Conn, byteArray []byte) {
	n, err := conn.Write(byteArray)
	if err != nil {
		go fmt.Println(err)
	}
	go fmt.Println("Written ", n, " bytes to socket")
}
