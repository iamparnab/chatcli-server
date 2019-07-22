package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

//Users info
func Users(res http.ResponseWriter, req *http.Request, allUsers map[string]net.Conn) {
	totalCount := GetUserCount(allUsers)
	respBody := AllUserType{
		UserCount: GetUserCount(allUsers),
		UserNames: GetUserNames(allUsers, totalCount),
	}
	res.Header().Set("Content-Type", "application/json")

	/**
	 * https://golang.org/pkg/encoding/json/#NewEncoder
	 */
	err := json.NewEncoder(res).Encode(respBody)
	if err != nil {
		fmt.Println(err)
	}
}
