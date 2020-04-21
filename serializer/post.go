// Package serializer provides ...
package serializer

import (
	"goblog/model"
	"math"
)

// Post ...
type Post struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	CreateAt int64  `json:"create_at"`
	UpdateAT int64  `json:"update_at"`
}

// BuildPostDetailResponse 帖子详情
func BuildPostDetailResponse(post model.Post) Response {
	return Response{
		Data: Post{
			ID:       post.ID,
			Title:    post.Title,
			Body:     post.Body,
			CreateAt: post.CreatedAt.Unix(),
			UpdateAT: post.UpdatedAt.Unix(),
		},
	}
}

// BuildPostOutline 帖子概要
func BuildPostOutline(post model.Post) interface{} {
	bodyLength := len(post.Body)
	pos := math.Min(200, float64(bodyLength))
	return Post{
		ID:       post.ID,
		Title:    post.Title,
		Body:     post.Body[:int(pos)],
		CreateAt: post.CreatedAt.Unix(),
		UpdateAT: post.UpdatedAt.Unix(),
	}
}

// BuildPostList 帖子列表
func BuildPostList(items []model.Post) (posts []interface{}) {
	for _, item := range items {
		post := BuildPostOutline(item)
		posts = append(posts, post)
	}
	return posts
}
