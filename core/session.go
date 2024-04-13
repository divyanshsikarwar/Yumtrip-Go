package core

import (
	"context"
	"time"
	"yumtrip/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//To be depricated, moved to Oauth2.0

func ValidateSession(context context.Context, sessionId primitive.ObjectID) (bool,primitive.ObjectID, error) {
	session, err := models.GetSessionByID(context, sessionId)
	if err != nil {
		return false, primitive.NilObjectID, err
	}
	return session.Expired, session.UserID, nil
}

func CreateUpdateSession(context context.Context, sessionId,UserId primitive.ObjectID) error {
	newSession := models.Session{
		ID: sessionId,
		Expired: false,
		UserID: UserId,
		Time: time.Now(),
	}
	return models.UpsertSession(context, newSession)
}

func GetSessionIdByUserId(context context.Context, userId primitive.ObjectID) primitive.ObjectID {
	session, err := models.GetSessionByUserID(context, userId)
	if err != nil {
		return primitive.NewObjectID()
	}
	return session.ID
}
func InValidateSession(context context.Context, sessionId primitive.ObjectID) error {
	return models.DeleteSessionByID(context, sessionId)
}

func GetNewExpiredSessions(context context.Context, ) ([]models.Session, error) {
	query := bson.M{
		"expired": false,
		"time": bson.M{
			"$lt": time.Now().Add(-time.Hour * 24),
		},
	}
	return models.GetSessionsByQuery(context, query)	
}

func GetOldExpiredSessions(context context.Context, ) ([]models.Session, error) {
	query := bson.M{
		"expired": true,
		"time": bson.M{
			"$lt": time.Now().Add(-time.Hour * 168),
		},
	}
	return models.GetSessionsByQuery(context, query)
}

func BulkUpdateSessions(context context.Context, sessions []models.Session) error {
	var bulkUpdateQuery  []bson.M
	for _, session := range sessions {
		bulkUpdateQuery = append(bulkUpdateQuery, bson.M{
			"updateOne": bson.M{
				"filter": bson.M{"_id": session.ID},
				"update": bson.M{
					"$set": bson.M{"expired": session.Expired, "time": session.Time}},
				},
		})
	}
	return models.BulkUpdateSessions(context, bulkUpdateQuery)

}

func BulkDeleteSessions(context context.Context, sessions []models.Session) error {
	var bulkDeleteQuery  []bson.M
	for _, session := range sessions {
		bulkDeleteQuery = append(bulkDeleteQuery, bson.M{
			"deleteOne": bson.M{
				"filter": bson.M{"_id": session.ID},
				},
		})
	}
	return models.BulkUpdateSessions(context, bulkDeleteQuery)
}