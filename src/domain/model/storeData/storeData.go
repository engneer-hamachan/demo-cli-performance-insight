package storeData

type StoreData struct {
	label string
	data  float64
	color string
}

func NewStoreData(label string, data float64, color string) *StoreData {
	storeData := StoreData{
		label: label,
		data:  data,
		color: color,
	}

	return &storeData
}

func (s *StoreData) GetLabel() string {
	return s.label
}

func (s *StoreData) GetData() float64 {
	return s.data
}

func (s *StoreData) GetColor() string {
	return s.color
}
