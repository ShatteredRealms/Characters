package main

import (
	"context"
	"fmt"
	"github.com/ShatteredRealms/Characters/internal/log"
	"github.com/ShatteredRealms/Characters/internal/option"
	"github.com/ShatteredRealms/Characters/pkg/characterspb"
	"github.com/ShatteredRealms/Characters/pkg/interceptor"
	"github.com/ShatteredRealms/Characters/pkg/repository"
	"github.com/ShatteredRealms/Characters/pkg/service"
	utilRepository "github.com/ShatteredRealms/GoUtils/pkg/repository"
	service2 "github.com/ShatteredRealms/GoUtils/pkg/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
)

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func main() {
	config := option.NewConfig()

	var logger log.LoggerService
	if config.IsRelease() {
		logger = log.NewLogger(log.Info, "")
	} else {
		logger = log.NewLogger(log.Debug, "")
	}

	file, err := ioutil.ReadFile(*config.DBFile.Value)
	if err != nil {
		logger.Log(log.Error, fmt.Sprintf("reading db file: %v", err))
		os.Exit(1)
	}

	c := &utilRepository.DBConnections{}
	err = yaml.Unmarshal(file, c)
	if err != nil {
		logger.Log(log.Error, fmt.Sprintf("parsing db file: %v", err))
		os.Exit(1)
	}

	db, err := utilRepository.DBConnect(*c)
	if err != nil {
		logger.Log(log.Error, fmt.Sprintf("db: %v", err))
		os.Exit(1)
	}

	jwtService, err := service2.NewJWTService(*config.KeyDir.Value)
	if err != nil {
		logger.Log(log.Error, fmt.Sprintf("jwt: %v", err))
		os.Exit(1)
	}

	characterRepository := repository.NewCharacterRepository(db)
	characterService := service.NewCharacterService(characterRepository)

	authInterceptor := interceptor.NewAuthInterceptor(jwtService)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authInterceptor.Unary()),
		grpc.StreamInterceptor(authInterceptor.Stream()),
	)

	characterServer := NewCharacterServer(characterService, logger)
	characterspb.RegisterCharacterServiceServer(grpcServer, characterServer)

	gwmux := runtime.NewServeMux()
	err = characterspb.RegisterCharacterServiceHandlerServer(
		context.Background(),
		gwmux,
		characterServer,
	)

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", config.Address())
	if err != nil {
		logger.Log(log.Error, fmt.Sprintf("listener: %v", err))
		os.Exit(1)
	}

	err = http.Serve(listener, grpcHandlerFunc(grpcServer, gwmux))

	if err != nil {
		logger.Log(log.Error, fmt.Sprintf("srv: %v", err))
		os.Exit(1)
	}
}

type characterServer struct {
	characterspb.UnimplementedCharacterServiceServer

	characterService service.CharacterService
	logger           log.LoggerService

	mu sync.Mutex
}

func (c *characterServer) Create(ctx context.Context, message *characterspb.CreateCharacterMessage) (*characterspb.EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "Create has not been implemented yet")
}

func NewCharacterServer(characterService service.CharacterService, logger log.LoggerService) *characterServer {
	return &characterServer{
		characterService: characterService,
		logger:           logger,
	}
}
