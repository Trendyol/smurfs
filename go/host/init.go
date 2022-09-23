package host

import (
	"fmt"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SmurfHost struct {
	Root   *cobra.Command
	Logger *Logger
}

type Options struct {
	Plugins     []Plugin
	RootCmd     *cobra.Command
	HostAddress string
}

type PluginBinary struct {
	Name   string
	Target string
	Args   interface{}
}

type Plugin struct {
	Name             string
	ShortDescription string
	LongDescription  string
	Usage            string
	Binaries         []PluginBinary
}

func InitializeHost(options Options) (*SmurfHost, error) {
	// TODO(peacecwz): register commands and flags to root command
	// TODO(peacecwz): register plugins
	if options.RootCmd == nil {
		options.RootCmd = &cobra.Command{
			Use:   "host",
			Short: "Host CLI",
			Long:  "Host CLI",
		}
	}

	for _, plugin := range options.Plugins {
		cmd := &cobra.Command{
			Use:   plugin.Name,
			Short: plugin.ShortDescription,
			Long:  plugin.LongDescription,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Micro CLI is running")
			},
		}

		options.RootCmd.AddCommand(cmd)
	}

	if options.HostAddress == "" {
		options.HostAddress = ":50051"
	}

	lis, err := net.Listen("tcp", options.HostAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
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
