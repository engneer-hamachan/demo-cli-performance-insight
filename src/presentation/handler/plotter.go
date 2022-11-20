package handler

import (
	"github.com/gin-gonic/gin"
	"main/src/usecase"
	"net/http"
	"strconv"
)

type PlotterHandler interface {
	RecieveData(c *gin.Context)
}

type plotterHandler struct {
	plotterUseCase usecase.PlotterUseCase
}

func NewPlotterHandler(pu usecase.PlotterUseCase) PlotterHandler {
	return &plotterHandler{
		plotterUseCase: pu,
	}
}

func (ph *plotterHandler) RecieveData(c *gin.Context) {

	label := c.Param("label")
	data, _ := strconv.ParseFloat(c.Param("data"), 64)
	color := c.Param("color")

	ph.plotterUseCase.RecieveData(label, float64(data), color)

	c.IndentedJSON(http.StatusOK, 1)
}
