package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	service2 "github.com/alexadastra/hisho/hisho-bot/internal/app/service"

	"github.com/alexadastra/ramme/config"
	"github.com/alexadastra/ramme/service"
	"github.com/alexadastra/ramme/system"

	advanced "github.com/alexadastra/hisho/hisho-bot/internal/config"
	"github.com/alexadastra/hisho/hisho-bot/internal/swagger"

	"github.com/alexadastra/hisho/hisho-bot/pkg/api"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	// Load ENV configuration
	config.ServiceName = "HISHO-BOT"
	basicConfManager, basicConfWatcher, err := config.InitBasicConfig()
	if err != nil {
		panic(err)
	}
	basicConfig := basicConfManager.GetBasic()

	advancedConfManager, advancedConfWatcher, err := advanced.InitAdvancedConfig()
	if err != nil {
		panic(err)
	}
	advancedConfig := advancedConfManager.Get()

	// Configure service and get router
	router, logger, err := service.Setup(basicConfig)
	if err != nil {
		logger.Fatal(err)
	}

	// TODO: figure out how to use advanced config
	logger.Info(advancedConfig)

	// Setup gRPC servers.
	baseGrpcServer := grpc.NewServer()
	userGrpcServer := service2.NewHishoBot()
	api.RegisterHishoBotServer(baseGrpcServer, userGrpcServer)

	// Setup gRPC gateway.
	ctx := context.Background()
	rmux := runtime.NewServeMux()
	mux := http.NewServeMux()
	mux.Handle("/", rmux)
	{
		err = api.RegisterHishoBotHandlerServer(ctx, rmux, userGrpcServer)
		if err != nil {
			logger.Fatal(err)
		}
	}

	if basicConfig.IsLocalEnvironment {
		mux.Handle(swagger.Pattern, swagger.HandlerLocal)
	} else {
		mux.Handle(swagger.Pattern, swagger.HandlerK8S)
	}

	// Setup secondary HTTP handlers
	// Listen and serve handlers
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprintf("%s:%d", basicConfig.Host, basicConfig.HTTPSecondaryPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Serve
	g := system.NewGroupOperator()

	g.Add(func() error {
		return basicConfWatcher.Run()
	}, func(err error) {
		_ = basicConfWatcher.Close()
	})
	g.Add(func() error {
		return advancedConfWatcher.Run()
	}, func(err error) {
		_ = advancedConfWatcher.Close()
	})

	grpcListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", basicConfig.Host, basicConfig.GRPCPort))
	if err != nil {
		logger.Fatal(err)
	}
	g.Add(func() error {
		logger.Warnf("Serving grpc address %s", fmt.Sprintf("%s:%d", basicConfig.Host, basicConfig.GRPCPort))
		return baseGrpcServer.Serve(grpcListener)
	}, func(error) {
		_ = grpcListener.Close()
	})

	httpListener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", basicConfig.Host, basicConfig.HTTPPort))
	if err != nil {
		logger.Fatal(err)
	}
	g.Add(func() error {
		logger.Warnf("Serving http address %s", fmt.Sprintf("%s:%d", basicConfig.Host, basicConfig.HTTPPort))
		return http.Serve(httpListener, mux)
	},
		func(err error) {
			_ = httpListener.Close()
		})

	g.Add(func() error {
		return srv.ListenAndServe()
	}, func(err error) {})

	signals := system.NewSignals()
	g.Add(func() error {
		return signals.Wait(logger, g)
	}, func(error) {})

	if err := g.Run(); err != nil {
		logger.Fatal(err)
	}
}
