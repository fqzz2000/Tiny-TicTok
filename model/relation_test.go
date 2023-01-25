package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRelationDAO_CountRelationsByID(t *testing.T) {
	Init_DB(true)
	var zero int64 
	zero = 0
	var one int64
	one = 1
	assert.Equal(t, zero, NewRelationDAO().CountRelationsByFansID(1))
	assert.Equal(t, one, NewRelationDAO().CountRelationsByFollowerID(1))
	assert.Equal(t, one, NewRelationDAO().CountRelationsByFansID(2))
	assert.Equal(t, zero, NewRelationDAO().CountRelationsByFollowerID(2),)
}

func TestRelationDAO_AddDelete(t *testing.T) {
	Init_DB(true)
	var rls []RelationDB
	DB.Find(&rls)
	old_len := len(rls)
	tme := time.Now()
	NewRelationDAO().AddNewRelation(&RelationDB{
		FollowerID: 1,
		FansID: 1,
		FollowDate: tme,
	})
	DB.Find(&rls)
	new_len := len(rls)
	assert.Equal(t, new_len - old_len, 1)
	NewRelationDAO().DeleteRelation(1, 1)
	DB.Find(&rls)
	new_len = len(rls)
	assert.Equal(t, new_len - old_len, 0)
}