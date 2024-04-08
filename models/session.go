package models

type Session struct {
	ID        string `json:"_id"`
	SessionId string `json:"sessionId"`
	Time      string `json:"time"`
	Key       string `json:"key"`
}