package usecase

import (
	"main/src/domain/model/plotData"
	"main/src/domain/repository"
	"time"
)

type CloudWatchUseCase interface {
	CheckSpeedInsight(url string) error
}

type cloudWatchUseCase struct {
	plotRepository        repository.PlotRepository
	httpRequestRepository repository.HttpRequestRepository
}

func NewCloudWatchUseCase(pr repository.PlotRepository, hrr repository.HttpRequestRepository) CloudWatchUseCase {
	return &cloudWatchUseCase{
		plotRepository:        pr,
		httpRequestRepository: hrr,
	}
}

func (cwu *cloudWatchUseCase) CheckSpeedInsight(url string) error {

	plotData := plotData.NewPlotData([]float64{})

	for {
		time.Sleep(time.Second * 1)

		speed, err := cwu.httpRequestRepository.GetSiteSpeed(url)
		if err != nil {
			return err
		}

		plotData = plotData.Append(*speed)
		cwu.plotRepository.PlotSpeedInsight(plotData)
	}
}
