package repository

import (
	"main/src/domain/model/plotData"
	"main/src/domain/model/storeData"
)

type PlotRepository interface {
	InsertStoreData(storeData *storeData.StoreData)
	GetPlotData() *plotData.PlotData
	PlotGraph(plotData *plotData.PlotData)
}
