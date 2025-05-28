package main

import (
	"common"
	"log"
	"net/http"

	pb "github.com/haileyism/commons"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)

}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")
	
	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r,&items); err != nil{
		common.WriteJSON(w,http.StatusBadRequest, err.Error())
		return 
	}



	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		customerID : customerID,
		Items: items,
	})

}
