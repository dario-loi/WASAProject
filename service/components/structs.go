package components

import "encoding/json"

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

const InternalServerError string = "{\"code\": 500, \"message\": \"Internal Server Error\"}"
const BadRequestError string = "{\"code\": 400, \"message\": \"Bad Request\"}"

func (e Error) ToJSON() ([]byte, error) {
	return json.MarshalIndent(e, "", "  ")
}

type SHA256hash struct {
	Hash string `json:"hash"`
}

func (h SHA256hash) ToJSON() ([]byte, error) {
	return json.MarshalIndent(h, "", "  ")
}
