package model

import (
	"time"
)
type RelationDB struct {
	FollowerID 	int64 `gorm:"primaryKey"`
	FansID 		int64 `gorm:"primaryKey"`
	FollowDate 	time.Time
}

func (*RelationDB) TableName() string {
	return "relations"
}

type RelationDAO struct {}

func NewRelationDAO() RelationDAO {
	return RelationDAO{}
}

func (RelationDAO) CountRelationsByFollowerID(id int64) int64 {
	var count int64
	DB.Model(&RelationDB{}).Where("follower_id = ?", id).Count(&count)
	return count
}

func (RelationDAO) CountRelationsByFansID(id int64) int64 {
	var count int64
	DB.Model(&RelationDB{}).Where("fans_id = ?", id).Count(&count)
	return count
}

func (RelationDAO) AddNewRelation(newRelation *RelationDB) {
	DB.Create(newRelation)
}

func (RelationDAO) DeleteRelation(follower_id int64, fans_id int64) {
	DB.Where("follower_id = ? AND fans_id = ?", follower_id, fans_id).Delete(&RelationDB{})
}