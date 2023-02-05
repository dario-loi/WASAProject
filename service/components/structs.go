package components

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	Uname string `json:"username-string"`
}

func (u User) ToJSON() ([]byte, error) {
	return json.MarshalIndent(u, "", "  ")
}

type SHA256hash struct {
	Hash string `json:"hash"`
}

func (h SHA256hash) ToJSON() ([]byte, error) {

	wrapper := struct {
		Token SHA256hash `json:"token"`
	}{
		Token: h,
	}

	return json.MarshalIndent(wrapper, "", "  ")
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

// From https://stackoverflow.com/questions/23695479/how-to-format-timestamp-in-outgoing-json
// Adapted to use time.RFC3339, the standard from our API.

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {

	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))
	return []byte(stamp), nil
}

type Comment struct {
	Comment_ID   SHA256hash `json:"comment_id"`
	Username     string     `json:"author"`
	Body         string     `json:"body"`
	CreationTime JSONTime   `json:"creation-time"`
	Parent       SHA256hash `json:"parent_post"`
}

func (c Comment) ToJSON() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}

type Photo struct {
	Data string `json:"photo_data"`
	Desc string `json:"photo_desc"`
}

type Post struct {
	Photo_ID     SHA256hash `json:"photo_id"`
	Author_Name  User       `json:"author_name"`
	Description  string     `json:"description"`
	CreationTime JSONTime   `json:"created_at"`
}

type Stream struct {
	Posts []Post `json:"posts"`
}
