package moodlogscontroller

type moodlogsCreateRequest struct {
	Mood   int    `json:"mood" binding:"required"`
	Note   string `json:"note"`
	Causes string `json:"causes" binding:"required"`
}
