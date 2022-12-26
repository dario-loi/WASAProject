package components

import (
	"database/sql"
	"encoding/json"
)

type User struct {
	Uname   Username   `json:"username-string"`
	PhotoID SHA256hash `json:"profile_photo"`
}

type Username struct {
	Username_string string `json:"username-string"`
}

func (u User) ToJSON() ([]byte, error) {
	return json.MarshalIndent(u, "", "  ")
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) ToJSON() ([]byte, error) {
	return json.MarshalIndent(e, "", "  ")
}

type SHA256hash struct {
	Hash sql.NullString `json:"hash"`
}

func (h SHA256hash) ToJSON() ([]byte, error) {
	return json.MarshalIndent(h, "", "  ")
}
