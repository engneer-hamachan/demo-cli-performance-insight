package dto

import (
	"main/src/domain/model/storeData"
	"time"
)

type StoreData struct {
	Label     string
	Data      float64
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StoreDatas struct {
	Data []StoreData
}

func ConvertStoreData(s *storeData.StoreData) *StoreData {
	return &StoreData{
		Label: s.GetLabel(),
		Data:  s.GetData(),
		Color: s.GetColor(),
	}
}

func AdaptStoreDatas(converted_store_datas *StoreDatas) []storeData.StoreData {
	var store_datas []storeData.StoreData

	for _, converted_store_data := range converted_store_datas.Data {
		d := *storeData.NewStoreData(
			converted_store_data.Label,
			converted_store_data.Data,
			converted_store_data.Color,
		)
		store_datas = append(store_datas, d)
	}

	return store_datas
}
