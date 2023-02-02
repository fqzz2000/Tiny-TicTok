package controller

import (
	"path/filepath"

	"github.com/fqzz2000/tiny-tictok/config"
	"github.com/fqzz2000/tiny-tictok/model"
)
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

// decorate user by user id
// TODO: add isFollow after register function completed
func DecorateUser(id int64) User{
	usrdb := model.NewUserDAO().QueryUserById(id)
	ans := User {
		Id: int64(usrdb.UserID),
		Name: usrdb.UserName, 
		FollowCount: model.NewRelationDAO().CountRelationsByFansID(int64(usrdb.UserID)),
		FollowerCount: model.NewRelationDAO().CountRelationsByFollowerID(int64(usrdb.UserID)),
	}
	return ans
}

// return a list of videos that can be returned to the front end
func DecorateVideos(videoDBs []model.VideoDB, userID int64) []Video {
	var ans []Video;
	for _, v := range videoDBs {
		isFavorate := false
		if userID > 0 {
			isFavorate = model.NewLikeDAO().GetIfUserLikeVideo(userID, v.VideoID)
		}
		ans = append(ans, Video{
			Id: v.VideoID,
			Author: DecorateUser(int64(v.VideoOwner)),
			PlayUrl: config.Info.StaticSourcePath+ "videos/"+ v.VideoFile,
			CoverUrl: filepath.Join(config.Info.StaticSourcePath, "covers", v.CoverFile),
			FavoriteCount: model.NewLikeDAO().CountLikesByVideoID(v.VideoID),
			CommentCount: model.NewCommentDAO().CountCommentsByVideoID(v.VideoID),
			IsFavorite: isFavorate, 
		})
	}
	return ans
}