package main

//AllUserType represents response object model
type AllUserType struct {
	UserCount int      `json:"uc"`
	UserNames []string `json:"un"`
}
