package main

//QueryZeroType type
type QueryZeroType struct {
	Q            int         `json:"q"`
	Ok           bool        `json:"ok"`
	ErrorMessage string      `json:"e"`
	Sender       string      `json:"s"`
	Users        AllUserType `json:"u"`
}

//QueryOneType type
type QueryOneType struct {
	Q            int    `json:"q"`
	Ok           bool   `json:"ok"`
	ErrorMessage string `json:"e"`
	Sender       string `json:"s"`
	Receiver     string `json:"r"`
	Message      string `json:"m"`
}
