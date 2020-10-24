package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "pixielabs.ai/pixielabs/src/stirling/protocols/http2/testing/proto"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) SayHelloClientStreaming(srv pb.StreamingGreeter_SayHelloClientStreamingServer) error {
	names := []string{}
	for {
		helloReq, err := srv.Recv()
		if err == io.EOF {
			return srv.SendAndClose(&pb.HelloReply{Message: "Hello " + strings.Join(names, ", ") + "!"})
		}
		if err != nil {
			return err
		}
		names = append(names, helloReq.Name)
	}
}

func (s *server) SayHelloServerStreaming(in *pb.HelloRequest, srv pb.StreamingGreeter_SayHelloServerStreamingServer) error {
	log.Printf("SayHelloServerStreaming, in: %v\n", in)
	// Send 3 responses each time. We do not care much about the exact number of responses, this is for executing the
	// server streaming mechanism and observe the underlying HTTP2 framing data.
	for i := 0; i < 3; i++ {
		srv.Send(&pb.HelloReply{Message: "Hello " + in.Name})
	}
	return nil
}

func (s *server) SayHelloBidirStreaming(stream pb.StreamingGreeter_SayHelloBidirStreamingServer) error {
	for {
		helloReq, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		stream.Send(&pb.HelloReply{Message: "Hello " + helloReq.Name})
	}
}

func main() {
	var port = flag.Int("port", 0, "The port to listen.")
	var https = flag.Bool("https", false, "Whether or not to use https")
	var cert = flag.String("cert", "", "Path to the .crt file.")
	var key = flag.String("key", "", "Path to the .key file.")
	var streaming = flag.Bool("streaming", false, "Whether or not to call streaming RPC")

	const keyPairBase = "src/stirling/protocols/http2/testing/go_grpc_server"

	flag.Parse()

	portStr := ":" + strconv.Itoa(*port)

	var lis net.Listener
	var err error
	if *https {
		certFile := keyPairBase + "/https-server.crt"
		if len(*cert) > 0 {
			certFile = *cert
		}
		keyFile := keyPairBase + "/https-server.key"
		if len(*key) > 0 {
			keyFile = *key
		}
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			log.Fatalf("failed to load certs: %v", err)
		}
		tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}

		log.Printf("Starting https server on port : %s cert: %s key: %s", portStr, certFile, keyFile)
		lis, err = tls.Listen("tcp", portStr, tlsConfig)
	} else {
		log.Printf("Starting http server on port : %s", portStr)
		lis, err = net.Listen("tcp", portStr)
	}

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Print(lis.Addr().(*net.TCPAddr).Port)

	s := grpc.NewServer()

	if *streaming {
		log.Printf("Launching streaming server")
		pb.RegisterStreamingGreeterServer(s, &server{})
	} else {
		log.Printf("Launching unary server")
		pb.RegisterGreeterServer(s, &server{})
	}
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}