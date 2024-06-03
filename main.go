package main

import (
	"fmt"

	"insurance/routes"

	"github.com/gin-gonic/gin"

	db "insurance/db/postgresql"

	redisDb "insurance/db/redis"
)

func main() {
	router := gin.Default()


	// Connect to PostgreSQL database
	db, errConnectDb := db.ConnectPostgreSQL()
	if errConnectDb != nil {
			fmt.Println("Error connecting to PostgreSQL database:", errConnectDb)
			return
	}
	defer db.Close()

	// Database connection is now available in the db variable
	fmt.Println("Connected to PostgreSQL database successfully")



	//Connect Redis
	rdb := redisDb.ConnectRedis()
	defer rdb.Close()


	//Setup router
	routesInsurance := routes.NewRouterInsurance(router, db, rdb)
	routesInsurance.Setup()


	// Run the server on port 6302
	router.Run(":6302")

}