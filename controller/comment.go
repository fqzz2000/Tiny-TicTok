package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fqzz2000/tiny-tictok/model"
	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

func CommentAction(c *gin.Context) {
	usrIdRaw, _ := c.Get("UserId")
	usrId, _ := strconv.ParseInt(fmt.Sprint(usrIdRaw), 10, 64)
	actionType := c.Query("action_type")
	if actionType == "" {
		actionType = c.PostForm("action_type")
	}
	videoIdRaw := c.Query("video_id")
	if videoIdRaw == "" {
		videoIdRaw = c.PostForm("video_id")
	}
	videoId, _ := strconv.ParseInt(videoIdRaw, 10, 64)


	//TODO: check if user exists
	// add comment
	if actionType == "1" {
		commentText := c.Query("comment_text")
		if commentText == "" {
			commentText = c.PostForm("comment_text")
		}
		t := time.Now()
		comment := model.CommentDB{
			CommentUserID: usrId,
			CommentVideoID: videoId,
			CommentContent: commentText,
			CommentCrtTime: t,
		}
		model.NewCommentDAO().AddNewComment(&comment)
		year, month, _ := t.Date()
		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
			Comment: Comment{
				Id:         comment.CommentID,
				User:       DecorateUser(usrId),
				Content:    commentText,
				CreateDate: fmt.Sprintf("%v-%02d",year, month),
			}})
		return
	// delete comments
	} else if actionType == "2" {
		commentIdRaw := c.Query("comment_id")
		if commentIdRaw == "" {
			commentIdRaw = c.PostForm("comment_id")
		}
		commentId, _ := strconv.ParseInt(commentIdRaw, 10, 64)
		model.NewCommentDAO().DeleteComment(commentId)
		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
			Comment: Comment{
				Id:         commentId,
				User:       DecorateUser(usrId),
			}})
			return
	}
}

// CommentAction no practical effect, just check if token is valid
func CommentAction____t(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			text := c.Query("comment_text")
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment: Comment{
					Id:         1,
					User:       user,
					Content:    text,
					CreateDate: "05-01",
				}})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoIdRaw := c.Query("video_id")
	// TODO: CHEK IF THE VIDEO ID EXISTS
	videoId, _ := strconv.ParseInt(videoIdRaw, 10, 64)
	comments := DocorateComments(model.NewCommentDAO().QueryCommentsByVideoID(videoId))
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: comments,
	})
}
