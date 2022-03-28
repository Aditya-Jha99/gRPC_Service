package server

import (
	"context"
	word_pb "wcservice/grpc/pb/word"
	"wcservice/service"
	"reflect"
	"testing"
)

func Test_wordServer_GetWords10(t *testing.T) {
	type fields struct {
		Service                 service.WordService
		UnimplementedWordServer word_pb.UnimplementedWordServer
	}
	type args struct {
		ctx     context.Context
		request *word_pb.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *word_pb.Response
		wantErr bool
	}{
		{
			"Success",
			fields{
				Service: service.WordServiceN_(),
			},
			args{
				ctx: context.TODO(),
				request: &word_pb.Request{
					Text: "a bb bb ccc ccc ccc",
				},
			},
			&word_pb.Response{
				WordCountList: []*word_pb.WordCount{
					{
						Word:  "ccc",
						Count: 3,
					},
					{
						Word:  "bb",
						Count: 2,
					},
					{
						Word:  "a",
						Count: 1,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := WordServer_(
				tt.fields.Service,
			)
			got, err := s.GetWords10(tt.args.ctx, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("wordServer.GetWords10() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordServer.GetWords10() = %v, want %v", got, tt.want)
			}
		})
	}
}
