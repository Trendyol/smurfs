# Go Library

## Overview

The Go library has been designed to be as simple as possible to use. It has Host and Client Libraries that are used to
create and interact with a MicroCLI. The Host Library is used to register host functions and the Client Library is used
to call client functions.

## Host Library

The Host Library is used to register host functions. Host functions are functions that are called by the client

### Installation

To install the Host Library, run the following command:

```bash

go get github.com/trendyol/smurfs

```

### Usage

To use the Host Library, import the library into your project:

```go

import "github.com/trendyol/smurfs/host"

var plugins = []plugin.Plugin{
    {
        Name:             "micro1",
        ShortDescription: "Micro CLI",
        LongDescription:  "Micro CLI",
        Usage:            "micro",
        Source: map[string]interface{}{
            "type":   "binary",
            "binary": "./micro1",
        },
    },
    {
        Name:             "micro2",
        ShortDescription: "Micro CLI",
        LongDescription:  "Micro CLI",
        Usage:            "micro1",
        Source: map[string]interface{}{
            "type":   "binary",
            "binary": "./micro2",
        },
    },
    {
        Name:             "onboarding",
        ShortDescription: "Onboarding CLI",
        LongDescription:  "Onboarding CLI",
        Usage:            "onboarding",
        Source: map[string]interface{}{
            "type":   "gitlab",
            "source": "https://gitlab.com/trendyol/onboarding-cli.git",
        },
    },
}

func main() {
    logger := &Logger{}
    smurfHost, err := host.InitializeHost(host.Options{
        Plugins: plugins,
        RootCmd: &cobra.Command{
            Use:   "host",
            Short: "Host CLI",
            Long:  "Host CLI",
        },
        Logger:  logger,
    })

    if err != nil {
        panic(err)
    }
    
    if err := smurfHost.Execute(); err != nil {
        panic(err)
    }
}


type Logger struct {
}

func (l *Logger) Debug(message string, args ...interface{}) {
    fmt.Printf("HOST-DEBUG: %s\n", message)
}

func (l *Logger) Info(message string, args ...interface{}) {
    fmt.Printf("HOST-INFO: %s\n", message)
}

func (l *Logger) Warn(message string, args ...interface{}) {
    fmt.Printf("HOST-WARN: %s\n", message)
}

func (l *Logger) Error(message string, args ...interface{}) {
    fmt.Printf("HOST-ERROR: %s\n", message)
}

func (l *Logger) Fatal(message string, args ...interface{}) {
    fmt.Printf("HOST-FATAL: %s\n", message)
}


```

### Plugin Structure

The plugin structure is as follows:

```go

type Plugin struct {
    Name             string
    ShortDescription string
    LongDescription  string
    Usage            string
    Source           map[string]interface{}
}

```

You can build static or fetch from available plugins and register them to the host. The source of the plugin can be a JSON file or git repository. When you execute the subcommand, the plugin will be download and execute the subcommand.



## Client Library

The Client Library is used to call client functions. Client functions are functions that are called by the host

### Installation

To install the Client Library, run the following command:

```bash

go get github.com/trendyol/smurfs

```

### Usage

To use the Client Library, import the library into your project:

```go

import "github.com/trendyol/smurfs/client"

func main() {
    host := "localhost:50051"
    smurfs, err := client.InitializeClient(client.Options{
        HostAddress: &host,
    })

    if err != nil {
        panic(err)
    }
	
	// Implement your logic here
	
	smurfs.Logger.Info("Send info logs")

	// Implement another logics here
	
	smurfs.Logger.Error("Send error logs")
	
    err = smurfs.Close()
    if err != nil {
        panic(err)
    }
}

```