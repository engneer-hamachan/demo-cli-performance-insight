package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main/src/usecase"
	"net/http"
	"strconv"
)

type CloudWatchHandler interface {
	PlotData(c *gin.Context)
}

type cloudWatchHandler struct {
	cloudWatchUseCase usecase.CloudWatchUseCase
}

func NewCloudWatchHandler(cwu usecase.CloudWatchUseCase) CloudWatchHandler {
	return &cloudWatchHandler{
		cloudWatchUseCase: cwu,
	}
}

func (cwh *cloudWatchHandler) PlotData(c *gin.Context) {

	label := c.Param("label")
	data, _ := strconv.ParseFloat(c.Param("data"), 64)
	color := c.Param("color")

	fmt.Println(data)

	cwh.cloudWatchUseCase.StoreData(label, float64(data), color)

	c.IndentedJSON(http.StatusOK, 1)
}
