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

func (pp *plotPersistence) GetStoreData() []storeData.StoreData {
	var store_datas dto.StoreDatas
	pp.Conn.Table("store_data").Find(&store_datas.Data)
	result_store_datas := dto.AdaptStoreDatas(&store_datas)

	return result_store_datas
}

func (pp *plotPersistence) PlotSpeedInsight(data *plotData.PlotData) {

	fmt.Print("\033[H\033[2J")
	graph := asciigraph.PlotMany(data.GetData(), asciigraph.Height(10), asciigraph.Precision(6))

	title := figure.NewFigure("Cloud Watch CLI", "", true)
	title.Print()

	fmt.Println("------------------------------------------------------------------------------------------")
	fmt.Println(graph)
	fmt.Println("------------------------------------------------------------------------------------------")
	fmt.Println("page speed watch now")
}
