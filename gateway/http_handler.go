package main

import (
	common "commons"
	pb "commons/api"
	"log"
	"net/http"
)

type handler struct {
	// gateway
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client: client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
    customerID := r.PathValue("customerID")
    log.Printf("Received request to create order for customerID: %s", customerID)

    var items []*pb.ItemsWithQuantity
    if err := common.ReadJSON(r, &items); err != nil {
        log.Printf("Error reading JSON: %v", err)
        common.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }

    resp, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
        CustomerID: customerID,
        Items:      items,
    })
    if err != nil {
        log.Printf("Error calling CreateOrder: %v", err)
        common.WriteError(w, http.StatusInternalServerError, err.Error())
        return
    }

    common.WriteJSON(w, http.StatusOK, resp)
}