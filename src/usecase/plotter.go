package usecase

import (
	"main/src/domain/model/storeData"
	"main/src/domain/repository"
)

type PlotterUseCase interface {
	RecieveData(label string, data float64, color string)
}

type plotterUseCase struct {
	plotRepository repository.PlotRepository
}

func NewPlotterUseCase(pr repository.PlotRepository) PlotterUseCase {
	return &plotterUseCase{
		plotRepository: pr,
	}
}

func (pu *plotterUseCase) RecieveData(label string, data float64, color string) {

	store_data := storeData.NewStoreData(
		label,
		data,
		color,
	)

	pu.plotRepository.InsertStoreData(store_data)

	plot_data := pu.plotRepository.GetPlotData()

	pu.plotRepository.PlotGraph(plot_data)
}
