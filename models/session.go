package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//To be depricated, moved to Oauth2.0

type Session struct {
	ID        primitive.ObjectID `json:"_id"`
	SessionId primitive.ObjectID `json:"sessionId"`
	UserId    primitive.ObjectID `json:"userId"`
	Time      time.Time `json:"time"`
	Expired   bool   `json:"expired"`
}

func (s *Session) New() *Session {
	return &Session{}
}

func GetSessionDocById(id primitive.ObjectID) (Session, error) {
	//query := bson.M{"_id": id}
	//session, err := GetSessionByQuery(query)
	return Session{}, nil
}

func GetSessionByUserId(userId primitive.ObjectID) (Session, error) {
	//query := bson.M{"userId": userId}
	//session, err := GetSessionByQuery(query)
	return Session{}, nil
}

func DeleteSessionById(id primitive.ObjectID) error {
	//_, err := DeleteSessionByQuery(bson.M{"_id":
	return nil
}

func UpsertSession(session Session) error {
	//_, err := UpsertSessionByQuery(bson.M{"_id": session.ID}, session)
	return nil
}

func GetSessionsByQuery(query bson.M) ([]Session, error) {
	//sessions, err := GetSessionByQuery(query)
	return []Session{}, nil
}

func BulkUpdateSessions(query []bson.M) error {
	//_, err := BulkUpdateSessionsByQuery(query)
	return nil
}