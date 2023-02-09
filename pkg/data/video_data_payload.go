package data

type VideoDataPayload struct {
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description" binding:"required"`
	Tags        []string   `json:"tags" binding:"required"`
	Visibility  Visibility `json:"visibility" binding:"required"`
}
