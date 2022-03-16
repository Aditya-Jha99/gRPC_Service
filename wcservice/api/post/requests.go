package post

import (
	"fmt"
	"log"
	"net/http"
	"wcservice/api/handler"
)

func homepage(m http.ResponseWriter, r *http.Request) {
	fmt.Fprint(m, "Aditya Jha Come to My API")

}
func Request() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/Word", handler.PostRequest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
