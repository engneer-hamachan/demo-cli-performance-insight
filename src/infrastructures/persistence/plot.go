package persistence

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/guptarohit/asciigraph"
	"github.com/jinzhu/gorm"
	"main/src/domain/model/plotData"
	"main/src/domain/model/storeData"
	"main/src/domain/repository"
	"main/src/infrastructures/persistence/dto"
)

type plotPersistence struct {
	Conn *gorm.DB
}

func NewPlotPersistence(conn *gorm.DB) repository.PlotRepository {
	return &plotPersistence{
		Conn: conn,
	}
}

func (pp *plotPersistence) InsertStoreData(storeData *storeData.StoreData) {
	save_store_data := dto.ConvertStoreData(storeData)
	pp.Conn.Save(&save_store_data)
}

func (pp *plotPersistence) GetPlotData() *plotData.PlotData {

	store_datas := dto.StoreDatas{}

	pp.Conn.Table("store_data").Order("data").Find(&store_datas.Data)
	result_store_datas := dto.AdaptStoreDatas(&store_datas)

	data_map := map[string][]float64{}
	color_map := map[string]string{}

	for _, d := range result_store_datas {
		data_map[d.GetLabel()] = append(data_map[d.GetLabel()], d.GetData())
		color_map[d.GetLabel()] = d.GetColor()
	}

	plot_data := plotData.NewPlotData([][]float64{}, []string{})

	for k, v := range data_map {
		ct := 1
		data := []float64{0}
		start_time := float64(0)

		for _, i := range v {
			if ct%2 == 0 {
				data = append(data, i-start_time)
			} else {
				start_time = i
			}
			ct += 1
		}
		plot_data.Append(data, color_map[k])
	}

	return plot_data
}

func (pp *plotPersistence) PlotGraph(data *plotData.PlotData) {

	colors := []asciigraph.AnsiColor{}

	for _, color := range data.GetColor() {
		colors = append(colors, asciigraph.ColorNames[color])
	}

	graph :=
		asciigraph.PlotMany(
			data.GetData(),
			asciigraph.Width(70),
			asciigraph.Height(10),
			asciigraph.Precision(8),
			asciigraph.SeriesColors(
				colors...,
			),
			asciigraph.LabelColor(
				asciigraph.SandyBrown,
			),
		)

	fmt.Print("\033[H\033[2J")

	title := figure.NewFigure("Speed Plotter", "", true)
	title.Print()

	fmt.Println("--------------------------------------------------------------------------------------")
	fmt.Println(graph)
	fmt.Println("--------------------------------------------------------------------------------------")
}
