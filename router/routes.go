package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/imperiustx/go_excercises/controllers"
	"github.com/imperiustx/go_excercises/module/address/addressmodel"
	"github.com/imperiustx/go_excercises/module/paymentmethod/paymentmethodmodel"
	"github.com/imperiustx/go_excercises/module/restaurant/restaurantmodel"
	"github.com/imperiustx/go_excercises/module/user/usermodel"
	"github.com/imperiustx/go_excercises/module/useraddress/useraddressmodel"
	"github.com/imperiustx/go_excercises/module/userpaymentmethod/userpaymentmethodmodel"
)

// StartDatabase start
func StartDatabase(r *gin.Engine) *gin.Engine {

	db, err := connectToDatabase()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&usermodel.User{})
	db.AutoMigrate(&useraddressmodel.UserAddress{})
	db.AutoMigrate(&userpaymentmethodmodel.UserPaymentMethod{})
	db.AutoMigrate(&addressmodel.Address{})
	db.AutoMigrate(&paymentmethodmodel.PaymentMethod{})
	db.AutoMigrate(&restaurantmodel.Restaurant{})

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:usr-id", controllers.GetUser)
	r.PUT("/users/:usr-id", controllers.UpdateUser)
	r.DELETE("/users/:usr-id", controllers.DeleteUser)
	r.POST("/users/address", controllers.CreateUserAddress)
	r.POST("/users/payment-method", controllers.CreateUserPaymentMethod)

	r.POST("/addresses", controllers.CreateAddress)
	r.GET("/addresses", controllers.GetAllAddresses)
	r.GET("/addresses/:add-id", controllers.GetAddress)
	r.PUT("/addresses/:add-id", controllers.UpdateAddress)

	r.POST("/payment-methods", controllers.CreatePaymentMethod)
	r.GET("/payment-methods", controllers.GetAllPayments)
	r.GET("/payment-methods/:pm-id", controllers.GetPayment)
	r.PUT("/payment-methods/:pm-id", controllers.UpdatePayment)

	r.POST("/restaurants", controllers.CreateRestaurant)
	r.GET("/restaurants", controllers.GetAllRestaurants)
	r.GET("/restaurants/:res-id", controllers.GetRestaurant)
	r.PUT("/restaurants/:res-id", controllers.UpdateRestaurant)
	r.DELETE("/restaurants/:res-id", controllers.DeleteRestaurant)

	return r
}