package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/trendyol/smurfs/host/protos"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "ty",
}

var commandManifests = []*protos.Command{
	{
		Name:        "login",
		Description: "to login with LDAP credentials",
		Flags: []*protos.CommandFlag {
			{
				Name: "email",
				Required: true,
				Repeated: false,
				Description: "User email",
			},
			{
				Name: "password",
				Required: true,
				Repeated: false,
				Description: "User password",
			},
		},
	},
	{
		Name:        "logout",
		Description: "to logout from CLI",
		Flags: []*protos.CommandFlag{
			{
				Name: "email",
				Required: false,
				Repeated: false,
				Description: "User email",
			},
		},

	},
}

func main() {
	for _, command := range commandManifests {
		subCommand := &cobra.Command{
			Use: command.Name,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Running the command " + command.Name)
				for i := range args {
					fmt.Println(i)
				}
			},
			Short:   command.Description,
			Example: command.Usage,
		}
		rootCmd.AddCommand(subCommand)
		for _, flag := range command.Flags {
			subCommand.Flags().StringP(flag.Name, flag.Short, "", flag.Description)
			if flag.Required {
				subCommand.MarkFlagRequired(flag.Name)
			}
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
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
