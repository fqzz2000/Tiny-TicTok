package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fqzz2000/tiny-tictok/model"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	userIdRaw, _ := c.Get("UserId")
	videoRaw := c.Query("video_id")
	actionRaw := c.Query("action_type")
	fmt.Println(userIdRaw)
	userIdString := fmt.Sprint(userIdRaw)
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	
	if videoRaw == "" {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Invalid Video ID"})
		return
	}
	if actionRaw == "" {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Invalid Action Type"})
		return
	}
	videoId, _ := strconv.ParseInt(videoRaw, 10, 64)
	action, _ := strconv.ParseInt(actionRaw, 10, 64)

	if action == 1 {
	// TODO: check if there are error in the database
	model.NewLikeDAO().AddNewLike(&model.LikeDB{
		UserID: userId,
		VideoID: videoId,
		LikeTime: time.Now(),
	})
} else if action == 2 {
	model.NewLikeDAO().DeleteLike(userId, videoId)
}
	c.JSON(http.StatusOK, Response{StatusCode: 0})
	
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	usrIdRaw, _ := c.Get("UserId")
	usrId, _ := strconv.ParseInt(fmt.Sprint(usrIdRaw), 10, 64)
	
	var videoList []model.VideoDB
	likeList := model.NewLikeDAO().GetLikeListByUserID(usrId)
	fmt.Println(len(likeList))
	for i:=0; i<len(likeList); i++ {
		vid := likeList[i].VideoID
		videoList = append(videoList, model.NewVideoDAO().QueryVideoByID(vid))
	}
	vList := DecorateVideos(videoList, usrId)

	
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: vList,
	})

}
