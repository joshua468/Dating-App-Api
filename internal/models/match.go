package models

type Match struct {
	UserID        uint `json:"user_id"`
	MatchedUserID uint `json:"matched_user_id"`
}
