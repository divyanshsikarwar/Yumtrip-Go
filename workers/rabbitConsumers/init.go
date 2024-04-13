package rabbitconsumers

import (
	"log"
	"yumtrip/constants"
)

func Init() {
	totalEmailConsumers := constants.TotalEmailConsumers
	for i := 0; i < totalEmailConsumers; i++ {
		go EmailConsumer()
	}
	totalNptificationConsumers := constants.TotalNotificationConsumers
	for i := 0; i < totalNptificationConsumers; i++ {
		go NotificationConsumer()
	}

}
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
