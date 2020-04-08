## adding unary API to the server

The server needs to implement the `GreetServiceServer interface`

Copy the interface code just before `func main()`

```go
Greet(context.Context, *GreetRequest) (*GreetResponse, error)
```

fill it out with a pointer to our server

```go
func (*server) Greet(context.Context, *GreetRequest) (*GreetResponse, error) {

}
```

`context` should be in imports, add in context parameter (ctx).

```go
func (*server) Greet(ctx context.Context, *GreetRequest) (*GreetResponse, error) {

}
```

`GreetRequest` & `GreetResponse` are not available, but are in the generated file, so add `*greetpb.GreetRequest` as type with `req` as the parameter and `*greetpb.GreetResponse`

```go
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

}
```

remember, `(*greetpb.GreetResponse, error)` is the return type.

---
Filling out the function.

The `req` object has `GetGreeting()` (find it in the greet.pb.go file) and `Greeting` has `GetFirstName()`

To form the `GreetResponse` (look at greet.pb.go to see the details), it takes a struct that takes a Result as a string. Notice that the parameter `*greetpb.GreetResponse` is a pointer, so need `&greetpb.GreetResponse` when building the struct

```go
func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
  fmt.Printf("Greet function was invoked with %v\n", req)
  // extract the first name
  firstName := req.GetGreeting().GetFirstName()
  // make our string
  result := "Hello " + firstName
  // build the struct
  res := &greetpb.GreetResponse{
    Result: result,
  }
  return res, nil
}
```

final `server.go` code should look like:

```go
package main

import (
    "context"
    "fmt"
    "go/gRPC/thomas/greet/greetpb"
    "log"
    "net"

    "google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
    fmt.Printf("Greet function was invoked with %v\n", req)
    firstName := req.GetGreeting().GetFirstName()
    result := "Hello " + firstName
    res := &greetpb.GreetResponse{
        Result: result,
    }
    return res, nil
}

func main() {
    fmt.Println("gRPC server running")

    // start the listener; 50051 is the default grpc port
    lis, err := net.Listen("tcp", "0.0.0.0:50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    // create gRPC server
    s := grpc.NewServer()
    greetpb.RegisterGreetServiceServer(s, &server{})

    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```

Run the server and let's build the client.

