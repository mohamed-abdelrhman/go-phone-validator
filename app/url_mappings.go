package app

import "github.com/mohamed-abdelrhman/go-phone-validator/controllers"

func mapUrls()  {
	router.GET("/ping", controllers.Ping)

	router.GET("/customers/:customer_id", controllers.Get)
	router.GET("/customers", controllers.CustomerControllers.GetAll)
	router.POST("/customers/filter", controllers.CustomerControllers.Filter)
	router.POST("/customers", controllers.CustomerControllers.Create)
	router.PUT("/customers/:customer_id", controllers.CustomerControllers.Update)
	router.PATCH("/customers/:customer_id", controllers.CustomerControllers.Update)
	router.DELETE("/customers/:customer_id", controllers.CustomerControllers.Delete)

	router.GET("/countries/:country_id", controllers.CountryControllers.Get)
	router.GET("/countries", controllers.CountryControllers.GetAll)
	router.POST("/countries", controllers.CountryControllers.Create)
	router.PUT("/countries/:country_id", controllers.CountryControllers.Update)
	router.PATCH("/countries/:country_id", controllers.CountryControllers.Update)
	router.DELETE("/countries/:country_id", controllers.CountryControllers.Delete)
}