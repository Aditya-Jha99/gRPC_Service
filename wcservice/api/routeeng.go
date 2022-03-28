package api

import (
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	ro := gin.Default()

	wordHandler := bindWord()
	ro.POST("Top10", wordHandler.GetWords10)

	return ro
}
