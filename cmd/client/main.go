package main

import (
	"context"
	"flag"
	"log"
	"time"

	v1 "github.com/ehsaaniqbal/dixdb/pkg/api/v1"
	"google.golang.org/grpc"
)

func main() {
	addr := flag.String("server", "", "host:port")
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}

	defer conn.Close()

	c := v1.NewDbServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req1 := v1.SetRequest{
		Key:   "hello",
		Value: "world",
	}

	res1, err := c.Set(ctx, &req1)
	if err != nil {
		log.Fatalf("Set failed: %v", err)
	}

	log.Printf("Set result: <%v>\n\n", res1.Status)

	req2 := v1.GetRequest{
		Key: "hello",
	}

	res2, err := c.Get(ctx, &req2)
	if err != nil {
		log.Fatalf("Get failed: %v", err)
	}

	log.Printf("Get result: <%+v>\n\n", res2)

	req3 := v1.DeleteRequest{
		Key: "hello",
	}
	res3, err := c.Delete(ctx, &req3)
	if err != nil {
		log.Fatalf("Delete failed: %v", err)
	}

	log.Printf("Delete result: <%v>\n\n", res3.Status)

}
