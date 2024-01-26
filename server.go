package main

import (
	"backend_02/food/proto"
	"context"
)

type Server struct {
	proto.UnimplementedUserServiceServer
}

func (s *Server) GetAllUsers(ctx context.Context, req *proto.FoodRequest) (*proto.FoodList, error) {
	println("GetAllUsers", req)
	foodlist := &proto.FoodList{
		Foods: []*proto.Food{
			{
				Name:        "Burger",
				Description: "Burger Description",
				Image:       "Burger Image",
				Price:       "Burger Price",
				Category:    "Burger Category",
				Restaurant:  "Burger Restaurant",
				Id:          "Burger Id",
			},
		},
	}
	return foodlist, nil
}
