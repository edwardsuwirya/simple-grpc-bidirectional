package main

import (
	"enigmacamp.com/go_grpc_bidirect/api"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

type reportingServer struct {
	api.UnimplementedReportingServer
}

func (s reportingServer) GenerateStatement(srv api.Reporting_GenerateStatementServer) error {
	log.Println("start new server")
	for {
		rnd := int32(rand.Intn(5))
		fmt.Println("Random ", rnd)
		req, err := srv.Recv()
		if err == io.EOF {
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			return err
		}
		time.Sleep(time.Second * time.Duration(rnd))
		resp := api.Response{Result: "Done processing " + req.ClientId}
		if err := srv.Send(&resp); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("send result=%s", req.ClientId)
	}

}
func main() {
	rand.Seed(time.Now().Unix())
	lis, err := net.Listen("tcp", "localhost:50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterReportingServer(s, &reportingServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
