package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/fqzz2000/tiny-tictok/middleware"
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
	tme, _ := c.GetQuery("latest_time")
	token, _ := c.GetQuery("token")
	usrId := int64(0)
	if (token != "") {
		claims, ok := middleware.ParseToken(token)
		if ok {
			usrId = claims.UserId
		}
	}
	var t time.Time = time.Date(2010, time.April, 2, 2, 2, 2, 2, time.UTC);
	if tme != "" {
		// parse time into unix time
		i, _ := strconv.ParseInt(tme, 10, 64)
		t = time.Unix(i, 0)
	}

	// TODO: current version: feed videos without token check
	// generate response based on provided token and latest time

	// given the token and latest time, return the response
	var videos []Video;
	videos = DecorateVideos(model.NewVideoDAO().QueryVideoBeforeTime(t, 30), usrId)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  t.Unix(), // TODO: need to set next time to proper value
	})
}





