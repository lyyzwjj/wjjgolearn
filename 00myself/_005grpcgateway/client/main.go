package main

import (
	"context"
	trippb "coolcar/proto/gen/go"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect server: %v", err)
	}
	tsClient := trippb.NewTripServiceClient(conn)
	trip, err := tsClient.GetTrip(context.Background(), &trippb.GetTripRequest{
		Id: "trip467",
	})
	if err != nil {
		log.Fatalf("cannot call GetTrip: %v", err)
	}
	fmt.Println(trip)
}
