package cronworkers

import (
	"context"
	"log"
	"yumtrip/core"
	"yumtrip/models"

	"github.com/robfig/cron"
)

//Create a cron job which checks the orders every 1 hour and move it to inactive state if status is pending or delivered using a cron library

var CronVar *cron.Cron

func Init() {
	c := cron.New()
	CronVar = c
	c.AddFunc("0 * * * *", UpdateOrderStatus) //Runs Every Hour
	c.AddFunc("0 0 * * *", ExpireSessions) //Runs Every Day
	c.AddFunc("0 0 * * 0", DeleteExpiredSession)//Every week
	c.Start()
}

func UpdateOrderStatus() {
	orders, err := core.GetNewInactiveOrders(context.Background())
	if err != nil {
		log.Println("Error getting orders by status", err)
		return
	}
	for _, order := range orders {
		order.Status = "inactive"
		err = order.UpdateOrder(context.Background())
		if err != nil {
			log.Println("Error updating order", err)
		}
	}
}

func ExpireSessions() {
	toBeExpiredSessions, err := core.GetNewExpiredSessions(context.Background())
	if err != nil {
		log.Println("Error getting expired sessions", err)
		return
	}
	var updatedSessions []models.Session
	for _, session := range toBeExpiredSessions {
		session.Expired = true
		updatedSessions = append(updatedSessions, session)
	}
	err = core.BulkUpdateSessions(context.Background(), updatedSessions)
	if err != nil {
		log.Println("Error updating sessions", err)
	}	
}

func DeleteExpiredSession() {
	toBeDeletedSessions, err := core.GetOldExpiredSessions(context.Background())
	if err != nil {
		log.Println("Error getting expired sessions", err)
		return
	}	
	err = core.BulkDeleteSessions(context.Background(), toBeDeletedSessions)
	if err != nil {
		log.Println("Error deleting sessions", err)
	}
}