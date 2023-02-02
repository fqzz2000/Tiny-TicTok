package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

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
// 
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