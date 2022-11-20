package repository

import (
	"github.com/guptarohit/asciigraph"
	"main/src/domain/model/plotData"
	"main/src/domain/model/storeData"
)

type PlotRepository interface {
	InsertStoreData(storeData *storeData.StoreData)
	GetPlotData() (*plotData.PlotData, []asciigraph.AnsiColor)
	PlotGraph(plotData *plotData.PlotData, colors []asciigraph.AnsiColor)
}
