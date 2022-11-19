package repository

import (
	"main/src/domain/model/plotData"
	"main/src/domain/model/storeData"
)

type PlotRepository interface {
	InsertStoreData(storeData *storeData.StoreData)
	GetStoreData() []storeData.StoreData
	PlotSpeedInsight(plotData *plotData.PlotData)
}
