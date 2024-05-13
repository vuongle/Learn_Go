package main

import (
	common "go-microservices/commons"
	pb "go-microservices/commons/api"
	"log"
	"net/http"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc(
		"POST /api/customers/{customerID}/orders",
		h.HandleCreateOrder,
	)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("HandleCreateOrder CALLED")
	customerID := r.PathValue("customerID")
	log.Println("customerID: ", customerID)
	var items []*pb.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
}
