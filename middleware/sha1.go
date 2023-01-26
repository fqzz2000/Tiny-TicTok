package middleware

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

// encode a string into SHA1
func SHA1(s string) string {
	o := sha1.New()
	o.Write([]byte(s))
	return hex.EncodeToString(o.Sum(nil))
}

// set context to sha1
func EncodeSHA1Password(c *gin.Context) {
	// dealing with GET
	pswd, ok := c.GetQuery("password")
	// dealing with POST
	if !ok {
		pswd = c.PostForm("password")
	}
	c.Set("password", SHA1(pswd))
}

func SHA1Middleware() gin.HandlerFunc {
	return EncodeSHA1Password
}