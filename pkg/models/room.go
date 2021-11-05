package models


type Room struct {
	UserID int `json:"-"`
	RoomTitle string `json:"room_title"`
	RoomState string `json:"room_state"`
	IsComplete bool `json:"is_complete"`
	Score int `json:"score"`
}