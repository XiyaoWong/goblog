package service

import (
	"goblog/common"
	"goblog/serializer"
)

// CheckTokenService 检查token有效性
type CheckTokenService struct {
	Token string `json:"token"`
}

// Check .
func (service *CheckTokenService) Check() *serializer.Response {
	token, _, err := common.ParseToken(service.Token)
	if err != nil || !token.Valid {
		return &serializer.Response{Msg: "token已失效"}
	}
	return &serializer.Response{Msg: "ok"}
}
