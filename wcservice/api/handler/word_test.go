package handler

import (
	service "wcservice/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_wordHandler_GetWords10(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()

	wordHandler := NewWordHandler(service.WordServiceN_())
	r.POST("Top10", wordHandler.GetWords10)

	type args struct {
		Text string
	}
	tests := []struct {
		texts       string
		args       args
		StatusCode int
	}{
		{
			"Success",
			args{
				Text: "a a b",
			},
			100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.texts, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "Top10", strings.NewReader(tt.args.Text))
			responses := httptest.NewRecorder()
			r.ServeHTTP(responses, req)
			if responses.Code != tt.StatusCode {
				t.Errorf("wordServer.GetWords10() = %v, want %v", responses.Code, tt.StatusCode)
			}
		})
	}
}
