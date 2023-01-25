package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCommentDAO_CountCommentsByID(t *testing.T) {
	Init_DB(true)
	var zero int64 
	zero = 0
	var one int64
	one = 1
	assert.Equal(t, one, NewCommentDAO().CountCommentsByVideoID(19))
	assert.Equal(t, zero, NewCommentDAO().CountCommentsByVideoID(1))
}

func TestCommentDAO_AddDelete(t *testing.T) {
	Init_DB(true)
	var rls []CommentDB
	DB.Find(&rls)
	old_len := len(rls)
	addComment := CommentDB{
		CommentUserID: 1,
		CommentVideoID: 19,
		CommentContent: "testContent",
		CommentCrtTime: time.Now(),
	}
	NewCommentDAO().AddNewComment(&addComment)
	DB.Find(&rls)
	new_len := len(rls)
	assert.Equal(t, new_len - old_len, 1)
	NewCommentDAO().DeleteComment(addComment.CommentID)
	DB.Find(&rls)
	new_len = len(rls)
	assert.Equal(t, new_len - old_len, 0)
}