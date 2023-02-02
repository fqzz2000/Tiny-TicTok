package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

var jwtKey = []byte("WLB")

type Claims struct {
	UserId int64
	jwt.RegisteredClaims
} 

// release a token given user Id
func ReleaseToken(userId int64) (string, error) {
	issueTime := time.Now()
	expirationTime := issueTime.Add(7 * 24 * time.Hour)
	// construct claim
	claims := &Claims{
		UserId : userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt: jwt.NewNumericDate(issueTime),
			NotBefore: jwt.NewNumericDate(issueTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// convert token into string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
// reconstruct Claims based on the token string
// only check if the token is valid, issue date and expired date will not be checked
func ParseToken(tokenString string) (*Claims, bool) {
	token, _:= jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if token != nil {
		if key, ok := token.Claims.(*Claims); ok {
			if token.Valid {
				return key, true
			} else {
				return key, false 
			}
		}
	}
	return nil, false

}

func jwtGetUserIDFromToken(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		tokenString = c.PostForm("token")
	}
	claims, ok := ParseToken(tokenString)
	if !ok {
		c.JSON(http.StatusOK, Response{
			StatusCode: 401, 
			StatusMsg: "Token Not Exists",
		})
		c.Abort()
		return
	}
	// check token time
	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		c.JSON(http.StatusOK, Response {
			StatusCode: 402,
			StatusMsg: "Token Expired",
		})
		c.Abort()
		return
	}
	c.Set("UserId", claims.UserId)
	c.Next()
}

// parse the token from the user, 
func JwtMiddleware() gin.HandlerFunc {
	return jwtGetUserIDFromToken
}