package persistence

import (
	"main/src/domain/repository"
	"os/exec"
	"strconv"
)

type httpRequestPersistence struct{}

func NewHttpRequestPersistence() repository.HttpRequestRepository {
	return &httpRequestPersistence{}
}

func (rh *httpRequestPersistence) GetSiteSpeed(url string) (speed *float64, err error) {

	out, err := exec.Command("curl", url, "-w", "%{time_total}", "-o", "/dev/null").Output()
	if err != nil {
		return nil, err
	}

	float, err := strconv.ParseFloat(string(out), 64)
	if err != nil {
		return nil, err
	}

	return &float, nil
}
