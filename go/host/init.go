package host

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
	"github.com/trendyol/smurfs/go/host/pkg/process"
	"github.com/trendyol/smurfs/go/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
	"net"
)

var hostLogger Logger

type SmurfHost struct {
	Root *cobra.Command
}

type Options struct {
	Plugins     []plugin.Plugin
	RootCmd     *cobra.Command
	HostAddress string
	Logger      Logger
}

type LogService struct {
	protos.UnimplementedLogServiceServer
}

func (l LogService) Info(stream protos.LogService_InfoServer) error {
	for {
		l, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&emptypb.Empty{})
		}

		if err != nil {
			return err
		}

		hostLogger.Info(l.Msg)
	}
}

func Start(hostAddress string, up chan struct{}) {
	lis, err := net.Listen("tcp", hostAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if up != nil {
		close(up)
	}

	s := grpc.NewServer()
	protos.RegisterLogServiceServer(s, &LogService{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func InitializeHost(options Options) (*SmurfHost, error) {
	if options.HostAddress == "" {
		options.HostAddress = "localhost:50051"
	}

	hostLogger = options.Logger

	up := make(chan struct{})
	go Start(options.HostAddress, up)
	<-up
	execManager := process.NewExec()
	if options.RootCmd == nil {
		options.RootCmd = &cobra.Command{
			Use:   "host",
			Short: "Host CLI",
			Long:  "Host CLI",
		}
	}

	for _, pl := range options.Plugins {
		cmd := &cobra.Command{
			Use:   pl.Name,
			Short: pl.ShortDescription,
			Long:  pl.LongDescription,
			Run: func(cmd *cobra.Command, args []string) {
				var currentPlugin plugin.Plugin
				for _, p := range options.Plugins {
					if p.Name == cmd.Name() {
						currentPlugin = p
						break
					}
				}

				err := execManager.Run(context.Background(), currentPlugin, args...)
				if err != nil {
					log.Fatalf(err.Error())
				}
				fmt.Println("Done")
			},
		}

		options.RootCmd.AddCommand(cmd)
	}

	return &SmurfHost{
		Root: options.RootCmd,
	}, nil
}

func (host SmurfHost) Execute() error {
	if err := host.Root.Execute(); err != nil {
		return err
	}

	return nil
}
