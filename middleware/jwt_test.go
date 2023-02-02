package middleware

import (
	"testing"

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