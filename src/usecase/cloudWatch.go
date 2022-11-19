package usecase

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
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

	m := map[string][]float64{}
	mc := map[string]asciigraph.AnsiColor{}

	for _, d := range store_datas {
		m[d.GetLabel()] = append(m[d.GetLabel()], d.GetData())
		mc[d.GetLabel()] = asciigraph.ColorNames[d.GetColor()]
	}

	plotData := plotData.NewPlotData([][]float64{})
	colorData := []asciigraph.AnsiColor{}

	for k, v := range m {
		ct := 1
		mm := []float64{}
		start_time := float64(0)
		fmt.Println(v)
		for _, i := range v {
			if ct%2 == 0 {
				mm = append(mm, i-start_time)
			} else {
				start_time = i
			}
			ct += 1
		}
		plotData.Append(mm)
		colorData = append(colorData, mc[k])
	}

	cwu.plotRepository.PlotSpeedInsight(plotData, colorData)

	return nil
}
