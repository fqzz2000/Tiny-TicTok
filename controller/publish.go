package controller

import (
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/fqzz2000/tiny-tictok/config"
	"github.com/fqzz2000/tiny-tictok/model"
	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}



var (
	videoIndexMap = map[string]struct{}{
		".mp4":  {},
		".avi":  {},
		".wmv":  {},
		".flv":  {},
		".mpeg": {},
		".mov":  {},
	}
	pictureIndexMap = map[string]struct{}{
		".jpeg":{},
		".jpg": {},
		".bmp": {},
		".png": {},
		".svg": {},
	}
)

func publishVideoError(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		StatusCode: 1,
		StatusMsg: msg,
	})
}

func publishVideoSuccess(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
	})
}

func GenerateCover(videoPath string, coverPath string) {
 	command := exec.Command(config.Info.FfmpegPath,
			"-i", videoPath,
			"-vf", "select=eq(n\\, 0)",
			"-q:v", "3",
			coverPath,
		)
	command.Run()
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	// Check if user exists
	userIdRaw, _ := c.Get("UserId")
	userId, _ := userIdRaw.(int64)
	videoTitle := c.Query("title")
	if videoTitle == "" {
		videoTitle = c.PostForm("title")
	}
	// get file 
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "fail to get the data",
		})
		return
	}
	// check if the file is a video
	suffix := filepath.Ext(data.Filename)
	if _, ok := videoIndexMap[suffix]; !ok {
		publishVideoError(c, "Invalid Video Format")
		return
	}

	// save file into disk
	
	filename := fmt.Sprint(userIdRaw) + "_" + fmt.Sprint(model.NewVideoDAO().CountVideoByOwnerID(userId)) + "_" + filepath.Base(data.Filename)
	saveFile := filepath.Join("./public/videos/", filename)
	err = c.SaveUploadedFile(data, saveFile)
	if err != nil {
		publishVideoError(c, "Fail to Save the file")
	}
	// generate cover for the video	save cover into disk
	covername := filename[:len(filename) - len(filepath.Ext(filename))] + ".png"
	GenerateCover(saveFile, filepath.Join("./public/covers/", covername))

	// generate video struct
	newVideo := model.VideoDB {
		VideoTitle: videoTitle,
		VideoDesc: "nothing",
		VideoOwner: (userId),
		VideoCrtTime: time.Now(),
		VideoFile: filename,
		CoverFile: covername,
	}


	// add video to the database
	model.NewVideoDAO().AddNewVideo(&newVideo)


	// return 
	c.JSON(http.StatusOK, Response {
		StatusCode: 0,
		StatusMsg:  filename + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	userIdString := c.Query("user_id")
	//TODO: need to check if user_id exists (user migh be deleted)
	userId, _ := strconv.ParseInt(userIdString, 10, 64)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response {
			StatusCode: 0,
		},
		VideoList: DecorateVideos(model.NewVideoDAO().QueryVideoByOwnerID(userId), userId),
	})
}
