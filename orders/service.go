package main

import (
	common "commons"
	pb "commons/api"
	"context"
	"log"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store: store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}

func (s *service) validateOrder(ctx context.Context, p *pb.CreateOrderRequest) error {

	if len(p.Items) == 0 {
		return common.ErrNoItems
	}

	mergedItems := mergeItemsQuantities(p.Items)

	log.Print("Merged items: ", mergedItems)

	// validate with the stock service 

	return nil

}

func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
 merged:=make([]*pb.ItemsWithQuantity,0)

 for _,item:=range items{
	 found:=false
	 for _,mItem:=range merged{
		 if mItem.ID==item.ID{
			 mItem.Quantity+=item.Quantity
			 found=true
			 break
		 }
	 }
	 if !found{
		 merged=append(merged,item)
	 }
 }

 return merged
}