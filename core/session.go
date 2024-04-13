package core

import (
	"time"
	"yumtrip/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//To be depricated, moved to Oauth2.0

func ValidateSession(sessionId primitive.ObjectID) (bool,primitive.ObjectID, error) {
	session, err := models.GetSessionDocById(sessionId)
	if err != nil {
		return false, primitive.NilObjectID, err
	}
	return session.Expired, session.UserId, nil
}

func CreateUpdateSession(sessionId,UserId primitive.ObjectID) error {
	newSession := models.Session{
		ID: sessionId,
		Expired: false,
		UserId: UserId,
		Time: time.Now(),
	}
	return models.UpsertSession(newSession)
}

func GetSessionIdByUserId(userId primitive.ObjectID) primitive.ObjectID {
	session, err := models.GetSessionByUserId(userId)
	if err != nil {
		return primitive.NewObjectID()
	}
	return session.ID

}
func InValidateSession(sessionId primitive.ObjectID) error {
	return models.DeleteSessionById(sessionId)
}

func GetNewExpiredSessions() ([]models.Session, error) {
	query := bson.M{
		"expired": false,
		"time": bson.M{
			"$lt": time.Now().Add(-time.Hour * 24),
		},
	}
	return models.GetSessionsByQuery(query)	
}

func GetOldExpiredSessions() ([]models.Session, error) {
	query := bson.M{
		"expired": true,
		"time": bson.M{
			"$lt": time.Now().Add(-time.Hour * 168),
		},
	}
	return models.GetSessionsByQuery(query)
}

func BulkUpdateSessions(sessions []models.Session) error {
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
	return models.BulkUpdateSessions(bulkUpdateQuery)

}

func BulkDeleteSessions(sessions []models.Session) error {
	var bulkDeleteQuery  []bson.M
	for _, session := range sessions {
		bulkDeleteQuery = append(bulkDeleteQuery, bson.M{
			"deleteOne": bson.M{
				"filter": bson.M{"_id": session.ID},
				},
		})
	}
	return models.BulkUpdateSessions(bulkDeleteQuery)
}