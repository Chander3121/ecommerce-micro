package clients

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "order-service/proto"
)

func GetProductGrpc(
	productID int32,
) (*pb.ProductResponse, error) {

	conn, err := grpc.Dial(
		"product-service:50051",
		grpc.WithInsecure(),
	)

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	client := pb.NewProductServiceClient(
		conn,
	)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Second,
	)

	defer cancel()

	response, err := client.GetProduct(
		ctx,
		&pb.ProductRequest{
			Id: productID,
		},
	)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return response, nil
}
