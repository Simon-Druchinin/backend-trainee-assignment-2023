package user_segmentation

import "time"

type User struct {
	Id int `json:"id"`
}

type UserSegment struct {
	User_id int    `json:"user_id"`
	Slug    string `json:"segment_slug"`
}

type UserSegmentHistory struct {
	User_id        int       `json:"user_id"`
	Slug           string    `json:"slug"`
	Operation_type string    `json:"operation_type"`
	Timestamp      time.Time `json:"timestamp"`
}
