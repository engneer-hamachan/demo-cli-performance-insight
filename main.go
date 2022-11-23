package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main/config"
	"main/src/infrastructures/persistence"
	"main/src/presentation/handler"
	"main/src/usecase"
	"time"
)

func main() {

	//Gin Default Setting
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//Gin Cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	//Db Connection
	db := config.Connect()
	defer db.Close()

	//DI
	plotRepository := persistence.NewPlotPersistence(db)
	plotterUseCase := usecase.NewPlotterUseCase(plotRepository)
	plotterHandler := handler.NewPlotterHandler(plotterUseCase)

	//Routing
	router.GET("/plotter/:label/:data/:color", plotterHandler.RecieveData)

	fmt.Println("waiting for data...")
	fmt.Println("please request => http://0.0.0.0:1988/plotter/:label/:data/:color")

	//Run
	router.Run("0.0.0.0:1988")
}
