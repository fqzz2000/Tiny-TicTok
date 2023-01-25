package controller

import (
	"net/http"
	"time"

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
		NextTime:  t.Unix(),
	})
}
// return a list of videos that can be returned to the front end
func decorateVideos(videoDBs []model.VideoDB) []Video {
	var ans []Video;
	for _, v := range videoDBs {
		ans = append(ans, Video{
			Id: v.VideoID,
			Author: User{},
			PlayUrl: v.VideoFile,
			CoverUrl: v.CoverFile,
			FavoriteCount: 12,
			CommentCount: 18,
			IsFavorite: false, 
		})
	}
	return ans
}




