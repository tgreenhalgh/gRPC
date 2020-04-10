## add unary API to client

From `c := greetpb.NewGreetServiceClient(cc)` there are now functions available on `c` from `greet.pb.go`, specifically `Greet`

```go
c.Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
```

We need to set the context: context.[`Background()`](https://golang.org/pkg/context/#Background) returns a non-nil, empty Context.

```go
c.Greet(context.Background(), in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
```

Next, we need to set the request (the parameter `in`). Looking at greet.pb.go, see that a GreetRequest takes a type `Greeting`, then add in what `Greeting` takes (FirstName, LastName):

```go
  req := &greetpb.GreetRequest{
    Greeting: &greetpb.Greeting{
      FirstName: "Thomas",
      LastName:  "Greenhalgh",
    },
  }
```

Now that we have our `req`, put that in as a parameter. We don't need the options, so just delete them:

```go
c.Greet(context.Background(), req) (*GreetResponse, error)
```

Use the returns (if err, log it; log res (remembering `Result` is returned))

```go
  res, err := c.Greet(context.Background(), req)
  if err != nil {
    log.Fatalf("error while calling Greet RPC : %v", err)
  }
  log.Printf("Response from Greet: %v", res.Result)
```

Organize the code into its own function. Note that `NewGreetServiceClient` returns a `GreetServiceClient`:

```go
{
  ...
  c := greetpb.NewGreetServiceClient(cc)
  doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
  fmt.Println("Starting to do a Unary RPC...")
  req := &greetpb.GreetRequest{
    Greeting: &greetpb.Greeting{
      FirstName: "Thomas",
      LastName:  "Greenhalgh",
    },
  }
  res, err := c.Greet(context.Background(), req)
  if err != nil {
    log.Fatalf("error while calling Greet RPC : %v", err)
  }
  log.Printf("Response from Greet: %v", res.Result)
}
```

---
all client code should now be:

```go
package main

import (
  "context"
  "fmt"
  "go/gRPC/thomas/greet/greetpb"
  "log"

  "google.golang.org/grpc"
)

func main() {
  fmt.Println("client is running")
  // create the connection
  cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("could not connect: %v", err)
  }

  // close the connection when done
  defer cc.Close()

  c := greetpb.NewGreetServiceClient(cc)
  doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
  fmt.Println("Starting to do a Unary RPC...")
  req := &greetpb.GreetRequest{
    Greeting: &greetpb.Greeting{
      FirstName: "Thomas",
      LastName:  "Greenhalgh",
    },
  }
  res, err := c.Greet(context.Background(), req)
  if err != nil {
    log.Fatalf("error while calling Greet RPC : %v", err)
  }
  log.Printf("Response from Greet: %v", res.Result)
}
```

making sure the server is still running, run the client:

`go run greet/greet_client/client.go`

and you should see something like

```text
Starting to do a Unary RPC...
2020/04/07 17:24:44 Response from Greet: Hello Thomas
```

check the tab where your server is running and you should see something like

```text
Greet function was invoked with greeting:<first_name:"Thomas" last_name:"Greenhalgh" >
```
