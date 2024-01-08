package main

import (
	"context"
	"fmt"
	"github.com/SaishNaik/ecom_microservices/product/app"
	"github.com/SaishNaik/ecom_microservices/product/config"
	"github.com/SaishNaik/ecom_microservices/product/logger"
	"github.com/SaishNaik/ecom_microservices/product/proxy"
	"github.com/sirupsen/logrus"
	"go.uber.org/automaxprocs/maxprocs"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	_, err := maxprocs.Set()
	if err != nil {
		slog.Error("failed sto set max process", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	cfg, err := config.NewConfig()
	_ = ctx
	_ = cfg
	if err != nil {
		slog.Error("Failed to get config \n", err)
		os.Exit(1)
	}

	slog.Info("init app", "name", cfg.Name, "version", cfg.Version)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	//todo need to check why is this needed
	logrus.SetLevel(logger.ConvertLogLevel(cfg.Log.Level))
	// integrate Logrus with the slog logger
	slog.New(logger.NewLogrusHandler(logrus.StandardLogger()))

	server := grpc.NewServer()
	go func() {
		defer server.GracefulStop()
		<-ctx.Done()
	}()

	//init app
	_, err = app.InitApp(cfg, server, cfg.PG.DNS)
	if err != nil {
		slog.Error("Failed To init app", err)
		cancel()
	}

	//grpc server
	address := fmt.Sprintf("%s:%s", cfg.GRPC.Host, cfg.GRPC.Port)
	network := "tcp"

	l, err := net.Listen(network, address)
	if err != nil {
		slog.Error("Unable to listen for grpc on address", address)
		cancel()
	}

	slog.Info("Starting grpc server on address", address)
	defer func() {
		if err := l.Close(); err != nil {
			slog.Error("Failed to close listener on address", address)
		}
	}()

	go func() {
		err = server.Serve(l)
		if err != nil {
			slog.Error("Failed to start grpc server on address", address)
			cancel()
		}
	}()

	//http server
	mux := http.NewServeMux()
	gw, err := proxy.NewHttpGRPCGateway(ctx, cfg, nil)
	if err != nil {
		slog.Error("Http GRPC gateway failed")
		cancel()
	}

	mux.Handle("/", gw)

	addr := fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	s := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	slog.Info("Starting http server on address:port", addr)
	if err := s.ListenAndServe(); err != nil {
		slog.Error("Http Server Failed to listen on", addr)
		cancel()
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		slog.Info("Notify Received", v)
	case <-ctx.Done():
		slog.Info("Done Received")
	}

}
