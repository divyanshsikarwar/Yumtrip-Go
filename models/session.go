package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Session is the model for user sessions
type Session struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	SessionID primitive.ObjectID `json:"sessionId" bson:"sessionId"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId"`
	Time      time.Time          `json:"time" bson:"time"`
	Expired   bool               `json:"expired" bson:"expired"`
}

func (s *Session) New() *Session {
	return &Session{}
}

// GetSessionByID retrieves a session by its ID
func GetSessionByID(ctx context.Context, id primitive.ObjectID) (Session, error) {
	var session Session
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&session)
	if err != nil {
		return Session{}, err
	}
	return session, nil
}

// GetSessionsByUserID retrieves sessions by user ID
func GetSessionByUserID(ctx context.Context, userID primitive.ObjectID) (Session, error) {
	var session Session
	err := collection.FindOne(ctx, bson.M{"userID": userID}).Decode(&session)
	if err != nil {
		return Session{}, err
	}
	return session, nil
}

// DeleteSessionByID deletes a session by its ID
func DeleteSessionByID(ctx context.Context, id primitive.ObjectID) error {
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

// UpsertSession inserts or updates a session
func UpsertSession(ctx context.Context, session Session) error {
	opts := options.Replace().SetUpsert(true)
	_, err := collection.ReplaceOne(ctx, bson.M{"_id": session.ID}, session, opts)
	if err != nil {
		return err
	}
	return nil
}

// GetSessionsByQuery retrieves sessions based on query
func GetSessionsByQuery(ctx context.Context, query bson.M) ([]Session, error) {
	cursor, err := collection.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var sessions []Session
	for cursor.Next(ctx) {
		var session Session
		if err := cursor.Decode(&session); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return sessions, nil
}

// BulkUpdateSessions updates sessions in bulk based on query
func BulkUpdateSessions(ctx context.Context, query []bson.M) error {
	// Your implementation for bulk update
	return nil
}
