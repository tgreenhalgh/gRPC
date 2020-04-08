## server

create `greet_server` dir

create server.go in that directory

```go
package main
import (
  "fmt"
)

func main() {
  fmt.Println("gRPC server running")
}
```

make sure everything is working as expected

```go
go run greet/greet_server/server.go
```

We need a 'server object' to use. Put this just after the imports and just before `func main()`

Eventually, we'll add services which will be added to the type

```go
type server struct{}
```

add to `func main()`:

```go
// start the listener; 50051 is the default grpc port
lis, err := net.Listen("tcp", "0.0.0.0:50051")
if err != nil {
  log.Fatalf("Failed to listen: %v", err)
}
```

```go
// create gRPC server
// make sure "google.golang.org/grpc" shows up in import
s := grpc.NewServer()

// register the service (RegisterGreetServiceServer was generated from greet.proto into the file greet.pb.go)
// it takes a server and the type (`type server struct{}` from the top of the file)
greetpb.RegisterGreetServiceServer(s, &server{})
```

modules need to be working, or it won't find `greetpb`
the import should look something like `"go/gRPC/thomas/greet/greetpb"`

add

```go
if err := s.Serve(lis); err != nil {
  log.Fatalf("failed to serve: %v", err)
}
```

---
at this point, `server.go` should look something like:

```go
package main

import (
  "fmt"
  "go/gRPC/thomas/greet/greetpb"
  "log"
  "net"

  "google.golang.org/grpc"
)

type server struct{}

func main() {
  fmt.Println("gRPC server running")

  // 50051 is the grpc port
  lis, err := net.Listen("tcp", "0.0.0.0:50051")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }
  // create gRPC server
  s := grpc.NewServer()
  // register the service
  greetpb.RegisterGreetServiceServer(s, &server{})

  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
```

you can test again with `go run greet/greet_server/server.go`

you should see: `gRPC server running`

hit ctrl-c to stop the server
