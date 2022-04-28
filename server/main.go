package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/jfmrm/grpc-loadbalancer-benchmark/helloworld"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

var (
	port        = flag.Int("port", 50051, "The server port")
	metricsPort = flag.Int("metrics-port", 2112, "The metrics port")
	latencyHist = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "say_hello_latency",
		Help:    "Histogram of processing latency of the say hello grpc method.",
		Buckets: prometheus.DefBuckets,
	})
	requestCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "say_hello_req_count",
		Help: "Counter to cont requests.",
	})
)

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	timer := prometheus.NewTimer(latencyHist)
	defer timer.ObserveDuration()
	requestCounter.Inc()
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterHelloServiceServer(s, &server{})

	go func() {
		log.Printf("Starting server on port %d", *port)
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	go func() {
		prometheus.MustRegister(latencyHist)
		prometheus.MustRegister(requestCounter)
		log.Printf("Starting metrics server on port %d", *metricsPort)
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":2112", nil)
	}()

	<-done

	log.Println("Shutting down server...")
	s.GracefulStop()
}
