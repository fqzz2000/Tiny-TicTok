package controller

import "time"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

// Structs for mySQL table
type UserDB struct {
	UserID 	 uint64 `gorm:"primaryKey column: user_id"`
	UserName string `gorm:"column: user_name"`
	UserPswd string `gorm:"column: user_pswd"`
}

func (*UserDB) TableName() string {
	return "users"
}

type VideoDB struct {
	VideoID      uint64		`gorm:"primaryKey column:video_id"`
	VideoTitle   string		`gorm:"column:video_title"`
	VideoDesc    string		`gorm:"column:video_desc"`
	VideoOwner   uint64		`gorm:"column:video_owner"`
	VideoCrtTime time.Time	`gorm:"column:video_crt_time"`
	VideoFile	 string		`gorm:"column:video_file"`
	CoverFile	 string		`gorm:"column:cover_file"`
}

func (*VideoDB) TableName() string {
	return "videos"
}