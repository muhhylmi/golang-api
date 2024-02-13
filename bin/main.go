package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	books "golang-api/modules/books/handler"
	carts "golang-api/modules/cart/handler"
	users "golang-api/modules/users/handler"

	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
)

func main() {
	apps := Init()
	apps.Apps.Validator = apps.Validator
	apps.Apps.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This service is running properly tests")
	})

	booksGroup := apps.Apps.Group("/v1/books")
	userGroup := apps.Apps.Group("/v1/users")
	cartGroup := apps.Apps.Group("/v1/cart")
	books.New(apps).Mount(booksGroup)
	users.New(apps).Mount(userGroup)
	carts.New(apps).Mount(cartGroup)

	l := apps.Logger.LogWithContext("server", "main")

	ctx, cancel := context.WithCancel(context.Background())
	g, gCtx := errgroup.WithContext(ctx)

	// run echo server
	g.Go(func() error {
		l.Info("Starting HTTP server", apps.GlobalConfig.HOST)
		return apps.Apps.Start(fmt.Sprintf(":%s", apps.GlobalConfig.HOST))
	})

	// run gRPC server
	g.Go(func() error {
		l.Info("Starting gRPC server", apps.GlobalConfig.GRPC_PORT)
		port, err := net.Listen("tcp", fmt.Sprintf(":%s", apps.GlobalConfig.GRPC_PORT))
		if err != nil {
			return err
		}

		return apps.GRPC.Serve(port)
	})

	// graceful shutdown
	g.Go(func() error {
		l.Info("Backend ready! Waiting for exit signal...")

		<-gCtx.Done()
		l.Info("Gracefully shutting down...")

		l.Info("Stopping gRPC server")
		apps.GRPC.GracefulStop()

		l.Info("Stopping HTTP server")
		httpServerErr := apps.Apps.Shutdown(ctx)
		if httpServerErr != nil {
			l.Error("Error while shutting down HTTP server", httpServerErr)
		}

		return nil
	})

	// listen for OS signal
	go func() {
		c := make(chan os.Signal, 1)

		signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		l.Info("Listening for OS signal...")
		<-c
		l.Info("Received OS signal, canceling context...")
		cancel()
	}()

	if err := g.Wait(); err != nil {
		fmt.Printf("Exit reason: %s \n", err)
	}
}
