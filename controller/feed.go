package controller

import (
	"net/http"
	"path/filepath"
	"time"

	"github.com/fqzz2000/tiny-tictok/config"
	"github.com/fqzz2000/tiny-tictok/model"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}
type user struct {
	UserID uint64 `gorm:"primaryKey column:user_id"`
	UserName string `gorm:"column:user_name"`
	UserPswd string `gorm:"column:user_pswd"`
}
// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// t, ok2 := c.GetQuery("latest_time")
	// token, ok := c.GetQuery("token")

	var t time.Time = time.Date(2010, time.April, 2, 2, 2, 2, 2, time.UTC);

	// TODO: current version: feed videos without token check
	// generate response based on provided token and latest time

	// given the token and latest time, return the response
	var videos []Video;
	videos = decorateVideos(model.NewVideoDAO().QueryVideoBeforeTime(t, 30))

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  t.Unix(), // TODO: need to set next time to proper value
	})
}
// decorate user by user id
// TODO: add isFollow after register function completed
func decorateUser(id int64) User{
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
func decorateVideos(videoDBs []model.VideoDB) []Video {
	var ans []Video;
	for _, v := range videoDBs {
		ans = append(ans, Video{
			Id: v.VideoID,
			Author: decorateUser(int64(v.VideoOwner)),
			PlayUrl: config.Info.StaticSourcePath + "/videos/" + v.VideoFile,
			CoverUrl: filepath.Join(config.Info.StaticSourcePath, "covers", v.CoverFile),
			FavoriteCount: model.NewLikeDAO().CountLikesByVideoID(v.VideoID),
			CommentCount: model.NewCommentDAO().CountCommentsByVideoID(v.VideoID),
			IsFavorite: false, 
		})
	}
	return ans
}




