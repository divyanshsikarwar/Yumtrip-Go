package cronworkers

import (
	"log"
	"yumtrip/core"

	"github.com/robfig/cron"
)

//Create a cron job which checks the orders every 1 hour and move it to inactive state if status is pending or delivered using a cron library

var CronVar *cron.Cron

func Init() {
	c := cron.New()
	CronVar = c
	c.AddFunc("0 * * * *", UpdateOrderStatus) //Runs Every Hour
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