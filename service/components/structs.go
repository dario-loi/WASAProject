package components

import (
	"encoding/json"
)

type User struct {
	Uname string `json:"username-string"`
}

func (u User) ToJSON() ([]byte, error) {
	return json.MarshalIndent(u, "", "  ")
}

type SHA256hash struct {
	Hash string `json:"username-string"`
}

func (h SHA256hash) ToJSON() ([]byte, error) {
	return json.MarshalIndent(h, "", "  ")
}

type Profile struct {
	Username string       `json:"username-string"`
	Photos   []SHA256hash `json:"photos"`
}

func (p Profile) ToJSON() ([]byte, error) {
	return json.MarshalIndent(p, "", "  ")
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) ToJSON() ([]byte, error) {
	return json.MarshalIndent(e, "", "  ")
}

type IDList struct {
	IDs []SHA256hash `json:"ids"`
}

func (i IDList) ToJSON() ([]byte, error) {
	if len(i.IDs) == 0 {
		return []byte("[]"), nil
	}

	return json.MarshalIndent(i, "", "  ")
}
