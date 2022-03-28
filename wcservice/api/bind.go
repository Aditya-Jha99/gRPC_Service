package api

import (
	"wcservice/api/handler"
	"wcservice/service"
)

func bindWord() handler.WordHandler {
	service := service.WordServiceN_()
	return handler.NewWordHandler(service)
}
