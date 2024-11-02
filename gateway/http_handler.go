package main

import (
	common "commons"
	pb "commons/api"
	"errors"
	"log"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

    if err := validateItems(items); err != nil {
        log.Printf("Error validating items: %v", err)

        common.WriteError(w, http.StatusBadRequest, err.Error())
        return
    }


    resp, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
        CustomerID: customerID,
        Items:      items,
    })

    rStatus:=status.Convert(err)
    if rStatus != nil {
        if rStatus.Code() != codes.InvalidArgument{
            common.WriteError(w, http.StatusBadRequest, rStatus.Message())
            return
        }

        log.Printf("Error calling CreateOrder: %v", err)
        common.WriteError(w, http.StatusInternalServerError, err.Error())
        return
    }

    common.WriteJSON(w, http.StatusOK, resp)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
   if len(items) == 0 {
       return common.ErrNoItems}

    for _, item := range items {
        if item.ID == "" {
            return errors.New("item ID is required")
        }
        if item.Quantity <= 0 {
            return errors.New("item quantity must be greater than 0")
        }
    }



    return nil
}