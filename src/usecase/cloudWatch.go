package usecase

import (
	"fmt"
	"main/src/domain/model/plotData"
	"main/src/domain/model/storeData"
	"main/src/domain/repository"
)

type CloudWatchUseCase interface {
	StoreData(label string, data float64, color string) error
}

type cloudWatchUseCase struct {
	plotRepository repository.PlotRepository
}

func NewCloudWatchUseCase(pr repository.PlotRepository) CloudWatchUseCase {
	return &cloudWatchUseCase{
		plotRepository: pr,
	}
}

func (cwu *cloudWatchUseCase) StoreData(label string, data float64, color string) error {
	store_data := storeData.NewStoreData(
		label,
		data,
		color,
	)

	cwu.plotRepository.InsertStoreData(store_data)
	store_datas := cwu.plotRepository.GetStoreData()
	fmt.Println(store_datas)
	m := map[string][]float64{}

	for _, d := range store_datas {
		m[d.GetLabel()] = append(m[d.GetLabel()], d.GetData())
	}

	fmt.Println(m)

	plotData := plotData.NewPlotData([][]float64{})

	for _, v := range m {
		plotData.Append(v)
	}

	cwu.plotRepository.PlotSpeedInsight(plotData)

	return nil
}
