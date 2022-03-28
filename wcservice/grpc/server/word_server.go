package server

import (
	"context"
	word_pb "src/wcservice/grpc/pb/word"
	"src/wcservice/service"
)

func WordServer(service service.WordService) *wordServer {
	return &wordServer{
		Service: service,
	}
}

type wordServer struct {
	Service service.WordService
	word_pb.UnimplementedWordServer
}

func (s *wordServer) GetWords(ctx context.Context, request *word_pb.Request) (*word_pb.Response, error) {

	wordCountList := s.Service.GetWords(request.Text)

	return &word_pb.Response{
		WordCountList: wordCountList,
	}, nil
}