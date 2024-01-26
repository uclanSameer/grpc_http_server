package main

import (
	"backend_02/food/proto"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	// run grpc server first
	go RunGRPCServer()

	// run http server later
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello World!")
		if err != nil {
			panic(err)
		}
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}

func RunGRPCServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
