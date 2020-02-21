package model

import (
	"encoding/json"
	"time"
)

const (
	// TableComment Table name
	TableComment = `comments`
)

// Comment Comment table definition
type Comment struct {
	ID             string `db:"id" json:"id"`
	MemberAsString string `db:"member" json:"member"`
	Member         Member `db:"-" json:"-"`
	Organization   string `db:"organization" json:"organization"`
	Message        string `db:"message" json:"message"`
	Visible        int    `db:"visible" json:"visible"`
	CreatedAt      string `db:"created_at" json:"created_at"`
	UpdatedAt      string `db:"updated_at" json:"updated_at"`
	DeletedAt      string `db:"deleted_at" json:"deleted_at"`
}

// Member Member definition
type Member struct {
	ID      int64  `json:"id"`
	Name    string `json:"login"`
	Profile string `json:"avatar_url"`
}

// NewComment ...
func NewComment() *Comment {
	return &Comment{
		ID:        GenerateUUID(),
		Visible:   1,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
}

// Prepare ...
func (me *Comment) Prepare() {
	me.MemberAsString = JSON(me.Member)
	return
}

// Format ...
func (me *Comment) Format() {
	json.Unmarshal([]byte(me.MemberAsString), &me.Member)
}
