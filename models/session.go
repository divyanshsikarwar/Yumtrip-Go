package models

type Session struct {
	ID        string `json:"_id"`
	SessionId string `json:"sessionId"`
	UserId    string `json:"userId"`
	Time      string `json:"time"`
	Expired  bool   `json:"expired"`
}