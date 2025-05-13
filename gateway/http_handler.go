package main

import ("net/http"
		"log"
)

type handler struct {
	//gateway, help service discover
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)

}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("hello")
}
