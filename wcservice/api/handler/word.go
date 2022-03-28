package handler

import (
	"wcservice/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WordHandler interface {
	GetWords10(c *gin.Context)
}

type wordHandler struct {
	Service service.WordService
}

func NewWordHandler(service service.WordService) WordHandler {
	return &wordHandler{
		Service: service,
	}
}

func (h *wordHandler) GetWords10(c *gin.Context) {

	bytes, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response := h.Service.GetWords10(string(bytes))

	c.JSON(http.StatusOK, response)
}
