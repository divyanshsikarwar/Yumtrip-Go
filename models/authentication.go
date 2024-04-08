package models

import "time"

type Authentication struct {
	ID    string `json:"_id"`
	Email string `json:"email"`
	Code  string `json:"code"`
	Time  time.Time `json:"time"`
}