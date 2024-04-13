package constants

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

const ( //Permission could be any of the following : View_Analytics, Manage_Orders, Manage_Menu, Manage_Coupons, Manage_Users
	ViewAnalytics  = "View_Analytics"
	ManageOrders string = "Manage_Orders"
	ManageMenu string = "Manage_Menu"
	ManageCoupons string = "Manage_Coupons"
	ManageUsers string = "Manage_Users"
)
var (
	AllPermissions []string = []string{ViewAnalytics, ManageOrders, ManageMenu, ManageCoupons, ManageUsers}
)
var (
	RabbitMQURL string
)
var (
	MongoURI string
	MongoClient *mongo.Client
	MongoCollections = map[string]*mongo.Collection{}
	MongoDBName = "yumtrip"
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
	RabbitMQURL = os.Getenv("RabbitMQURL")
	TotalEmailConsumers,err = strconv.Atoi(os.Getenv("TotalEmailConsumers"))
	if err != nil {
		log.Println("Could not parse TotalEmailConsumers, using default value")
	}
	TotalNotificationConsumers,err = strconv.Atoi(os.Getenv("TotalNotificationConsumers"))
	if err != nil {
		log.Println("Could not parse TotalNotificationConsumers, using default value")
	}
	MongoURI = os.Getenv("MongoURI")
	if len(MongoURI) == 0 {
		log.Panic("Could not connect to MongoDB, MongoURI not provided")
	}

}