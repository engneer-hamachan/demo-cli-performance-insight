package handler

import (
	"main/src/usecase"
)

type CloudWatchHandler interface {
	RunCheckSpeedInsight(url string)
}

type cloudWatchHandler struct {
	cloudWatchUseCase usecase.CloudWatchUseCase
}

func NewCloudWatchHandler(cwu usecase.CloudWatchUseCase) CloudWatchHandler {
	return &cloudWatchHandler{
		cloudWatchUseCase: cwu,
	}
}

func (cwh *cloudWatchHandler) RunCheckSpeedInsight(url string) {
	cwh.cloudWatchUseCase.CheckSpeedInsight(url)
}
