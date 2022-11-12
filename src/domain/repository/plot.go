package repository

import (
	"main/src/domain/model/plotData"
)

type PlotRepository interface {
	PlotSpeedInsight(plotData *plotData.PlotData)
}
