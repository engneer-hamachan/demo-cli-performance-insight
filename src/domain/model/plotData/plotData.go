package plotData

type PlotData struct {
	data []float64
}

func NewPlotData(data []float64) *PlotData {
	plotData := PlotData{
		data: data,
	}
	return &plotData
}

func (p *PlotData) GetData() []float64 {
	return p.data
}

func (p *PlotData) Append(data float64) {

	if len(p.data) > 78 {
		p.data = p.data[1:]
	}

	p.data = append(p.data, data)
}
