package user_segmentation

type Segment struct {
	Id   int    `json:"id"`
	Slug string `json:"slug" binding:"required"`
}
