package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	// t, ok2 := c.GetQuery("latest_time")
	// token, ok := c.GetQuery("token")

	// var t time.Time = time.Now();

	// TODO: current version: feed videos without token check
	// generate response based on provided token and latest time

	// given the token and latest time, return the response
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}




