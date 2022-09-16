# Project Structure

We are designing multi projects in one repository so we need to describe folder/project/package structure for languages & stacks. Every language & stack has own project structure. If you would like to be contributor. You should read this doc. It will help and onboard you for projects

## Development Library

* Host and Client Library into one package
* Every standalone application can be host, client, or both

### Go

If you would like to contribute or develop Go library. You have to develop with this structure

```md

- go
  - client
  - host
  - internal
  - pkg

```

### go/client

Client specific code blocks should be into `go/client/` It's only for client operations

### go/host

Host specific code blocks should be into `go/host/` It's only for host operations

### go/internal

You must add utils, helpers, and other code blocks into internals. Internals mean doesn't share to referenced project

### go/pkg

You should add adapters, providers and core components. For example we implemented Gitlab Adapter and It's in `pkg/dvc/` package. If you want to implement Github Adapter. You have to put into `pkg/dvc/`. In another case, you would like to implement UI components into `pkg/ui/`. You can share with host and client packages
