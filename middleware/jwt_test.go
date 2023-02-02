package middleware

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestReleaseParseToken(t *testing.T) {
	var id int64 = 10086
	claims, _ := ReleaseToken(id)
	restoredClaims, ok := ParseToken(claims)
	assert.Equal(t, true, ok) 
	assert.Equal(t, id, restoredClaims.UserId)
}

func TestParseError(t *testing.T) {
	restoredClaims, ok := ParseToken("123")
	assert.Equal(t, false, ok)
	assert.Equal(t, nil == restoredClaims, true)
}

func TestJwtMiddleWare(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	var id int64 = 10086
	tokenString, _ := ReleaseToken(id)
	c.Params = []gin.Param{{Key: "token", Value: tokenString}}
	jwtGetUserIDFromToken(c)
	assert.Equal(t, id , c.Query("UserId"))
}