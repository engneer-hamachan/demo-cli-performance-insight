package vo

type TimeStamp int64

func NewTimeStamp(value int64) (*TimeStamp, error) {
	timestamp := TimeStamp(value)
	return &timestamp, nil
}
