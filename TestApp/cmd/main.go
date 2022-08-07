package main

import (
	"TestApp/internal/app"
)

func main() {
	/*
		r := gin.Default()
		r.GET("/ping", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run(":9090")
	*/
	//godotenv.Load(".env")
	//connection := os.Getenv("SQL_DRIVER")
	//fmt.Println("driver: " + connection)
	//server.RunServer()
	app.RunServer()

}
