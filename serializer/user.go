// Package serializer provides ...
package serializer

import (
	"goblog/model"
)

// User ...
type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// BuildUserInfoResponse ..
func BuildUserInfoResponse(user model.User) Response {
	return Response{
		Data: User{
			ID:       user.ID,
			Username: user.Username,
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
		},
	}
}
