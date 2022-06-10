package interceptor

import (
	"context"
	"github.com/ShatteredRealms/GoUtils/pkg/service"
	"google.golang.org/grpc"
	"log"
)

type authInterceptor struct {
	jwtService service.JWTService
}

func NewAuthInterceptor(jwtService service.JWTService) *authInterceptor {
	return &authInterceptor{
		jwtService: jwtService,
	}
}

func (interceptor *authInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)
		return handler(ctx, req)
	}
}

func (interceptor *authInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		log.Println("--> stream interceptor: ", info.FullMethod)
		return handler(srv, stream)
	}
}
