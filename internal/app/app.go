package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"os"

	gRPCctrl "github.com/Artenso/Geo-Provider/internal/controller/grpc_geo_provider"
	jsonRPCctrl "github.com/Artenso/Geo-Provider/internal/controller/json_rpc_geo_provider"
	"github.com/Artenso/Geo-Provider/internal/service"
	desc "github.com/Artenso/Geo-Provider/pkg/grpc_geo_provider"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	service           service.GeoProvider
	gRPCcontroller    *gRPCctrl.Controller
	jsonRPCcontroller *jsonRPCctrl.Controller
	gRPCServer        *grpc.Server
	jsonRPCServer     net.Listener
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run(ctx context.Context) error {

	group, _ := errgroup.WithContext(ctx)

	group.Go(func() error {
		list, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
		if err != nil {
			return fmt.Errorf("failed to mapping port: %s", err.Error())
		}

		if err := a.gRPCServer.Serve(list); err != nil {
			return fmt.Errorf("failed to run server: %s", err.Error())
		}

		return nil
	})

	group.Go(func() error {
		for {
			conn, err := a.jsonRPCServer.Accept()
			if err != nil {
				return fmt.Errorf("accept failed: %s", err)
			}

			go jsonrpc.ServeConn(conn)
		}
	})

	if err := group.Wait(); err != nil {
		return err
	}

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(ctx context.Context) error{
		a.initConfig,
		a.initService,
		a.initGrpcController,
		a.initJSONrpcController,
		a.initGRPCServer,
		a.initJSONrpcSrver,
	}

	for _, init := range inits {
		err := init(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
		return err
	}

	return nil
}

func (a *App) initService(_ context.Context) error {
	a.service = service.NewGeoService(os.Getenv("DADATA_APIKEY"), os.Getenv("DADATA_SECRETKEY"))
	return nil
}

func (a *App) initGrpcController(_ context.Context) error {
	a.gRPCcontroller = gRPCctrl.NewController(a.service)
	return nil
}

func (a *App) initJSONrpcController(_ context.Context) error {
	a.jsonRPCcontroller = jsonRPCctrl.NewController(a.service)
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	s := grpc.NewServer()

	desc.RegisterGeoProviderServer(s, a.gRPCcontroller)

	reflection.Register(s)

	a.gRPCServer = s

	return nil
}

func (a *App) initJSONrpcSrver(_ context.Context) error {
	rpc.Register(a.jsonRPCcontroller)

	listener, err := net.Listen("tcp", os.Getenv("JSON_RPC_PORT"))
	if err != nil {
		return err
	}

	a.jsonRPCServer = listener

	return nil
}
