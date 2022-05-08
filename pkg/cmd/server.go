package cmd

import (
	"context"
	"flag"
	"fmt"

	m "github.com/ehsaaniqbal/dixdb/pkg/db/memory"
	grpc "github.com/ehsaaniqbal/dixdb/pkg/protocol/grpc"
	v1 "github.com/ehsaaniqbal/dixdb/pkg/service/v1"
)

type ServerConfig struct {
	// TCP port to listen on
	Port string
}

func RunServer() error {
	ctx := context.Background()

	var config ServerConfig

	flag.StringVar(&config.Port, "port", "", "TCP port to bind")
	flag.Parse()

	if len(config.Port) == 0 {
		return fmt.Errorf("nvalid TCP port provided: '%s", config.Port)
	}

	db := *m.NewMemoryStore()

	v1API := v1.NewServer(db)

	return grpc.RunServer(ctx, v1API, config.Port)

}
