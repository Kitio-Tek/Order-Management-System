package main

import (
	common "commons"
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "commons/api"
)

var(
	httpAddr = common.EnvString("HTTP_ADDR", "localhost:8080")
	ordersServiceAddr= "localhost:2000"

   
)

func main() {

	conn, err := grpc.Dial(ordersServiceAddr,grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to orders service: %v", err)
    }
	defer conn.Close()
    
	log.Println("Dialing orders service at ", ordersServiceAddr)

	c:=	pb.NewOrderServiceClient(conn)
	
	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server on %s...\n", httpAddr)
    
	if err:= http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server: ", err)
	}
}