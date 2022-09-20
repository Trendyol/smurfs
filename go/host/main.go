package main

import (
	"github.com/trendyol/smurfs/host/pkg/environment"
	installation "github.com/trendyol/smurfs/host/pkg/install"
	"github.com/trendyol/smurfs/host/pkg/plugin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func HomeDir() string {
	if runtime.GOOS == "windows" {

		// First prefer the HOME environmental variable
		if home := os.Getenv("HOME"); len(home) > 0 {
			if _, err := os.Stat(home); err == nil {
				return home
			}
		}
		if homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"); len(homeDrive) > 0 && len(homePath) > 0 {
			homeDir := homeDrive + homePath
			if _, err := os.Stat(homeDir); err == nil {
				return homeDir
			}
		}
		if userProfile := os.Getenv("USERPROFILE"); len(userProfile) > 0 {
			if _, err := os.Stat(userProfile); err == nil {
				return userProfile
			}
		}
	}
	return os.Getenv("HOME")
}

func main() {
	base := filepath.Join(HomeDir(), ".ty")
	paths := environment.NewPaths(base)
	p := plugin.Plugin{
		TypeMeta: metav1.TypeMeta{
			Kind: "Plugin",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "test",
		},
		Spec: plugin.PluginSpec{
			Version: "v0.0.1",
			Platform: plugin.Runnable{
				URI:    "https://github.com/sysdiglabs/kube-policy-advisor/releases/download/v1.0.2/kube-policy-advisor_v1.0.2_darwin_amd64.tar.gz",
				Sha256: "043e6dd1608eae2b2845db41052fd7876c986fd82392166c176d119554cafbb4",
				Bin:    "kubectl-advise-policy",
			},
		},
	}
	err := installation.Install(paths, p, installation.InstallOpts{})
	if err != nil {
		log.Fatalf("failed to install plugin: %v", err)
	}

	//lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8080"))
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//server := grpc.NewServer()
	//reflection.Register(server)
	//protos.RegisterLogServiceServer(server, service.NewLogService())
	//if err = server.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
}
