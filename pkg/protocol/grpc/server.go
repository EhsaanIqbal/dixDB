package grpc

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	v1 "github.com/ehsaaniqbal/dixdb/pkg/api/v1"
	"google.golang.org/grpc"
)

func RunServer(ctx context.Context, v1API v1.DbServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// register grpc service
	server := grpc.NewServer()
	v1.RegisterDbServiceServer(server, v1API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			log.Println("[INFO] Shuting down server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start server
	log.Println("[INFO] Starting server on PORT: " + port)
	return server.Serve(listen)
}
