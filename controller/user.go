package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/fqzz2000/tiny-tictok/middleware"
	"github.com/fqzz2000/tiny-tictok/model"
	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}


func Register(c *gin.Context) {
	userName := c.Query("username")
	rawVal, _ := c.Get("password")
	password, ok := rawVal.(string)
	if !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Fail to parse password"},
		})
	}
	// check if user name exist
	if model.NewUserDAO().QueryNameExists(userName) {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
	// write info into database
	newUser := model.UserDB {
		UserName: userName,
		UserPswd: password,
	}
	model.NewUserDAO().AddNewUser(&newUser)
	// construct token
	token, err := middleware.ReleaseToken(newUser.UserID)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Fail to Generate Token"},
		})
	}
	// construct response
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   newUser.UserID,
		Token:    token,
	})
}
}

func Register__T(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println(username)
	token := username + password

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		newUser := User{
			Id:   userIdSequence,
			Name: username,
		}
		usersLoginInfo[token] = newUser
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userIdSequence,
			Token:    username + password,
		})
	}
}


func Login(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		username = c.PostForm("username")
	}
	rawPswd, _ := c.Get("password")
	password, ok := rawPswd.(string)
	if !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{
				StatusCode : 1,
				StatusMsg: "Cannot Parse the Password",
			},
		})
		return
	}

	if !model.NewUserDAO().QueryNameExists(username) {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: "User Does Not Exists",
			},
		})
		return 
	}
	
	userInfo := model.NewUserDAO().QueryUserByName(username)
	if userInfo.UserPswd != password {
		c.JSON(http.StatusOK, Response{
			StatusCode: 40,
			StatusMsg: "Wrong Password",
		})
	} else {
		token, _ := middleware.ReleaseToken(userInfo.UserID)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserId: int64(userInfo.UserID),
			Token: token,
		})
	}


}


func Login___t(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token := username + password

	if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   user.Id,
			Token:    token,
		})
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}


func UserInfo (c *gin.Context) {
	tokenUserId, _ := c.Get("UserId")
	userId := c.Query("user_id")

	if userId != fmt.Sprint(tokenUserId) {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: "Inconsistent UserId",
			},
		})
		return
	}
	id, err := strconv.ParseInt(fmt.Sprint(tokenUserId), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg: "Invalid ID",
			},
		})
		return
	}
	userInfo := model.NewUserDAO().QueryUserById(id)
	usr := User{
		Id : id,
		Name: userInfo.UserName,
		FollowCount: 114514,
		FollowerCount: 1919810,
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{
			StatusCode: 0,
		},
		User : usr, 
	})

}

