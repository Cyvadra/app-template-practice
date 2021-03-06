package api

import (
	"context"

	"github.com/NpoolPlatform/go-service-app-template/message/npool"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	npool.UnimplementedAppPracticeServer
}

func Register(server grpc.ServiceRegistrar) {
	npool.RegisterAppPracticeServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return npool.RegisterAppPracticeHandlerFromEndpoint(context.Background(), mux, endpoint, opts)
}
