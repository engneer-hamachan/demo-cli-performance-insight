package persistence

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/guptarohit/asciigraph"
	"main/src/domain/model/plotData"
	"main/src/domain/repository"
)

type plotPersistence struct{}

func NewPlotPersistence() repository.PlotRepository {
	return &plotPersistence{}
}

func (pp *plotPersistence) PlotSpeedInsight(data *plotData.PlotData) {

	fmt.Print("\033[H\033[2J")
	graph := asciigraph.Plot(data.GetData(), asciigraph.Height(10), asciigraph.Precision(6))

	title := figure.NewFigure("Cloud Watch CLI", "", true)
	title.Print()

	fmt.Println("------------------------------------------------------------------------------------------")
	fmt.Println(graph)
	fmt.Println("------------------------------------------------------------------------------------------")
	fmt.Println("page speed watch now")
}
