package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/guptarohit/asciigraph"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type PlotData struct {
	data []float64
}

func NewPlotData(data []float64) *PlotData {
	plotData := PlotData{
		data: data,
	}
	return &plotData
}

func (p PlotData) GetData() []float64 {
	return p.data
}

func main() {

	template := "%{time_total}"

	url := os.Args[1]
	data := []float64{}

	for {

		time.Sleep(time.Second * 1)
		out, err := exec.Command("curl", url, "-w", template, "-o", "/dev/null").Output()
		if err != nil {
			fmt.Println(err)
		}

		float, err := strconv.ParseFloat(string(out), 64)

		if len(data) > 69 {
			data = data[1:]
		}

		data = append(data, float)

		plotData := NewPlotData(data)

		fmt.Print("\033[H\033[2J")
		graph := asciigraph.Plot(plotData.GetData(), asciigraph.Height(10), asciigraph.Precision(6))

		title := figure.NewFigure("Cloud Watch CLI", "", true)
		title.Print()

		fmt.Println("------------------------------------------------------------------------------------------")
		fmt.Println(graph)
		fmt.Println("------------------------------------------------------------------------------------------")

	}
}
