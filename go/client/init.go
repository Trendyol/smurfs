package client

import (
	"context"
	"flag"
	"github.com/trendyol/smurfs/go/client/pkg/service"
	"google.golang.org/grpc"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type SmurfClient struct {
	Logger  *service.LoggerClient
	Auth    *service.AuthClient
	Storage *service.StorageClient
	Close   func() error
}

type Options struct {
	HostAddress *string
}

func InitializeClient(opt Options) (*SmurfClient, error) {
	if opt.HostAddress == nil {
		flag.StringVar(opt.HostAddress, "host", "localhost:8080", "host address")
	}
	ctx, cancel := context.WithCancel(context.Background())

	dial, err := grpc.Dial(*opt.HostAddress, grpc.WithBlock(), grpc.WithTimeout(3*time.Second), grpc.WithInsecure())
	if err != nil {
		log.Printf("failed to dial: %+v", err)
		return nil, err
	}

	loggerClient, err := service.NewLoggerClient(dial, ctx)
	if err != nil {
		return nil, err
	}

	authClient, err := service.NewAuthClient(dial)
	if err != nil {
		return nil, err
	}

	storageClient, err := service.NewStorageClient(dial)
	if err != nil {
		return nil, err
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-ctx.Done():
			return
		case s := <-sigCh:
			log.Printf("got signal %v, attempting graceful shutdown", s)
			cancel()
		}
	}()

	smurf := &SmurfClient{
		Logger:  loggerClient,
		Auth:    authClient,
		Storage: storageClient,
		Close: func() error {
			// TODO: implement wait all requests to finish
			time.Sleep(500 * time.Millisecond)
			return dial.Close()
		},
	}
	return smurf, nil
}
