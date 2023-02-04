package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/mavz49/ping/ping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	t1 := time.Now() // Метка 1
	fmt.Println("t1 =", t1)

	r, err := c.SayPing(ctx, &pb.PingRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	//time.Sleep(3 * time.Second) // Имитация фрагмента кода
	t2 := time.Now() // Метка 2
	fmt.Println("t2 =", t2)
	d := t2.Sub(t1) // Разница между метками
	fmt.Println("d =", d.Milliseconds())

	log.Printf("Greeting: %s", r.GetMessage())
	//log.Printf("Greeting: %s", r.GetMessage())
}

/**/
