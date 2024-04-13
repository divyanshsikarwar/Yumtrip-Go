package cronworkers

import (
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
	orders, err := core.GetNewInactiveOrders()
	if err != nil {
		log.Println("Error getting orders by status", err)
		return
	}
	for _, order := range orders {
		order.Status = "inactive"
		err = order.UpdateOrder()
		if err != nil {
			log.Println("Error updating order", err)
		}
	}
}

func ExpireSessions() {
	toBeExpiredSessions, err := core.GetNewExpiredSessions()
	if err != nil {
		log.Println("Error getting expired sessions", err)
		return
	}
	var updatedSessions []models.Session
	for _, session := range toBeExpiredSessions {
		session.Expired = true
		updatedSessions = append(updatedSessions, session)
	}
	err = core.BulkUpdateSessions(updatedSessions)
	if err != nil {
		log.Println("Error updating sessions", err)
	}	
}

func DeleteExpiredSession() {
	toBeDeletedSessions, err := core.GetOldExpiredSessions()
	if err != nil {
		log.Println("Error getting expired sessions", err)
		return
	}	
	err = core.BulkDeleteSessions(toBeDeletedSessions)
	if err != nil {
		log.Println("Error deleting sessions", err)
	}
}