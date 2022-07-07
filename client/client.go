package main

import (
	"context"
	"enigmacamp.com/go_grpc_bidirect/api"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(":50005", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	client := api.NewReportingClient(conn)
	stream, err := client.GenerateStatement(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	ctx := stream.Context()
	done := make(chan bool)
	go func() {
		for i := 1; i <= 10; i++ {
			req := api.Request{ClientId: fmt.Sprintf("Client %d", i)}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			log.Printf("Client %d sent", i)
			time.Sleep(time.Millisecond * 500)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			result := resp.Result
			log.Printf("Status %s received", result)
		}
	}()
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	//log.Printf("finished with max=%d", max)
}
