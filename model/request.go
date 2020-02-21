package model

// Request Request table definition
type Request struct {
	Member  Member `json:"member"`
	Message string `json:"comment"`
}
