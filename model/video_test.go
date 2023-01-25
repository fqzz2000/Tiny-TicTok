package model

import (
	"fmt"
	"testing"
	"time"
)


func TestVideoDAO_AddVideo(t *testing.T) {
	Init_DB(true)
	NewVideo := VideoDB{
		VideoTitle: "testVideo",
		VideoDesc: "",
		VideoOwner: 1,
		VideoCrtTime: time.Now(),
		VideoFile: "",
		CoverFile: "",
	}
	DB.Where("video_title = ?", "testVideo").Delete(&VideoDB{})
	NewVideoDAO().AddNewVideo(&NewVideo)
	fmt.Printf("Video id = %d\n", NewVideo.VideoID);
	var testV VideoDB
	DB.Where("video_title = ?", "testVideo").Find(&testV)
	fmt.Printf("%+v", testV);
	DB.Where("video_title = ?", "testVideo").Delete(&VideoDB{})
}

func TestVideoDAO_QueryVideoBeforeTime(t *testing.T) {
	Init_DB(true)
	newVideo := VideoDB{
		VideoTitle: "testQuery",
		VideoOwner: 1,
		VideoCrtTime: time.Now(),
	}
	NewVideoDAO().AddNewVideo(&newVideo)
	NewVideoDAO().AddNewVideo(&newVideo)
	NewVideoDAO().AddNewVideo(&VideoDB{
		VideoTitle: "testQuery",
		VideoOwner: 1,
		VideoCrtTime: time.Date(2010, time.April, 3, 2, 1, 2, 3, time.UTC),
	})
	testTime := time.Date(2016, time.August, 18, 23, 15, 8, 4, time.UTC)
	s := NewVideoDAO().QueryVideoBeforeTime(testTime, 30)
	fmt.Printf("%+v\n", s);
	DB.Where("video_title = ?", "testQuery").Delete(&VideoDB{})

}