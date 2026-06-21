package moodlogscontroller

type MoodlogsCreateRequest struct {
	Mood   int    `json:"mood" binding:"required,min=1,max=5"`
	Note   string `json:"note"`
	Causes string `json:"causes" binding:"required"`
}
