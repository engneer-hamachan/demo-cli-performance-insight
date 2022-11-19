package usecase

import (
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

	plotData, plotColor := cwu.plotRepository.GetStoreData()

	cwu.plotRepository.PlotSpeedInsight(plotData, plotColor)

	return nil
}
