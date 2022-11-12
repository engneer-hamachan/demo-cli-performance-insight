package benchmark

import (
	"github.com/google/uuid"
	"main/src/domain/model/vo"
)

type BenchMark struct {
	benchMarkId    vo.UuId
	label          vo.Label
	startTimeStamp vo.TimeStamp
	endTimeStamp   vo.TimeStamp
}

func New(benchMarkId string, label string, startTimeStamp int64, endTimeStamp int64) (*BenchMark, error) {

	createdBenchMarkId, err := vo.NewUuId(benchMarkId)
	if err != nil {
		return nil, err
	}

	createdLabel, err := vo.NewLabel(label)
	if err != nil {
		return nil, err
	}

	createdStartTimeStamp, err := vo.NewTimeStamp(startTimeStamp)
	if err != nil {
		return nil, err
	}

	createdEndTimeStamp, err := vo.NewTimeStamp(endTimeStamp)
	if err != nil {
		return nil, err
	}

	benchMark := BenchMark{
		benchMarkId:    *createdBenchMarkId,
		label:          *createdLabel,
		startTimeStamp: *createdStartTimeStamp,
		endTimeStamp:   *createdEndTimeStamp,
	}
	return &benchMark, nil

}

// Create Constructor
func Create(label string, startTimeStamp int64, endTimeStamp int64) (*BenchMark, error) {
	benchMarkId := uuid.New().String()
	benchMark, err := New(benchMarkId, label, startTimeStamp, endTimeStamp)
	if err != nil {
		return nil, err
	}

	return benchMark, err
}

func (b BenchMark) GetBenchMarkId() vo.UuId {
	return b.benchMarkId
}

func (b BenchMark) GetLabel() vo.Label {
	return b.label
}

func (b BenchMark) GetStartTimeStamp() vo.TimeStamp {
	return b.startTimeStamp
}

func (b BenchMark) GetEndTimeStamp() vo.TimeStamp {
	return b.endTimeStamp
}
