package repository

type HttpRequestRepository interface {
	GetSiteSpeed(url string) (speed *float64, err error)
}
