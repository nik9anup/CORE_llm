// Setup a gRPC server with unary RPC
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// Setup a gRPC client and call a unary RPC
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}




// Implement server streaming RPC
package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type StreamerServer struct{}

func (s *StreamerServer) CountNumbers(req *NumberRequest, stream Streamer_CountNumbersServer) error {
	for i := 1; i <= int(req.Number); i++ {
		if err := stream.Send(&NumberResponse{Result: int64(i)}); err != nil {
			return err
		}
		time.Sleep(time.Second) // Simulating some processing time
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterStreamerServer(s, &StreamerServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// Implement client streaming RPC
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewStreamerClient(conn)
	stream, err := client.RecordNumbers(context.Background())
	if err != nil {
		log.Fatalf("error recording numbers: %v", err)
	}

	for _, number := range []int64{1, 2, 3, 4, 5} {
		if err := stream.Send(&NumberRequest{Number: number}); err != nil {
			log.Fatalf("error sending request: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error closing stream: %v", err)
	}

	log.Printf("Sum of numbers: %d", resp.Sum)
}




// Implement bidirectional streaming RPC
package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type BidirectionalServer struct{}

func (s *BidirectionalServer) Communicate(stream Bidirectional_CommunicateServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received message: %s", req.Message)

		if err := stream.Send(&BidirectionalResponse{Message: "Hello, " + req.Message}); err != nil {
			return err
		}
		time.Sleep(time.Second) // Simulating some processing time
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterBidirectionalServer(s, &BidirectionalServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// Implement unary RPC with deadline
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &HelloRequest{Name: "Bob"})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			log.Fatalf("timeout: %v", statusErr.Message())
		}
		log.Fatalf("error: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}




// Implement server-side interceptor
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Unary RPC call: %s", info.FullMethod)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("Unary RPC error: %v", err)
	}
	return resp, err
}

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// Implement client-side interceptor
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("Sending RPC request: %s", method)
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			log.Printf("RPC request timed out: %v", err)
		} else {
			log.Printf("RPC request error: %v", err)
		}
	}
	return err
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(clientInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &HelloRequest{Name: "Charlie"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}




// Implement RPC with retry and deadline
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func retryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	for {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err == nil || ctx.Err() != nil {
			return err
		}
		log.Printf("RPC failed: %v. Retrying...", err)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(retryInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &HelloRequest{Name: "Charlie"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}







// Implement gRPC reflection service
package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	reflection.Register(s)
	log.Println("gRPC server with reflection started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}



// Setup a gRPC server with TLS
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	certificate, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("could not load server key pair: %v", err)
	}
	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatalf("could not read ca certificate: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(bs); !ok {
		log.Fatalf("failed to append client certs")
	}
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server with TLS started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// Setup a gRPC client with TLS
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		log.Fatalf("could not read ca certificate: %v", err)
	}
	if ok := certPool.AppendCertsFromPEM(bs); !ok {
		log.Fatalf("failed to append client certs")
	}
	creds := credentials.NewTLS(&tls.Config{
		RootCAs: certPool,
	})
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	resp, err := client.SayHello(context.Background(), &HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}





// Error handling in gRPC server
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name is required")
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// Unary RPC with metadata
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("Received metadata: %v", md)
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// Unary RPC client with metadata
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	md := metadata.Pairs("authorization", "Bearer some-token")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	resp, err := client.SayHello(ctx, &HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}




package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func authInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata not provided")
	}

	token := md["authorization"]
	if len(token) == 0 || token[0] != "Bearer some-token" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return handler(ctx, req)
}

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor),
	)
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// Server-side streaming RPC with authentication
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type StreamerServer struct{}

func (s *StreamerServer) ListNumbers(req *NumberRequest, stream Streamer_ListNumbersServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata not provided")
	}

	token := md["authorization"]
	if len(token) == 0 || token[0] != "Bearer some-token" {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	for i := 1; i <= int(req.Number); i++ {
		if err := stream.Send(&NumberResponse{Result: int64(i)}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterStreamerServer(s, &StreamerServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// Client-side streaming RPC with authentication
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewStreamerClient(conn)
	md := metadata.Pairs("authorization", "Bearer some-token")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	stream, err := client.RecordNumbers(ctx)
	if err != nil {
		log.Fatalf("error recording numbers: %v", err)
	}

	for _, number := range []int64{1, 2, 3, 4, 5} {
		if err := stream.Send(&NumberRequest{Number: number}); err != nil {
			log.Fatalf("error sending request: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error closing stream: %v", err)
	}

	log.Printf("Sum of numbers: %d", resp.Sum)
}





// Bidirectional streaming RPC with authentication
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type BidirectionalServer struct{}

func (s *BidirectionalServer) Chat(stream Bidirectional_ChatServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata not provided")
	}

	token := md["authorization"]
	if len(token) == 0 || token[0] != "Bearer some-token" {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received message: %s", req.Message)

		if err := stream.Send(&BidirectionalResponse{Message: "Hello, " + req.Message}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterBidirectionalServer(s, &BidirectionalServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// Cancel RPC context to terminate ongoing operation
package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	select {
	case <-time.After(time.Second * 3):
		return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// Unary RPC with client-side cancellation
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &HelloRequest{Name: "Bob"})
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			log.Fatalf("timeout: %v", statusErr.Message())
		}
		log.Fatalf("error: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}




// Unary RPC with server-side context cancellation
package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	select {
	case <-time.After(time.Second * 5):
		return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// Unary RPC with context propagation
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("Received metadata: %v", md)
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// Streaming RPC with deadline
package main

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StreamerServer struct{}

func (s *StreamerServer) RecordNumbers(stream Streamer_RecordNumbersServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received number: %d", req.Number)

		// Simulate processing time
		select {
		case <-stream.Context().Done():
			return status.Error(codes.DeadlineExceeded, "client deadline exceeded")
		default:
			if err := stream.Send(&NumberResponse{Result: req.Number * req.Number}); err != nil {
				return err
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterStreamerServer(s, &StreamerServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// Client-side streaming RPC with deadline
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewStreamerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stream, err := client.RecordNumbers(ctx)
	if err != nil {
		log.Fatalf("error recording numbers: %v", err)
	}

	for _, number := range []int64{1, 2, 3, 4, 5} {
		if err := stream.Send(&NumberRequest{Number: number}); err != nil {
			log.Fatalf("error sending request: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error closing stream: %v", err)
	}

	log.Printf("Sum of squares: %d", resp.Sum)
}




// Bidirectional streaming RPC with deadline
package main

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BidirectionalServer struct{}

func (s *BidirectionalServer) Chat(stream Bidirectional_ChatServer) error {
	ctx := stream.Context()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			req, err := stream.Recv()
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}
			log.Printf("Received message: %s", req.Message)

			if err := stream.Send(&BidirectionalResponse{Message: "Hello, " + req.Message}); err != nil {
				return err
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterBidirectionalServer(s, &BidirectionalServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// Unary RPC with custom error handling
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Name cannot be empty")
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// Server-side streaming RPC with custom error handling
package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StreamerServer struct{}

func (s *StreamerServer) ListNumbers(req *NumberRequest, stream Streamer_ListNumbersServer) error {
	if req.Number <= 0 {
		return status.Error(codes.InvalidArgument, "Number should be greater than zero")
	}

	for i := 1; i <= int(req.Number); i++ {
		if err := stream.Send(&NumberResponse{Result: int64(i)}); err != nil {
			return err
		}
		time.Sleep(time.Second) // Simulate processing time
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterStreamerServer(s, &StreamerServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// Client-side streaming RPC with custom error handling
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewStreamerClient(conn)
	ctx := context.Background()
	stream, err := client.RecordNumbers(ctx)
	if err != nil {
		log.Fatalf("error recording numbers: %v", err)
	}

	for _, number := range []int64{0, -1, 2, 3, 4} { // Sending invalid numbers
		if err := stream.Send(&NumberRequest{Number: number}); err != nil {
			log.Fatalf("error sending request: %v", err)
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.InvalidArgument {
			log.Fatalf("invalid argument: %v", statusErr.Message())
		}
		log.Fatalf("error closing stream: %v", err)
	}

	log.Printf("Sum of squares: %d", resp.Sum)
}





// Bidirectional streaming RPC with custom error handling
package main

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BidirectionalServer struct{}

func (s *BidirectionalServer) Chat(stream Bidirectional_ChatServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received message: %s", req.Message)

		if req.Message == "error" {
			return status.Error(codes.InvalidArgument, "Received 'error' message")
		}

		if err := stream.Send(&BidirectionalResponse{Message: "Hello, " + req.Message}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterBidirectionalServer(s, &BidirectionalServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// Handling gRPC errors on client-side
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &HelloRequest{Name: ""}) // Sending empty name
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.InvalidArgument {
			log.Fatalf("invalid argument: %v", statusErr.Message())
		}
		log.Fatalf("error: %v", err)
	}

	log.Printf("Greeting: %s", resp.Message)
}





// Unary RPC with error propagation from server
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "error" {
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// Unary RPC with error handling and logging
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "" {
		err := status.Error(codes.InvalidArgument, "Name cannot be empty")
		log.Printf("Invalid request: %v", err)
		return nil, err
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}






// Streaming RPC with error handling and logging
package main

import (
	"context"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StreamerServer struct{}

func (s *StreamerServer) RecordNumbers(stream Streamer_RecordNumbersServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Received number: %d", req.Number)

		if req.Number <= 0 {
			err := status.Error(codes.InvalidArgument, "Number should be greater than zero")
			log.Printf("Invalid request: %v", err)
			return err
		}

		if err := stream.Send(&NumberResponse{Result: req.Number * req.Number}); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterStreamerServer(s, &StreamerServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// Handling gRPC errors with retry mechanism
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	retryAttempts := 3
	for attempt := 1; attempt <= retryAttempts; attempt++ {
		client := NewGreeterClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := client.SayHello(ctx, &HelloRequest{Name: "error"}) // Simulating error scenario
		if err != nil {
			statusErr, ok := status.FromError(err)
			if ok && statusErr.Code() == codes.Internal {
				log.Printf("Attempt %d: Internal server error - %v", attempt, statusErr.Message())
				if attempt < retryAttempts {
					log.Println("Retrying...")
					time.Sleep(2 * time.Second)
					continue
				}
			}
			log.Fatalf("error: %v", err)
		}

		log.Printf("Greeting: %s", resp.Message)
		break
	}
}





// Unary RPC with TLS encryption
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	certFile := "server.crt"
	keyFile := "server.key"
	caFile := "ca.crt"

	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load server certificate: %v", err)
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started with TLS on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// Code for client-side streaming RPC with TLS encryption


package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	certFile := "client.crt"
	keyFile := "client.key"
	caFile := "ca.crt"

	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load client certificate: %v", err)
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      caCertPool,
	})

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewStreamerClient(conn)
	// Perform client-side streaming RPC operations here...
}





// Bidirectional streaming RPC with TLS encryption
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	certFile := "client.crt"
	keyFile := "client.key"
	caFile := "ca.crt"

	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load client certificate: %v", err)
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      caCertPool,
	})

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewBidirectionalClient(conn)
	ctx := context.Background()
	stream, err := client.Chat(ctx)
	if err != nil {
		log.Fatalf("error opening stream: %v", err)
	}

	// Perform bidirectional streaming RPC operations here...
}





// Server-side streaming RPC with TLS encryption
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	certFile := "server.crt"
	keyFile := "server.key"
	caFile := "ca.crt"

	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load server certificate: %v", err)
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	RegisterStreamerServer(s, &StreamerServer{})
	log.Println("gRPC server started with TLS on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}






// gRPC with interceptors for logging
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	log.Printf("Received request: %v", req)
	if req.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "Name cannot be empty")
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("gRPC method: %s, request: %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("gRPC method: %s, error: %v", info.FullMethod, err)
	}
	return resp, err
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}







// gRPC with authentication interceptor
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func authInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	creds, err := credentials.ParseAuthorizationHeader(ctx, "Bearer")
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Invalid credentials")
	}
	// Perform authentication logic here with `creds`
	log.Printf("Authenticated with credentials: %v", creds)
	return handler(ctx, req)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// gRPC with custom metadata
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("Received metadata: %v", md)
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// gRPC with cancellation propagation
package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		time.Sleep(2 * time.Second) // Simulate processing time
		return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}






// gRPC with context metadata propagation
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("Received metadata: %v", md)
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}




// gRPC with deadline propagation on client-side
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Greeting: %s", resp.Message)
}




// gRPC with unary interceptor for logging
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	log.Printf("Received request: %v", req)
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("gRPC method: %s, request: %v", info.FullMethod, req)
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("gRPC method: %s, error: %v", info.FullMethod, err)
	}
	return resp, err
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}






// gRPC with server-side streaming and TLS encryption
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type StreamerServer struct{}

func (s *StreamerServer) RecordNumbers(req *NumberRequest, stream Streamer_RecordNumbersServer) error {
	for i := int64(1); i <= req.Number; i++ {
		if err := stream.Send(&NumberResponse{Result: i}); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	certFile := "server.crt"
	keyFile := "server.key"
	caFile := "ca.crt"

	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load server certificate: %v", err)
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	RegisterStreamerServer(s, &StreamerServer{})
	log.Println("gRPC server started with TLS on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// gRPC with client-side streaming and TLS encryption
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type StreamerClient struct{}

func (s *StreamerClient) ComputeAverage(stream Streamer_ComputeAverageClient) error {
	numbers := []int64{1, 2, 3, 4, 5} // Example numbers
	for _, num := range numbers {
		if err := stream.Send(&NumberRequest{Number: num}); err != nil {
			return err
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}
	log.Printf("Average: %.2f", resp.Average)
	return nil
}

func main() {
	certFile := "client.crt"
	keyFile := "client.key"
	caFile := "ca.crt"

	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load client certificate: %v", err)
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      caCertPool,
	})

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewStreamerClient(conn)
	stream, err := client.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for {
		resp, err := stream.Recv()
		if err != nil {
			break
		}
		log.Printf("Received response: %v", resp)
	}
}




// gRPC with custom error handling
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name cannot be empty")
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}






// gRPC server with unary interceptor for authorization
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func authInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Example: Perform authorization logic here
	if token := ctx.Value("token"); token == nil || token.(string) != "valid_token" {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
	}
	return handler(ctx, req)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// gRPC server with unary interceptor for logging and recovery
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	panic("test panic") // Simulate a panic
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func loggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("gRPC method: %s, request: %v", info.FullMethod, req)
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v", r)
		}
	}()
	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("gRPC method: %s, error: %v", info.FullMethod, err)
	}
	return resp, err
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// gRPC client with retry and timeout
package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retryInterceptor),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Greeting: %s", resp.Message)
}

func retryInterceptor(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	for attempt := 0; attempt < 3; attempt++ {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err == nil {
			return nil
		}
		st, ok := status.FromError(err)
		if !ok || st.Code() != codes.Unavailable {
			return err
		}
		time.Sleep(100 * time.Millisecond)
	}
	return status.Error(codes.Unavailable, "Service unavailable")
}






// gRPC server with TLS and mutual authentication
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	certFile := "server.crt"
	keyFile := "server.key"
	caFile := "ca.crt"

	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load server certificate: %v", err)
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started with TLS on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// gRPC client with custom metadata
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	md := metadata.Pairs("authorization", "Bearer token")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	resp, err := client.SayHello(ctx, &HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Greeting: %s", resp.Message)
}





// gRPC server with custom metadata interceptor
package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Printf("Received metadata: %v", md)
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func metadataInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		md = metadata.New(nil)
	}
	md.Append("server-interceptor", "true")
	newCtx := metadata.NewIncomingContext(ctx, md)
	return handler(newCtx, req)
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.UnaryInterceptor(metadataInterceptor))
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}






// gRPC server with TLS and custom error handling
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/credentials"
)

type GreeterServer struct{}

func (s *GreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name cannot be empty")
	}
	return &HelloResponse{Message: "Hello, " + req.Name + "!"}, nil
}

func main() {
	certFile := "server.crt"
	keyFile := "server.key"
	caFile := "ca.crt"

	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load server certificate: %v", err)
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    caCertPool,
	})

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	RegisterGreeterServer(s, &GreeterServer{})
	log.Println("gRPC server started with TLS on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}





// gRPC client with TLS and custom timeout
package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	certFile := "client.crt"
	keyFile := "client.key"
	caFile := "ca.crt"

	certificate, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("failed to load client certificate: %v", err)
	}

	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatalf("failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      caCertPool,
	})

	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(creds),
		grpc.WithTimeout(2 * time.Second),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := NewGreeterClient(conn)
	ctx := context.Background()

	resp, err := client.SayHello(ctx, &HelloRequest{Name: "Alice"})
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Greeting: %s", resp.Message)
}






