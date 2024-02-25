package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/ruan-molinari/calculator-gRPC/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Specify a CLI flag
	serverAddr := flag.String(
		"server", "localhost:8000",
		"The server address in the format of host:port",
	)
	flag.Parse()

	// Enabling TLS
	// creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})
	creds := insecure.NewCredentials()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Creating a gRPC connection
	conn, err := grpc.DialContext(ctx, *serverAddr, opts...)
	if err != nil {
		log.Fatalln("fail to dial: ", err)
	}
	defer conn.Close()

	// Creating a new instance of the Calculator client on the gRPC connection
	client := pb.NewCalculatorClient(conn)
	res, err := client.Sum(ctx, &pb.RepeatedInput{
		Numbers: []int64{1, 2, 3, 4, 5},
	})
	if err != nil {
		log.Fatalln("error sending request: ", err)
	}
	log.Println(res)
}
