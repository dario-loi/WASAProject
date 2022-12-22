package api

import (
	"image"
	"time"
)

// Date type for times.
type Date string

// a Photo.
//
// The type interface will probably change as we will surely have to change the way
// we store photos, especially depending on the delivery system to the vue.js frontend.
//
// The current implementation is a simple wrapper around the image.Image type.
type Photo image.Image

// List of banned users.
type BanList struct {
	BannedUsers *UserList `json:"banned_users,omitempty"`
	Owner       *struct {
		ProfilePhoto *struct {
			Hash *string `json:"hash,omitempty"`
		} `json:"profile_photo,omitempty"`

		Username *Username `json:"username,omitempty"`
	} `json:"owner,omitempty"`
}

// The Biography of a WASAPhoto user.
type Biography struct {
	Birthdate  *Date `json:"birthdate,omitempty"`
	Birthplace *struct {
		City *string `json:"city,omitempty"`

		Coordinates *struct {
			Latitude *float32 `json:"latitude,omitempty"`

			Longitude *float32 `json:"longitude,omitempty"`
		} `json:"coordinates,omitempty"`

		Country *string `json:"country,omitempty"`

		State *string `json:"state,omitempty"`
	} `json:"birthplace,omitempty"`

	CurrentPlace *struct {
		Location *struct {
			City *string `json:"city,omitempty"`

			Coordinates *struct {
				Latitude *float32 `json:"latitude,omitempty"`

				Longitude *float32 `json:"longitude,omitempty"`
			} `json:"coordinates,omitempty"`

			Country *string `json:"country,omitempty"`

			State *string `json:"state,omitempty"`
		} `json:"location,omitempty"`

		Name *string `json:"name,omitempty"`
	} `json:"current_place,omitempty"`

	CurrentState *BiographyCurrentState `json:"current_state,omitempty"`

	Education *[]struct {
		EndDate *Date `json:"end_date,omitempty"`

		Name  *string `json:"name,omitempty"`
		Place *struct {
			City *string `json:"city,omitempty"`

			Coordinates *struct {
				Latitude *float32 `json:"latitude,omitempty"`

				Longitude *float32 `json:"longitude,omitempty"`
			} `json:"coordinates,omitempty"`

			Country *string `json:"country,omitempty"`

			State *string `json:"state,omitempty"`
		} `json:"place,omitempty"`

		StartDate *Date `json:"start_date,omitempty"`
	} `json:"education,omitempty"`

	Employment *[]struct {
		EndDate *Date `json:"end_date,omitempty"`

		Name  *string `json:"name,omitempty"`
		Place *struct {
			City *string `json:"city,omitempty"`

			Coordinates *struct {
				Latitude *float32 `json:"latitude,omitempty"`

				Longitude *float32 `json:"longitude,omitempty"`
			} `json:"coordinates,omitempty"`

			Country *string `json:"country,omitempty"`

			State *string `json:"state,omitempty"`
		} `json:"place,omitempty"`

		StartDate *Date `json:"start_date,omitempty"`
	} `json:"employment,omitempty"`
	Residence *struct {
		City *string `json:"city,omitempty"`

		Coordinates *struct {
			Latitude *float32 `json:"latitude,omitempty"`

			Longitude *float32 `json:"longitude,omitempty"`
		} `json:"coordinates,omitempty"`

		Country *string `json:"country,omitempty"`

		State *string `json:"state,omitempty"`
	} `json:"residence,omitempty"`

	ShortDescription *string `json:"short_description,omitempty"`
}

// A WASAPhoto comment.
type Comment struct {
	Author *User `json:"author,omitempty"`

	Body *string `json:"body,omitempty"`

	CreationTime *time.Time `json:"creation_time,omitempty"`

	ParentPost *SHA256hash `json:"parent_post,omitempty"`
}

// A WASAPhoto comment ID.
type CommentID struct {
	Token *SHA256hash `json:"token,omitempty"`
}

// A WASAPhoto comment list.
type CommentList struct {
	Comments *[]Comment `json:"comments,omitempty"`
}

// Error is a generic error response.
type Error struct {
	Code *int `json:"code,omitempty"`

	Message *string `json:"message,omitempty"`
}

// A WASAPhoto follow list.
type FollowList struct {
	FollowList *UserList `json:"follow-list,omitempty"`
	Owner      *struct {
		ProfilePhoto *struct {
			Hash *string `json:"hash,omitempty"`
		} `json:"profile_photo,omitempty"`

		Username *Username `json:"username,omitempty"`
	} `json:"owner,omitempty"`
}

// A WASAPhoto like list.
type LikesList struct {
	Likes *UserList `json:"likes,omitempty"`
	Owner *struct {
		AuthorIdentifier *UserID `json:"author_identifier,omitempty"`

		CommentsList *CommentList `json:"comments-list,omitempty"`

		CreatedAt *time.Time `json:"created_at,omitempty"`

		Description *string `json:"description,omitempty"`

		Id *SHA256hash `json:"id,omitempty"`

		Photo *Photo `json:"photo,omitempty"`
	} `json:"owner,omitempty"`
}

// A WASAPhoto post.
type Photopost struct {
	AuthorIdentifier *UserID `json:"author_identifier,omitempty"`

	CommentsList *CommentList `json:"comments-list,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`

	Description *string `json:"description,omitempty"`

	Id *SHA256hash `json:"id,omitempty"`

	Photo *Photo `json:"photo,omitempty"`
}

// A Physical location.
type Place struct {
	City *string `json:"city,omitempty"`

	Coordinates *struct {
		Latitude *float32 `json:"latitude,omitempty"`

		Longitude *float32 `json:"longitude,omitempty"`
	} `json:"coordinates,omitempty"`

	Country *string `json:"country,omitempty"`

	State *string `json:"state,omitempty"`
}

// A wrapper for a SHA256 hash, used to identify users, posts, etc.
type SHA256hash struct {
	Hash *string `json:"hash,omitempty"`
}

// A WASAPhoto stream.
type Stream struct {
	Posts *[]Photopost `json:"posts,omitempty"`
}

// A WASAPhoto user.
type User struct {
	ProfilePhoto *struct {
		Hash *string `json:"hash,omitempty"`
	} `json:"profile_photo,omitempty"`

	Username *Username `json:"username,omitempty"`
}

// A WASAPhoto user ID.
type UserID struct {
	Token *SHA256hash `json:"token,omitempty"`
}

// A WASAPhoto user list.
type UserList struct {
	Users *[]UserID `json:"users,omitempty"`
}

// A WASAPhoto username.
type Username struct {
	UsernameString *string `json:"username-string,omitempty"`
}
