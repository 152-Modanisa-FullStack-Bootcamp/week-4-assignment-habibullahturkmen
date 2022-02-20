package model

type Video struct {
	ID                 int    `json:"id"`
	VideoAddress       string `json:"videoAddress"`
	CoverImage         string `json:"coverImage"`
	HoverImage         string `json:"hoverImage"`
	Title              string `json:"title"`
	ViewCount          int    `json:"viewCount"`
	PublishDateInMonth int    `json:"publishDateInMonth"`
	OwnerImage         string `json:"ownerImage"`
	OwnerName          string `json:"ownerName"`
	Description        string `json:"description"`
	Favorite           bool   `json:"favorite"`
}

type Videos []Video
