# vmapi

### QuickStart

```
go run cmd/vmapi/main.go
```

### API Server Options

```
$ go run cmd/vmapi/main.go --help
Usage:
  main [OPTIONS]

Options for VM APIs

Application Options:
      --host= The IP to listen on (default: 0.0.0.0) [$HOST]
      --port= The port to listen on (default: 8080) [$PORT]

Help Options:
  -h, --help  Show this help message
```

### Project Structure

```
├── Makefile                # build
├── README.md
├── cache                   # cache implement vmstore
│   └── vm.go
├── cmd
│   └── vmapi
│       └── main.go         # entrypoint
├── filestore               # filestore implement vmstore
│   └── vm.go
├── go.mod
├── go.sum
├── http                    # http (api)
│   ├── http.go             # interface
│   ├── mux.go              # mux (route)
│   ├── server.go           # http server
│   ├── service.go          # api service
│   └── vm.go               # vm rest api
└── vmapi.go                # interface
```
