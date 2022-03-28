package grpc

import (
	word_pb "wcservice/grpc/pb/word"
	"wcservice/grpc/server"
	"wcservice/service"
)

func bindWord() word_pb.WordServer {
	service := service.WordServiceN_()
	return server.WordServer_(service)
}
