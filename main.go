package main

import (
	"main/src/infrastructures/persistence"
	"main/src/presentation/handler"
	"main/src/usecase"
	"os"
)

func main() {
	//DI
	plotRepository := persistence.NewPlotPersistence()
	httpRequestRepository := persistence.NewHttpRequestPersistence()
	cloudWatchUseCase := usecase.NewCloudWatchUseCase(plotRepository, httpRequestRepository)
	cloudWatchHandler := handler.NewCloudWatchHandler(cloudWatchUseCase)

	//Run
	cloudWatchHandler.RunCheckSpeedInsight(os.Args[1])
}
