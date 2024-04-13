package constants

import (
	"log"
	"os"
	"strconv"
	"yumtrip/models"

	"github.com/joho/godotenv"
)

const ( //Permission could be any of the following : View_Analytics, Manage_Orders, Manage_Menu, Manage_Coupons, Manage_Users
	ViewAnalytics models.Permission = "View_Analytics"
	ManageOrders models.Permission = "Manage_Orders"
	ManageMenu models.Permission = "Manage_Menu"
	ManageCoupons models.Permission = "Manage_Coupons"
	ManageUsers models.Permission = "Manage_Users"
)

var (
	AllPermissions []models.Permission = []models.Permission{ViewAnalytics, ManageOrders, ManageMenu, ManageCoupons, ManageUsers}
)
var (
	PasswordSalt string
	RabbitMQURL string
)

var (
	RabbitEmailQueue = "email_queue"
	RabbitSMSQueue = "sms_queue"
	RabbitNotificationQueue = "notification_queue"
	TotalEmailConsumers = 20
	TotalNotificationConsumers = 1
)

func init() {
	err := godotenv.Load(".env")
	if err != nil{
		panic(err)
	}
	
	salt := os.Getenv("DefaultPasswordSalt")
	PasswordSalt = salt
	RabbitMQURL = os.Getenv("RabbitMQURL")
	TotalEmailConsumers,err = strconv.Atoi(os.Getenv("TotalEmailConsumers"))
	if err != nil {
		log.Println("Could not parse TotalEmailConsumers, using default value")
	}
	TotalNotificationConsumers,err = strconv.Atoi(os.Getenv("TotalNotificationConsumers"))
	if err != nil {
		log.Println("Could not parse TotalNotificationConsumers, using default value")
	}
	
}