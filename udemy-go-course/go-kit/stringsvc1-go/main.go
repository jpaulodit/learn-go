package main

import (
	"context"
	"errors"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	pb "github.com/jpaulodit/learn-go/go-kit/stringsvc1-go/stringsvc1"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

var ErrEmpty = errors.New("empty string")

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(pb.UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return pb.UppercaseResponse{V: v, Err: err.Error()}, nil
		}
		return pb.UppercaseResponse{V: v, Err: ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(pb.CountRequest)
		v := svc.Count(req.S)
		return pb.CountResponse{V: int32(v)}, nil
	}
}

func NewService() StringService {
	return stringService{}
}

type Set struct {
	CountEndpoint     endpoint.Endpoint
	UppercaseEndpoint endpoint.Endpoint
}

func NewEndpoints(svc StringService) Set {
	var countEndpoint endpoint.Endpoint
	countEndpoint = makeCountEndpoint(svc)

	var uppercaseEndpoint endpoint.Endpoint
	uppercaseEndpoint = makeUppercaseEndpoint(svc)

	return Set{
		CountEndpoint:     countEndpoint,
		UppercaseEndpoint: uppercaseEndpoint,
	}
}

type grpcServer struct {
	count     grpctransport.Handler
	uppercase grpctransport.Handler
}

func NewGRPCServer(endponts Set) pb.StringServiceServer {
	return &grpcServer{
		count: grpctransport.NewServer(
			endponts.CountEndpoint,
			decodeGRPCCountRequest,
			encodeGRPCCountResponse,
		),
		uppercase: grpctransport.NewServer(
			endponts.UppercaseEndpoint,
			decodeGRPCUppercaseRequest,
			encodeGRPCUppercaseResponse,
		),
	}
}

func (s *grpcServer) Count(ctx context.Context, req *pb.CountRequest) (*pb.CountResponse, error) {
	_, rep, err := s.count.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.CountResponse), nil
}

func (s *grpcServer) Uppercase(ctx context.Context, req *pb.UppercaseRequest) (*pb.UppercaseResponse, error) {
	_, rep, err := s.uppercase.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UppercaseResponse), nil
}

type CountRequest struct {
	S string
}

type UppercaseRequest struct {
	S string
}

type CountResponse struct {
	V int32
}

type UppercaseResponse struct {
	V   string
	Err error
}

func decodeGRPCCountRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.CountRequest)
	return CountRequest{S: req.S}, nil
}

func decodeGRPCUppercaseRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.UppercaseRequest)
	return UppercaseRequest{S: req.S}, nil
}

func encodeGRPCCountResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(CountResponse)
	return &pb.CountResponse{V: int32(resp.V)}, nil
}

func err2str(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}

func encodeGRPCUppercaseResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(UppercaseResponse)
	return &pb.UppercaseResponse{V: resp.V, Err: err2str(resp.Err)}, nil
}

func main() {

	var (
		service    = NewService()
		endpoints  = NewEndpoints(service)
		grpcServer = NewGRPCServer(endpoints)
	)

	grpcListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error with grpc listener: %v", err)
	}

	baseServer := grpc.NewServer(grpc.UnaryInterceptor(grpctransport.Interceptor))
	pb.RegisterStringServiceServer(baseServer, grpcServer)
	if err := baseServer.Serve(grpcListener); err != nil {
		log.Fatalf("failed to service: %v", err)
	}
}
