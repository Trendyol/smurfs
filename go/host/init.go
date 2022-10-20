package host

import (
	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/go/host/auth"
	"github.com/trendyol/smurfs/go/host/logger"
	"github.com/trendyol/smurfs/go/host/pkg/archive"
	"github.com/trendyol/smurfs/go/host/pkg/cli"
	"github.com/trendyol/smurfs/go/host/pkg/download"
	"github.com/trendyol/smurfs/go/host/pkg/environment"
	"github.com/trendyol/smurfs/go/host/pkg/plugin"
	"github.com/trendyol/smurfs/go/host/pkg/process"
	"github.com/trendyol/smurfs/go/host/pkg/providers"
	"github.com/trendyol/smurfs/go/host/pkg/verifier"
	"github.com/trendyol/smurfs/go/host/storage"
	"github.com/trendyol/smurfs/go/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

type SmurfHost struct {
	Root *cobra.Command
}

func InitializeHost(options *cli.Options) (*SmurfHost, error) {
	buildDefaultOptions(options)

	up := make(chan struct{})
	go Start(options, up)
	<-up

	paths := environment.NewPaths(options.PluginPath)
	httpClient := http.Client{}
	providers.InitProviders(httpClient)
	execManager := process.NewExec()
	extractor := archive.NewExtractorManager(map[string]archive.Extractor{
		"zip":    archive.NewZipExtractor(),
		"tar.gz": archive.NewTarGzExtractor(),
	})
	downloader := plugin.NewDownloader(paths, providers.GetProviders(), download.NewFileDownloader(&httpClient))
	pluginManager := plugin.NewManager(paths, downloader, extractor, verifier.NewSha256Verifier("sha256"))
	cmdWrapper := cli.NewCmdWrapper(execManager, pluginManager)
	builder := cli.NewBuilder(paths, cmdWrapper)

	return &SmurfHost{
		Root: builder.Build(options),
	}, nil
}

func buildDefaultOptions(options *cli.Options) {
	if options.HostAddress == "" {
		options.HostAddress = "localhost:50051"
	}

	if options.RootCmd == nil {
		options.RootCmd = &cobra.Command{
			Use:   "host",
			Short: "Host CLI",
			Long:  "Host CLI",
		}
	}
}

func Start(options *cli.Options, up chan struct{}) {
	lis, err := net.Listen("tcp", options.HostAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if up != nil {
		close(up)
	}

	s := grpc.NewServer()
	protos.RegisterLogServiceServer(s, logger.NewLogService(options.Logger))
	protos.RegisterAuthServiceServer(s, auth.NewAuthService(options.Auth))
	protos.RegisterMetadataStorageServiceServer(s, storage.NewMetadataStorageService(options.MetadataStorage))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (host SmurfHost) Execute() error {
	if err := host.Root.Execute(); err != nil {
		return err
	}

	return nil
}
