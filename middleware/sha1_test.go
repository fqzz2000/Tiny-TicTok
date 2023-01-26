package middleware

import (
	"fmt"
	"testing"
)

func TestSHA1function(t *testing.T) {
	s := "happy birthday to you"
	fmt.Println(SHA1(s))
}

