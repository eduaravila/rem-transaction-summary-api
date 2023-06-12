package ports

import (
	"net/http"

	"github.com/eduaravila/stori-challenge/pkg/summary/app"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(app app.Application) *HTTPServer {
	return &HTTPServer{
		app: app,
	}
}

func (h *HTTPServer) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPServer) CreateTransaction(w http.ResponseWriter, r *http.Request, userID string) {
}

func (h *HTTPServer) GetTransactions(w http.ResponseWriter, r *http.Request, userID string) {

}
