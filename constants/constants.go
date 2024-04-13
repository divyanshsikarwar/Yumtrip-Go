package constants

import (
	"yumtrip/models"
	"github.com/joho/godotenv"
	"os"
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
	PasswordSalt = "yumtrip" //Initial value
)

func init() {
	err := godotenv.Load(".env")
	if err != nil{
		panic(err)
	}
	
	salt := os.Getenv("PasswordSalt")
	PasswordSalt = salt
}