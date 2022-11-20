package plotData

type PlotData struct {
	data  [][]float64
	color []string
}

func NewPlotData(data [][]float64, color []string) *PlotData {
	plotData := PlotData{
		data:  data,
		color: color,
	}
	return &plotData
}

func (p *PlotData) GetData() [][]float64 {
	return p.data
}

func (p *PlotData) GetColor() []string {
	return p.color
}

func (p *PlotData) Append(data []float64, color string) {
	p.data = append(p.data, data)
	p.color = append(p.color, color)
}
