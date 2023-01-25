package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLikeDAO_CountLikeByID(t *testing.T) {
	Init_DB(true)
	var zero int64 
	zero = 0
	var one int64
	one = 1
	assert.Equal(t, zero, NewLikeDAO().CountLikesByUserID(2))
	assert.Equal(t, one, NewLikeDAO().CountLikesByVideoID(19))
	assert.Equal(t, one, NewLikeDAO().CountLikesByUserID(1))
	assert.Equal(t, zero, NewLikeDAO().CountLikesByVideoID(20),)
}

func TestLikeDAO_AddDelete(t *testing.T) {
	Init_DB(true)
	var rls []LikeDB
	DB.Find(&rls)
	old_len := len(rls)
	NewLikeDAO().AddNewLike(&LikeDB{
		UserID: 2,
		VideoID: 19,
		LikeTime: time.Now(),
	})
	DB.Find(&rls)
	new_len := len(rls)
	assert.Equal(t,  1, new_len - old_len)
	NewLikeDAO().DeleteLike(2, 19)
	DB.Find(&rls)
	new_len = len(rls)
	assert.Equal(t,  0, new_len - old_len,)
}